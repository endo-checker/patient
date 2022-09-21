package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	daprpb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	dapr "github.com/dapr/go-sdk/client"
	gw "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	sv "github.com/endo-checker/common/server"
	pb "github.com/endo-checker/patient/gen/proto/go/patient/v1"
	"github.com/endo-checker/patient/handler"
	"github.com/endo-checker/patient/store"
)

func main() {

	// initiate dapr
	time.Sleep(2 * time.Second)
	client, err := dapr.NewClient()
	if err != nil {
		log.Fatalf("failed to initialise Dapr client: %v", err)
	}
	defer client.Close()

	// initiate port
	defPort := os.Getenv("PORT")
	// if defPort == "" {
	// 	defPort = "8080"
	// }

	grpcSrv := grpc.NewServer()
	defer grpcSrv.Stop()         // stop server on exit
	reflection.Register(grpcSrv) // for postman

	h := &handler.PatientServer{
		Store: store.Connect(),
		Dapr:  client,
	}

	ch := handler.CallbackServer{}
	daprpb.RegisterAppCallbackServer(grpcSrv, ch)

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

	mux := sv.HttpGrpcMux(httpMux, grpcSrv)
	httpSrv := &http.Server{
		Addr:    ":" + defPort,
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}

	if err := httpSrv.ListenAndServe(); err != http.ErrServerClosed {
		return
	}
}
