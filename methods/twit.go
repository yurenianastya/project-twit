package methods

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/yurenianastya/project-twit/proto/twit"
	"github.com/yurenianastya/project-twit/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)


type Twit struct {
	twit.UnimplementedTwitServiceServer
}

func (t *Twit) WriteTwit(c context.Context, input *twit.Twit) (*twit.ResponseTwit, error) {
	input.Id.Value = uuid.New().String()
	input.Date = timestamppb.Now()
	reactCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("REACTIONS_COLLECTION",".."))
	_, err := reactCollection.InsertOne(MongoCtx, bson.M{"_id": input.Id.Value, "likes": 0,
		"retwits": 0})
	if err != nil {
		return nil, err
	}
	twitCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("TWIT_COLLECTION",".."))
	_, err = twitCollection.InsertOne(MongoCtx,
		bson.M{"_id": input.Id.Value,
			"date": input.Date,
			"text": input.Text,
			"user_id": input.User})
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
	userCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("USER",".."))
	var data = utils.MongoDecodedTwitData{}
	var usernameData = utils.MongoDecodedUsernames{}
	result := twitCollection.FindOne(MongoCtx, bson.M{"_id": id.Value})
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not find twit with Object ID %s: %v", id.Value, err))
	}
	result = userCollection.FindOne(MongoCtx, bson.M{"_id": data.User})
	if err := result.Decode(&usernameData); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not find user with Object ID %s: %v", id.Value, err))
	}
	response := &twit.Twit{
		Id: id,
		Date: data.Date,
		Text: data.Text,
		User: usernameData.Username}
	return response, nil
}

func (t *Twit) GetTwits(n *emptypb.Empty, stream twit.TwitService_GetTwitsServer) error {
	userCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("USER",".."))
	twitCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("TWIT_COLLECTION",".."))
	data := utils.MongoDecodedTwitData{}
	usernameData := utils.MongoDecodedUsernames{}
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
		result := userCollection.FindOne(MongoCtx, bson.M{"_id": data.User})
		if err := result.Decode(&usernameData); err != nil {
			return status.Errorf(codes.NotFound,
				fmt.Sprintf("Could not find user with Object ID %s: %v", data.User, err))
		}
		response := &twit.Twit{
			Id: &twitID,
			Date: data.Date,
			Text: data.Text,
			User: usernameData.Username}
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
	reactCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("REACTIONS_COLLECTION",".."))
	twitCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("TWIT_COLLECTION",".."))
	_, err := twitCollection.DeleteOne(MongoCtx, bson.M{"_id": id.Value})
	if err != nil {
		return nil, err
	}
	_, err = reactCollection.DeleteOne(MongoCtx, bson.M{"_id": id.Value})
	if err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not delete twit with id %s: %v", id.Value, err))
		}
		twitStatus := &twit.ResponseTwit{Value: "twit was successfully deleted"}
	return twitStatus, nil
}
