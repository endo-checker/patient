package main

import (
	"net/http"
	"os"

	sv "github.com/endo-checker/common/server"
	"github.com/joho/godotenv"

	"github.com/endo-checker/patient/handler"
	pbcnn "github.com/endo-checker/patient/internal/gen/patient/v1/patientv1connect"
	"github.com/endo-checker/patient/store"
)

type Server struct {
	*http.ServeMux
}

var addr = ":8080"

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")
	if port != "" {
		addr = ":" + port
	}
	uri := os.Getenv("MONGO_URI")

	svc := &handler.PatientServer{
		Store: store.New(uri),
	}
	path, hndlr := pbcnn.NewPatientServiceHandler(svc)

	srvr := sv.Server{
		ServeMux: &http.ServeMux{},
	}

	sv.Server.ConnectServer(srvr, path, hndlr, addr)
}
