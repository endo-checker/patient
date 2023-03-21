package handler

import (
	"context"
	"errors"
	"regexp"
	"time"

	rc "github.com/AvraamMavridis/randomcolor"
	"github.com/bufbuild/connect-go"
	"github.com/google/uuid"

	pb "github.com/endo-checker/patient/internal/gen/patient/v1"
	pbcnn "github.com/endo-checker/patient/internal/gen/patient/v1/patientv1connect"
	"github.com/endo-checker/patient/store"
)

type PatientServer struct {
	Store store.Storer
	pbcnn.UnimplementedPatientServiceHandler
}

func (p PatientServer) Create(ctx context.Context, req *connect.Request[pb.CreateRequest]) (*connect.Response[pb.CreateResponse], error) {
	reqMsg := req.Msg
	ptnt := reqMsg.Patient
	ptnt.Id = uuid.NewString()
	ptnt.CreatedAt = time.Now().Unix()
	ptnt.IconColor = rc.GetRandomColorInHex()

	if err := p.Store.Create(ctx, ptnt); err != nil {
		return nil, connect.NewError(connect.CodeAborted, err)
	}

	rsp := &pb.CreateResponse{
		Patient: ptnt,
	}
	return connect.NewResponse(rsp), nil
}

func (p PatientServer) Query(ctx context.Context, req *connect.Request[pb.QueryRequest]) (*connect.Response[pb.QueryResponse], error) {
	reqMsg := req.Msg

	if reqMsg.SearchText != "" {
		pattern, err := regexp.Compile(`^[a-zA-Z@. ]+$`)
		if err != nil {
			return nil, connect.NewError(connect.CodeAborted, err)
		}
		if !pattern.MatchString(reqMsg.SearchText) {
			return nil, connect.NewError(connect.CodeInvalidArgument,
				errors.New("invalid search text format"))
		}
	}

	cur, mat, err := p.Store.Fetch(ctx, reqMsg)
	if err != nil {
		return nil, connect.NewError(connect.CodeAborted, err)
	}

	rsp := &pb.QueryResponse{
		Cursor:  cur,
		Matches: mat,
	}
	return connect.NewResponse(rsp), nil
}

func (p PatientServer) Get(ctx context.Context, req *connect.Request[pb.GetRequest]) (*connect.Response[pb.GetResponse], error) {
	reqMsg := req.Msg

	ptnt, err := p.Store.Get(ctx, reqMsg.PatientId)
	if err != nil {
		return nil, connect.NewError(connect.CodeAborted, err)
	}

	rsp := &pb.GetResponse{Patient: ptnt}
	return connect.NewResponse(rsp), nil
}

func (p PatientServer) Update(ctx context.Context, req *connect.Request[pb.UpdateRequest]) (*connect.Response[pb.UpdateResponse], error) {
	reqMsg := req.Msg

	if err := p.Store.Update(ctx, reqMsg.PatientId, reqMsg.Patient); err != nil {
		return nil, connect.NewError(connect.CodeAborted, err)
	}

	rsp := &pb.UpdateResponse{
		Patient: reqMsg.Patient,
	}
	return connect.NewResponse(rsp), nil

}

func (p PatientServer) Delete(ctx context.Context, req *connect.Request[pb.DeleteRequest]) (*connect.Response[pb.DeleteResponse], error) {
	reqMsg := req.Msg

	if err := p.Store.Delete(ctx, reqMsg.PatientId); err != nil {
		return nil, connect.NewError(connect.CodeAborted, err)
	}

	rsp := &pb.DeleteResponse{}
	return connect.NewResponse(rsp), nil
}