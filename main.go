package main

import (
	"log"
	"net/http"
	"os"

	sv "github.com/endo-checker/protostore/server"
	"github.com/joho/godotenv"

	"github.com/endo-checker/patient/handler"
	pbcnn "github.com/endo-checker/patient/internal/gen/patient/v1/patientv1connect"
	"github.com/endo-checker/patient/store"
)

type Server struct {
	*http.ServeMux
}

var addr = "8080"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env found: %v", err)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = addr
	}
	uri := os.Getenv("MONGO_URI")

	svc := &handler.PatientServer{
		Store: store.New(uri),
	}
	path, hndlr := pbcnn.NewPatientServiceHandler(svc)

	srvr := sv.Server{
		ServeMux: &http.ServeMux{},
	}

	if err := srvr.ConnectServer(path, ":"+port, hndlr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
