package main

import (
	// "github.com/endo-checker/common/store"

	"github.com/endo-checker/common/server"
	"github.com/endo-checker/patient/handler"
	pbcnn "github.com/endo-checker/patient/internal/gen/patient/v1/patientv1connect"
	"github.com/endo-checker/patient/store"
	"github.com/joho/godotenv"
)

func main() {
	port := "localhost:8080"
	godotenv.Load()
	// uri := os.Getenv("mongouri")

	s := store.Connect()
	svc := &handler.PatientServer{
		Store: s,
	}

	path, h := pbcnn.NewPatientServiceHandler(svc)
	server.ConnectServer(path, h, port)
}
