package store

import (
	st "github.com/endo-checker/common/store"

	patientv1 "github.com/endo-checker/patient/internal/gen/patient/v1"
)

type Storer interface {
	st.Storer[*patientv1.Patient]
	// QueryPatient(ctx context.Context, qr *patientv1.QueryRequest) ([]*patientv1.Patient, int64, error)
}

type PatientStore struct {
	*st.Store[*patientv1.Patient]
}

// func (s Storer) QueryPatient(ctx context.Context, qr *patientv1.QueryRequest) ([]*patientv1.Patient, int64, error) {
// 	filter := bson.M{}

// 	if qr.SearchText != "" {
// 		filter = bson.M{"$and": bson.A{filter,
// 			bson.M{"$or": bson.A{
// 				bson.M{"givennames": primitive.Regex{Pattern: qr.SearchText, Options: "i"}},
// 				bson.M{"familyname": primitive.Regex{Pattern: qr.SearchText, Options: "i"}}}}}}
// 	}

// 	if qr.AuthId != "" {
// 		filter = bson.M{"$and": bson.A{filter, bson.M{"specialistid": qr.AuthId}}}
// 	}

// 	opt := options.FindOptions{
// 		Skip:  &qr.Offset,
// 		Limit: &qr.Limit,
// 		Sort:  bson.M{"risk": -1},
// 	}

// 	cursor, err := s.locaColl.Find(ctx, filter, &opt)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	var ptnts []*patientv1.Patient
// 	if err := cursor.All(ctx, &ptnts); err != nil {
// 		return nil, 0, err
// 	}

// 	matches, err := s.locaColl.CountDocuments(ctx, filter)
// 	if err != nil {
// 		return nil, 0, err
// 	}

// 	return ptnts, matches, err
// }
