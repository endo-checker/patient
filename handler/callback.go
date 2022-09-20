package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
			PubsubName: "pubsubsrv",
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
		createAuthUser(p.Email, p.GivenNames, p.Id)
	default:
		return &pb.TopicEventResponse{},
			status.Errorf(codes.Aborted, "unexpected path in OnTopicEvent: %s", in.Path)
	}

	return &pb.TopicEventResponse{}, nil
}

func createAuthUser(email, nickname, id string) {

	data := map[string]string{
		"client_id": id,
	}

	values := map[string]interface{}{
		"nickname":      nickname,
		"email":         email,
		"password":      "Wfbuebf45YYvche",
		"connection":    "Username-Password-Authentication",
		"picture":       "http://example.org/jdoe.png",
		"client_id":     "JKv5m5C6LYyIaObQMvgJ8tn4fBnvHDFR",
		"user_metadata": data,
	}

	json_data, _ := json.Marshal(values)

	url := store.LoadEnv("AUTH0_DOMAIN")
	req, _ := http.NewRequest("POST", "https://"+url+"/dbconnections/signup", bytes.NewBuffer(json_data))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
