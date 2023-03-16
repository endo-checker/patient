package main

import (
	"net/http"
	"os"

	sv "github.com/endo-checker/common/server"
	st "github.com/endo-checker/common/store"
	"github.com/joho/godotenv"

	"github.com/endo-checker/patient/handler"
	patientv1 "github.com/endo-checker/patient/internal/gen/patient/v1"
	pbcnn "github.com/endo-checker/patient/internal/gen/patient/v1/patientv1connect"
)

type Server struct {
	*http.ServeMux
}

func main() {
	port := "localhost:8080"
	godotenv.Load()
	uri := os.Getenv("MONGO_URI")

	svc := &handler.PatientServer{
		Store: st.Connect[*patientv1.Patient](uri, "patient"),
	}

	path, hndlr := pbcnn.NewPatientServiceHandler(svc)

	sv.Server.ConnectServer(
		sv.Server{
			ServeMux: &http.ServeMux{},
		},
		path,
		hndlr,
		port,
	)
}
