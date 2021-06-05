package methods

import (
	"context"
	"github.com/magiconair/properties/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"project-twit/proto/twit"
	"project-twit/utils"
	"testing"
	"time"
)

type mockTwitService_GetTwitsServer struct {
	grpc.ServerStream
	Twits []*twit.Twit
}

func (_m *mockTwitService_GetTwitsServer) Send(twit *twit.Twit) error {
	_m.Twits = append(_m.Twits, twit)
	return nil
}

func DatabaseInit() (*mongo.Client, context.Context, func()){
	Client, _ = mongo.NewClient(options.Client().ApplyURI(GetEnvVariable("MONGO_ADDRESS", "..")))
	mongoCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err := Client.Connect(mongoCtx)
	if err != nil {
		log.Fatal(err)
	}
	return Client, mongoCtx, func() {
		cancel()
		err := Client.Disconnect(mongoCtx)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getRandomTwitDocument(collection *mongo.Collection) utils.MongoDecodedTwitData {
	pipeline := []bson.M{{"$sample": bson.M{"size": 1}}}
	cursor, err := collection.Aggregate(context.Background(), pipeline)
	var data utils.MongoDecodedTwitData
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
	}
	return data
}

func TestWriteTwit(t *testing.T) {
	_, ctx, closeCon := DatabaseInit()
	defer closeCon()
	twitStruct := Twit{}
	testCases := []struct{
			Id *twit.TwitUUID
			Date *timestamppb.Timestamp
			Text string
			Nickname string
	} {
		{
			Id: &twit.TwitUUID{Value: ""},
			Date:     nil,
			Text:     "tests",
			Nickname: "test",
		},
		{
			Id: &twit.TwitUUID{Value: ""},
			Date: nil,
			Text: "",
			Nickname: "",
		},
		{
			Id: &twit.TwitUUID{Value: ""},
			Date: nil,
			Text: "very long unnecessary twit text which no one will ever read",
			Nickname: "secret",
		},
	}
	for _, tt := range testCases {
		req := &twit.Twit{
			Id: tt.Id,
			Date: tt.Date,
			Text: tt.Text,
			Nickname: tt.Nickname,
		}
		response, err := twitStruct.WriteTwit(ctx, req)
		if err != nil {
			t.Errorf("Test got unexpected error")
		}
		twitStatus := &twit.ResponseTwit{Value: "successfully added new twit"}
		assert.Equal(t, response.Value, twitStatus.Value, "response texts are not equal")
	}
}

func TestGetTwit(t *testing.T) {
	_, ctx, closeCon := DatabaseInit()
	defer closeCon()
	//getting id that exists from db to test func
	collection := Client.Database(GetEnvVariable("TWIT_DB",".")).
		Collection(GetEnvVariable("REACTIONS_COLLECTION","."))
	data := getRandomTwitDocument(collection)
	twitStruct := Twit{}
	id := twit.TwitUUID{Value: data.ID}
	response, err := twitStruct.GetTwit(ctx, &id)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, response.Text, response.Text, "response text should be the same")
}

func TestDeleteTwit(t *testing.T) {
	_, ctx, closeCon := DatabaseInit()
	defer closeCon()
	//getting id that exists from db to test func
	collection := Client.Database(GetEnvVariable("TWIT_DB",".")).
		Collection(GetEnvVariable("TWIT_COLLECTION","."))
	data := getRandomTwitDocument(collection)
	twitStruct := Twit{}
	id := twit.TwitUUID{Value: data.ID}
	response, err := twitStruct.DeleteTwit(ctx, &id)
	if err != nil {
		log.Fatal(err)
	}
	twitStatus := &twit.ResponseTwit{Value: "twit was successfully deleted"}
	assert.Equal(t, response.Value, twitStatus.Value, "response text should be the same")
}

func TestGetTwits(t *testing.T) {
	_, _, closeCon := DatabaseInit()
	defer closeCon()
	twitStruct := &Twit{}
	stream := &mockTwitService_GetTwitsServer{}
	var n *emptypb.Empty
	response := twitStruct.GetTwits(n, stream)
	if response != nil {
		t.Errorf("Test failed, method returned an error")
	}
}
