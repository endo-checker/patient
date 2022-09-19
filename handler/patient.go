package handler

import (
	"context"
	"time"

	rc "github.com/AvraamMavridis/randomcolor"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	pb "github.com/endo-checker/patient/gen/proto/go/patient/v1"
	"github.com/endo-checker/patient/store"
)

type PatientServer struct {
	Store store.Storer
	pb.UnimplementedPatientServiceServer
}

func (p PatientServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.CreateResponse{}, status.Errorf(codes.Aborted, "%s", "no incoming context")
	}

	ptnt := req.Patient
	ptnt.Id = uuid.NewString()
	ptnt.CreatedAt = time.Now().Unix()
	ptnt.IconColor = rc.GetRandomColorInHex()

	if err := p.Store.AddPatient(ptnt, md); err != nil {
		return &pb.CreateResponse{}, status.Errorf(codes.Aborted, "%v", err)
	}
	return &pb.CreateResponse{Patient: ptnt}, nil
}

func (p PatientServer) Query(ctx context.Context, req *pb.QueryRequest) (*pb.QueryResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.QueryResponse{}, status.Errorf(codes.Aborted, "%s", "no incoming context")
	}

	cur, mat, err := p.Store.QueryPatient(req, md)
	if err != nil {
		return &pb.QueryResponse{}, status.Errorf(codes.Aborted, "%v", err)
	}

	return &pb.QueryResponse{
		Cursor:  cur,
		Matches: mat,
	}, nil
}

func (p PatientServer) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.GetResponse{}, status.Errorf(codes.Aborted, "%s", "no incoming context")
	}

	ptnt, err := p.Store.GetPatient(req.Id, md)
	if err != nil {
		return &pb.GetResponse{}, status.Errorf(codes.Aborted, "%v", err)
	}

	return &pb.GetResponse{Patient: ptnt}, nil
}

func (p PatientServer) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.UpdateResponse{}, status.Errorf(codes.Aborted, "%s", "no incoming context")
	}

	ptnt := req.Patient
	id := ptnt.Id

	if err := p.Store.UpdatePatient(id, md, ptnt); err != nil {
		return &pb.UpdateResponse{}, status.Errorf(codes.Aborted, "%v", err)
	}

	return &pb.UpdateResponse{Patient: ptnt}, nil
}

func (p PatientServer) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pb.DeleteResponse{}, status.Errorf(codes.Aborted, "%s", "no incoming context")
	}

	if err := p.Store.DeletePatient(req.Id, md); err != nil {
		return &pb.DeleteResponse{}, status.Errorf(codes.Aborted, "%v", err)
	}

	return &pb.DeleteResponse{}, nil
}
