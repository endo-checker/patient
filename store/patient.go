package store

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	pb "github.com/endo-checker/patient/internal/gen/patient/v1"
)

type Storer interface {
	AddPatient(ctx context.Context, p *pb.Patient) error
	QueryPatient(ctx context.Context, qr *pb.QueryRequest) ([]*pb.Patient, int64, error)
	GetPatient(ctx context.Context, id string) (*pb.Patient, error)
	UpdatePatient(id string, ctx context.Context, u *pb.Patient) error
	DeletePatient(id string) error
}

func (s Store) AddPatient(ctx context.Context, p *pb.Patient) error {
	_, err := s.locaColl.InsertOne(ctx, p)
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (s Store) QueryPatient(ctx context.Context, qr *pb.QueryRequest) ([]*pb.Patient, int64, error) {
	filter := bson.M{}

	if qr.SearchText != "" {
		filter = bson.M{"$and": bson.A{filter,
			bson.M{"$or": bson.A{
				bson.M{"givennames": primitive.Regex{Pattern: qr.SearchText, Options: "i"}},
				bson.M{"familyname": primitive.Regex{Pattern: qr.SearchText, Options: "i"}}}}}}
	}

	if qr.AuthId != "" {
		filter = bson.M{"$and": bson.A{filter, bson.M{"specialistid": qr.AuthId}}}
	}

	opt := options.FindOptions{
		Skip:  &qr.Offset,
		Limit: &qr.Limit,
		Sort:  bson.M{"risk": -1},
	}

	cursor, err := s.locaColl.Find(ctx, filter, &opt)
	if err != nil {
		return nil, 0, err
	}

	var ptnts []*pb.Patient
	if err := cursor.All(ctx, &ptnts); err != nil {
		return nil, 0, err
	}

	matches, err := s.locaColl.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return ptnts, matches, err
}

func (s Store) GetPatient(ctx context.Context, id string) (*pb.Patient, error) {
	var p pb.Patient

	if err := s.locaColl.FindOne(context.Background(), bson.M{"id": id}).Decode(&p); err != nil {
		if err == mongo.ErrNoDocuments {
			return &p, err
		}
		return &p, err
	}

	return &p, nil
}

func (s Store) UpdatePatient(id string, ctx context.Context, p *pb.Patient) error {
	insertResult, err := s.locaColl.ReplaceOne(ctx, bson.M{"id": id}, p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)

	return err
}

func (s Store) DeletePatient(id string) error {
	if _, err := s.locaColl.DeleteOne(context.Background(), bson.M{"id": id}); err != nil {
		return err
	}
	return nil
}
