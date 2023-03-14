package main

import (
	"net/http"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/endo-checker/patient/handler"
	pbcnn "github.com/endo-checker/patient/internal/gen/patient/v1/patientv1connect"
	"github.com/endo-checker/patient/store"
	"github.com/rs/cors"
)

const port = "localhost:8080"

func main() {
	s := store.Connect()

	svc := &handler.PatientServer{
		Store: s,
	}

	c := setCORS()

	mux := http.NewServeMux()
	path, h := pbcnn.NewPatientServiceHandler(svc)
	mux.Handle(path, h)
	handler := c.Handler(mux)

	http.ListenAndServe(
		port,
		h2c.NewHandler(handler, &http2.Server{}),
	)
}

func setCORS() *cors.Cors {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Content-Type"},
		AllowedMethods:   []string{"POST", "OPTIONS"},
		AllowCredentials: true,
	})
	return c
}
