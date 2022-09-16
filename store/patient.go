package store

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/metadata"

	pb "github.com/endo-checker/patient/gen/proto/go/patient/v1"
)

type Storer interface {
	AddPatient(u *pb.Patient, md metadata.MD) error
	QueryPatient(qr *pb.QueryRequest, md metadata.MD) ([]*pb.Patient, int64, error)
	GetPatient(id string, md metadata.MD) (*pb.Patient, error)
	UpdatePatient(id string, md metadata.MD, u *pb.Patient) error
	DeletePatient(id string, md metadata.MD) error
}

func (s Store) AddPatient(p *pb.Patient, md metadata.MD) error {
	_, err := s.locaColl.InsertOne(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (s Store) QueryPatient(qr *pb.QueryRequest, md metadata.MD) ([]*pb.Patient, int64, error) {
	filter := bson.M{}

	if qr.SearchText != "" {
		filter = bson.M{"$text": bson.M{"$search": `"` + qr.SearchText + `"`}}
	}

	opt := options.FindOptions{
		Skip:  &qr.Offset,
		Limit: &qr.Limit,
		Sort:  bson.M{"medicaldetails.risk": -1},
	}

	ctx := context.Background()
	cursor, err := s.locaColl.Find(ctx, filter, &opt)
	if err != nil {
		return nil, 0, err
	}

	var ptnts []*pb.Patient
	if err := cursor.All(context.Background(), &ptnts); err != nil {
		return nil, 0, err
	}

	matches, err := s.locaColl.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return ptnts, matches, err
}

func (s Store) GetPatient(id string, md metadata.MD) (*pb.Patient, error) {
	var p pb.Patient

	if err := s.locaColl.FindOne(context.Background(), bson.M{"id": id}).Decode(&p); err != nil {
		if err == mongo.ErrNoDocuments {
			return &p, err
		}
		return &p, err
	}

	return &p, nil
}

func (s Store) UpdatePatient(id string, md metadata.MD, p *pb.Patient) error {
	insertResult, err := s.locaColl.ReplaceOne(context.Background(), bson.M{"id": id}, p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nInserted a Single Document: %v\n", insertResult)

	return err
}

func (s Store) DeletePatient(id string, md metadata.MD) error {
	if _, err := s.locaColl.DeleteOne(context.Background(), bson.M{"id": id}); err != nil {
		return err
	}
	return nil
}
