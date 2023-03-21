package store

import (
	st "github.com/endo-checker/common/store"
	patientv1 "github.com/endo-checker/patient/internal/gen/patient/v1"
)

func New(uri string) PatientStore {
	s := st.Connect[*patientv1.Patient](uri)

	return PatientStore{
		Store: &s,
	}
}