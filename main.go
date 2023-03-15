package main

import (
	"github.com/endo-checker/common/server"

	"github.com/endo-checker/patient/handler"
	pbcnn "github.com/endo-checker/patient/internal/gen/patient/v1/patientv1connect"
	"github.com/endo-checker/patient/store"
)

func main() {
	port := "localhost:8080"

	s := store.Connect()
	svc := &handler.PatientServer{
		Store: s,
	}

	path, h := pbcnn.NewPatientServiceHandler(svc)
	server.ConnectServer(path, h, port)
}
