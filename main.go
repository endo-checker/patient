package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/felixge/httpsnoop"
	gw "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/endo-checker/patient/gen/proto/go/patient/v1"
	"github.com/endo-checker/patient/handler"
	"github.com/endo-checker/patient/store"
)

const defPort = "localhost:8080"

func main() {

	grpcSrv := grpc.NewServer()
	defer grpcSrv.Stop()         // stop server on exit
	reflection.Register(grpcSrv) // for postman

	h := &handler.PatientServer{
		Store: store.Connect(),
	}

	pb.RegisterPatientServiceServer(grpcSrv, h)
	mux := gw.NewServeMux()

	dopts := []grpc.DialOption{grpc.WithInsecure()}

	err := pb.RegisterPatientServiceHandlerFromEndpoint(context.Background(), mux, defPort, dopts)
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Handler: withLogger(mux),
	}
	// creating a listener for server
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	m := cmux.New(lis)
	
	httpL := m.Match(cmux.HTTP1Fast())

	grpcL := m.Match(cmux.HTTP2())

	go server.Serve(httpL)
	
	go grpcSrv.Serve(grpcL)
	
	m.Serve()


}

func withLogger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, writer, request)
		log.Printf("http[%d]-- %s -- %s\n", m.Code, m.Duration, request.URL.Path)
	})
}
