package grpc

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
	"project-twit/proto/twit"
)


var twitcol *mongo.Collection


type Twit struct {
	twit.UnimplementedTwitServiceServer
}

type mongoTwitData struct {
	ID       string                 `json:"_id" bson:"_id"`
	Date     *timestamppb.Timestamp `json:"date" bson:"date"`
	Text     string                 `json:"text" bson:"text"`
	Nickname string                 `json:"nickname" bson:"nickname"`
}

func (t *Twit) WriteTwit(c context.Context, input *twit.Twit) (*twit.Twit, error) {
	input.Id.Value = uuid.New().String()
	input.Date = timestamppb.Now()
	twitcol = client.Database("twit").Collection("reactions")
	_, err := twitcol.InsertOne(mongoCtx, bson.M{"_id": input.Id.Value, "likes": 0,
		"retwits": 0})
	if err != nil {
		return nil, err
	}
	twitcol = client.Database("twit").Collection("twits")
	_, err = twitcol.InsertOne(mongoCtx, bson.M{"_id": input.Id.Value, "date": input.Date,
		"text": input.Text, "nickname": input.Nickname})
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return &twit.Twit{Id: input.Id, Date: input.Date, Text:
	input.Text, Nickname: input.Nickname}, nil
}

func (t *Twit) GetTwit(c context.Context, id *twit.TwitUUID) (*twit.Twit, error) {
	twitcol = client.Database("twit").Collection("twits")
	var data = mongoTwitData{}
	result := twitcol.FindOne(mongoCtx, bson.M{"_id": id.Value})
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not find twit with Object ID %s: %v", id.Value, err))
	}
	response := &twit.Twit{Id: id, Date: data.Date, Text: data.Text, Nickname: data.Nickname}
	return response, nil
}

func (t *Twit) GetTwits(n *emptypb.Empty, stream twit.TwitService_GetTwitsServer) error {
	twitcol = client.Database("twit").Collection("twits")
	data := mongoTwitData{}
	cursor, err := twitcol.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&data)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		twitID := twit.TwitUUID{Value: data.ID}
		protoData := &twit.Twit{Id: &twitID, Date: data.Date, Text: data.Text, Nickname: data.Nickname}
		err = stream.Send(protoData)
		if err != nil {
			return err
		}
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown cursor error: %v", err))
	}
	return nil
}

func (t *Twit) DeleteTwit(c context.Context, id *twit.TwitUUID) (*twit.Twit, error) {
	twitcol = client.Database("twit").Collection("reactions")
	_, err := twitcol.DeleteOne(mongoCtx, bson.M{"_id": id.Value})
	if err != nil {
		return nil, err
	}
	twitcol = client.Database("twit").Collection("twits")
	_, err = twitcol.DeleteOne(mongoCtx, bson.M{"_id": id.Value})
	if err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not delete twit with id %s: %v", id.Value, err))
		}
	return &twit.Twit{Id: nil, Date: nil, Text: "", Nickname: ""}, nil
}
