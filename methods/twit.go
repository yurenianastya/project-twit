package methods

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"project-twit/proto/twit"
	"project-twit/utils"
)


type Twit struct {
	twit.UnimplementedTwitServiceServer
}

func (t *Twit) WriteTwit(c context.Context, input *twit.Twit) (*twit.ResponseTwit, error) {
	input.Id.Value = uuid.New().String()
	input.Date = timestamppb.Now()
	twitCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("REACTIONS_COLLECTION",".."))
	_, err := twitCollection.InsertOne(MongoCtx, bson.M{"_id": input.Id.Value, "likes": 0,
		"retwits": 0})
	if err != nil {
		return nil, err
	}
	twitCollection = Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("TWIT_COLLECTION",".."))
	_, err = twitCollection.InsertOne(MongoCtx,
		bson.M{"_id": input.Id.Value,
			"date": input.Date,
			"text": input.Text,
			"nickname": input.Nickname})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	twitStatus := &twit.ResponseTwit{Value: "successfully added new twit"}
	return twitStatus, nil
}

func (t *Twit) GetTwit(c context.Context, id *twit.TwitUUID) (*twit.Twit, error) {
	twitCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("TWIT_COLLECTION",".."))
	var data = utils.MongoDecodedTwitData{}
	result := twitCollection.FindOne(MongoCtx, bson.M{"_id": id.Value})
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not find twit with Object ID %s: %v", id.Value, err))
	}
	response := &twit.Twit{
		Id: id,
		Date: data.Date,
		Text: data.Text,
		Nickname: data.Nickname}
	return response, nil
}

func (t *Twit) GetTwits(n *emptypb.Empty, stream twit.TwitService_GetTwitsServer) error {
	twitCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("TWIT_COLLECTION",".."))
	data := utils.MongoDecodedTwitData{}
	cursor, err := twitCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(cursor, context.Background())
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&data)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		twitID := twit.TwitUUID{Value: data.ID}
		response := &twit.Twit{
			Id: &twitID,
			Date: data.Date,
			Text: data.Text,
			Nickname: data.Nickname}
		err = stream.Send(response)
		if err != nil {
			return err
		}
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown cursor error: %v", err))
	}
	return nil
}

func (t *Twit) DeleteTwit(c context.Context, id *twit.TwitUUID) (*twit.ResponseTwit, error) {
	twitCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("REACTIONS_COLLECTION",".."))
	_, err := twitCollection.DeleteOne(MongoCtx, bson.M{"_id": id.Value})
	if err != nil {
		return nil, err
	}
	twitCollection = Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("TWIT_COLLECTION",".."))
	_, err = twitCollection.DeleteOne(MongoCtx, bson.M{"_id": id.Value})
	if err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not delete twit with id %s: %v", id.Value, err))
		}
		twitStatus := &twit.ResponseTwit{Value: "twit was successfully deleted"}
	return twitStatus, nil
}
