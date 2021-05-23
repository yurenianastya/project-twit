package grpc

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"project-twit/proto/reactions"
	"project-twit/proto/twit"
	"time"
)

var client *mongo.Client
var mongoCtx context.Context

func Server() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("failed to listen: ", err)
		return
	}
	// Creates a new gRPC TwitServer
	s := grpc.NewServer()
	twit.RegisterTwitServiceServer(s, &Twit{})
	reactions.RegisterLikeServiceServer(s, &Like{})
	reactions.RegisterRetwitServiceServer(s, &Retwit{})
	reflection.Register(s)
	// connection to mongodb
	client, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	mongoCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(mongoCtx)
	if err != nil {
		return
	}
	err = client.Disconnect(mongoCtx)
	if err != nil {
		return
	}
	// start listening rpc server
	err = s.Serve(lis)
	if err != nil {
		return
	}
}
