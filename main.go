package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	gw "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/spf13/viper"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	pb "github.com/endo-checker/patient/gen/proto/go/patient/v1"
	"github.com/endo-checker/patient/handler"
	"github.com/endo-checker/patient/store"
)

func main() {
	defPort := os.Getenv("PORT")
	if defPort == "" {
		defPort = "8080"
	}

	grpcSrv := grpc.NewServer()
	defer grpcSrv.Stop()         // stop server on exit
	reflection.Register(grpcSrv) // for postman

	h := &handler.PatientServer{
		Store: store.Connect(),
	}

	hm := gw.WithIncomingHeaderMatcher(func(key string) (string, bool) {
		switch key {
		case "X-Token-C-Tenant", "X-Token-C-User", "Permissions":
			return key, true
		default:
			return gw.DefaultHeaderMatcher(key)
		}

	})

	mo := gw.WithMarshalerOption("*", &gw.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			EmitUnpopulated: false,
		},
	})

	pb.RegisterPatientServiceServer(grpcSrv, h)
	httpMux := gw.NewServeMux(hm, mo)
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterPatientServiceHandlerFromEndpoint(context.Background(), httpMux, ":"+defPort, dopts); err != nil {
		log.Fatal(err)
	}

	mux := httpGrpcMux(httpMux, grpcSrv)
	httpSrv := &http.Server{
		Addr:    ":" + defPort,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	if err := httpSrv.ListenAndServe(); err != http.ErrServerClosed {
		return
	}
}

func httpGrpcMux(httpHandler http.Handler, grpcServer *grpc.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			if allowedOrigin(r.Header.Get("Origin")) {
				w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
			}
			if r.Method == "OPTIONS" {
				return
			}
			httpHandler.ServeHTTP(w, r)
		}
	})
}

func allowedOrigin(origin string) bool {
	if viper.GetString("cors") == "*" {
		return true
	}
	if matched, _ := regexp.MatchString(viper.GetString("cors"), origin); matched {
		return true
	}
	return false
}
