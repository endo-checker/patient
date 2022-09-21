package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	pb "github.com/dapr/dapr/pkg/proto/runtime/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/emptypb"

	pbsub "github.com/endo-checker/patient/gen/proto/go/patient/v1"
	"github.com/endo-checker/patient/store"
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
		createAuthUser(p.GivenNames, p.FamilyName, p.Email, p.Id)
	default:
		return &pb.TopicEventResponse{},
			status.Errorf(codes.Aborted, "unexpected path in OnTopicEvent: %s", in.Path)
	}

	return &pb.TopicEventResponse{}, nil
}

// creates a new tenant on Auth0
func createAuthUser(givenName, familyName, email, id string) {

	url := store.LoadEnv("AUTH0_DOMAIN")
	key := store.LoadEnv("AUTH_CLIENT_ID")

	data := map[string]string{
		"patient_id": id,
	}

	values := map[string]interface{}{
		"given_name":    givenName,
		"family_name":   familyName,
		"email":         email,
		"password":      "Wfbuebf45YYvche",
		"connection":    "Username-Password-Authentication",
		"client_id":     key,
		"user_metadata": data,
	}

	json_data, _ := json.Marshal(values)

	req, _ := http.NewRequest("POST", "https://"+url+"/dbconnections/signup", bytes.NewBuffer(json_data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}
