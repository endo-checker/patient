package store

import (
	"context"

	st "github.com/endo-checker/common/store"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	patientv1 "github.com/endo-checker/patient/internal/gen/patient/v1"
)

type Storer interface {
	st.Storer[*patientv1.Patient]
	Fetch(ctx context.Context, qr *patientv1.QueryRequest) ([]*patientv1.Patient, int64, error)
}

type PatientStore struct {
	*st.Store[*patientv1.Patient]
}

func (s PatientStore) Fetch(ctx context.Context, qr *patientv1.QueryRequest) ([]*patientv1.Patient, int64, error) {
	filter := bson.M{}

	if qr.SearchText != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"$or": bson.A{
				bson.M{"givennames": primitive.Regex{Pattern: qr.SearchText, Options: "i"}},
				bson.M{"familyname": primitive.Regex{Pattern: qr.SearchText, Options: "i"}}}}}}
	}

	f := st.WithFilter(filter)

	fo := st.WithFindOptions(options.FindOptions{
		Skip:  &qr.Offset,
		Limit: &qr.Limit,
		Sort:  bson.M{"risk": -1},
	})

	return s.List(ctx, f, fo)
}