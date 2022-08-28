package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/endo-checker/patient/handler"
	"github.com/endo-checker/patient/store"

	pb "github.com/endo-checker/patient/gen/proto/go/patient/v1"
)

func main() {
	
	grpcSrv := grpc.NewServer()
	defer grpcSrv.Stop()         // stop server on exit
	reflection.Register(grpcSrv) // for postman

	h := &handler.PatientServer{
		Store: store.Connect(),
	}
	pb.RegisterPatientServiceServer(grpcSrv, h)

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
