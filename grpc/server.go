package grpc

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"project-twit/proto/reactions"
	"project-twit/proto/twit"
	"time"
)


var client *mongo.Client
var mongoCtx context.Context

func GetEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func Configure() {
	lis, err := net.Listen("tcp", GetEnvVariable("SERVER_PORT"))
	if err != nil {
		fmt.Println("failed to listen: ", err)
		return
	}
	// Creates a new gRPC Server
	server := grpc.NewServer()
	twit.RegisterTwitServiceServer(server, &Twit{})
	reactions.RegisterLikeServiceServer(server, &Like{})
	reactions.RegisterRetwitServiceServer(server, &Retwit{})
	reflection.Register(server)
	// connection to mongodb
	client, _ = mongo.NewClient(options.Client().ApplyURI(GetEnvVariable("MONGO_ADDRESS")))
	mongoCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(mongoCtx)
	if err != nil {
		log.Fatal(err)
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(client, mongoCtx)
	// start listening rpc server
	err = server.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
