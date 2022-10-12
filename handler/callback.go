package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"

	sv "github.com/endo-checker/common/store"
	pbauth "github.com/endo-checker/patient/gen/proto/go/auth/v1"
	pbsub "github.com/endo-checker/patient/gen/proto/go/patient/v1"
)

type CallbackServer struct {
	Server PatientServer
	pb.UnimplementedAppCallbackServer
}

// Dapr will call this method to get the list of topics the app wants to subscribe to.
func (d CallbackServer) ListTopicSubscriptions(ctx context.Context, in *emptypb.Empty) (*pb.ListTopicSubscriptionsResponse, error) {

	return &pb.ListTopicSubscriptionsResponse{
		Subscriptions: []*pb.TopicSubscription{{
			PubsubName: "patient",
			Topic:      "create",
			Routes:     &pb.TopicRoutes{Default: "/create"},
		}},
	}, nil
}

// OnTopicEvent is fired for events subscribed to.
// Dapr sends published messages in a CloudEvents 0.3 envelope.
func (d CallbackServer) OnTopicEvent(ctx context.Context, in *pb.TopicEventRequest) (*pb.TopicEventResponse, error) {
	var p pbsub.Patient
	if err := protojson.Unmarshal(in.Data, &p); err != nil {
		return &pb.TopicEventResponse{Status: pb.TopicEventResponse_DROP},
			status.Errorf(codes.Aborted, "issue unmarshalling data: %v", err)
	}

	switch in.Path {
	case "/create":
		createAuth(ctx, p.GivenNames, p.FamilyName, p.Email, p.Id)
	default:
		return &pb.TopicEventResponse{},
			status.Errorf(codes.Aborted, "unexpected path in OnTopicEvent: %s", in.Path)
	}

	return &pb.TopicEventResponse{}, nil
}

// Creates a new tenant on Auth0
func createAuth(ctx context.Context, givenName, familyName, email, id string) (*pbauth.CreateResponse, error) {
	_, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return &pbauth.CreateResponse{}, status.Errorf(codes.Aborted, "%s", "no incoming context")
	}

	url := sv.LoadEnv("AUTH0_DOMAIN")
	key := sv.LoadEnv("AUTH_CLIENT_ID")

	authUser := &pbauth.AuthUser{
		GivenName:  givenName,
		FamilyName: familyName,
		Email:      email,
		Password:   "Wfbuebf45YYvche",
		Connection: "Username-Password-Authentication",
		ClientId:   key,
		UserMetadata: &pbauth.UserMetadata{
			PatientId: id,
			Role:      "patient",
		},
	}

	json_data, _ := json.Marshal(authUser)

	resp, _ := http.NewRequest("POST", "https://"+url+"/dbconnections/signup", bytes.NewBuffer(json_data))
	resp.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	_, err := client.Do(resp)
	if err != nil {
		return &pbauth.CreateResponse{}, err
	}
	return &pbauth.CreateResponse{AuthUser: authUser}, nil
}
