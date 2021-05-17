package grpc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
	"project-twit/proto"
	tw "project-twit/proto"
	"time"
)


var db *mongo.Client
var mongoCtx context.Context
var twitdb *mongo.Collection

const port = ":9000"

type mongoData struct {
	Id       string                 `json:"_id" bson:"_id"`
	Date     *timestamppb.Timestamp `json:"date" bson:"date"`
	Text     string                 `json:"text" bson:"text"`
	Nickname string                 `json:"nickname" bson:"nickname"`
}

type Twit struct {
	twitsList [] *proto.Twit
	tw.UnimplementedTwitServiceServer
}

func (t *Twit) WriteTwit(c context.Context, input *proto.Twit) (*proto.Twit, error) {
	input.Id.Value = uuid.New().String()
	input.Date = timestamppb.Now()
	_, err := twitdb.InsertOne(mongoCtx, bson.M{"_id": input.Id.Value, "date": input.Date,
		"text": input.Text, "nickname": input.Nickname})
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return &proto.Twit{Id: input.Id, Date: input.Date, Text:
	input.Text, Nickname: input.Nickname}, nil
}

func (t *Twit) GetTwit(c context.Context, Id *proto.UUID) (*proto.Twit, error) {
	var data = mongoData{}
	result := twitdb.FindOne(mongoCtx, bson.M{"_id": Id.Value})
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not find blog with Object ID %s: %v", Id.Value, err))
	}
	response := &proto.Twit{Id: Id, Date: data.Date, Text: data.Text, Nickname: data.Nickname}
	return response, nil
}

func (t *Twit) GetTwits(n *emptypb.Empty, stream proto.TwitService_GetTwitsServer) error {
	data := &mongoData{}
	cursor, err := twitdb.Find(context.Background(), bson.M{})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&data)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		twitId := proto.UUID{Value: data.Id}
		protoData := &proto.Twit{Id: &twitId, Date: data.Date, Text: data.Text, Nickname: data.Nickname}
		stream.Send(protoData)
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	//for _, elem := range t.twitsList {
	//	err := stream.Send(elem)
	//	if err != nil {
	//		return err
	//	}
	//}
	return nil
}

func (t *Twit) DeleteTwit(c context.Context, id *proto.UUID) (*proto.Twit, error) {
	_, err := twitdb.DeleteOne(mongoCtx, bson.M{"_id": id.Value})
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Could not delete blog with id %s: %v", id.Value, err))
		}
	return &proto.Twit{Id: nil, Date: nil, Text: "", Nickname: ""}, nil
}

func Server() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("failed to listen: ", err)
		return
	}
	// Creates a new gRPC TwitServer
	s := grpc.NewServer()
	tw.RegisterTwitServiceServer(s, &Twit{})
	reflection.Register(s)
	fmt.Println("connecting to mongodb")
	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	mongoCtx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = db.Connect(mongoCtx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Disconnect(mongoCtx)
	// check connection by printing all db names
	databases, err := db.ListDatabaseNames(mongoCtx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(databases)
	twitdb = db.Database("twit").Collection("twits")
	// start listening rpc server
	err = s.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}

}