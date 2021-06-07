package methods

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/yurenianastya/project-twit/proto/account"
	"github.com/yurenianastya/project-twit/proto/lists"
	"github.com/yurenianastya/project-twit/proto/reactions"
	"github.com/yurenianastya/project-twit/proto/twit"
	"github.com/yurenianastya/project-twit/proto/user"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)


var Client *mongo.Client
var MongoCtx context.Context

func GetEnvVariable(key string, path string) string {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error while reading app.env file %s\n", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		fmt.Println("Invalid type assertion")
	}
	return value
}

func Configure() {
	listener, err := net.Listen("tcp",
		GetEnvVariable("SERVER_PORT", "."))
	if err != nil {
		fmt.Println("failed to listen: ", err)
		return
	}
	// Creates a new gRPC Server
	server := grpc.NewServer()
	twit.RegisterTwitServiceServer(server, &Twit{})
	reactions.RegisterReactionServiceServer(server, &Reaction{})
	lists.RegisterListServiceServer(server, &List{})
	account.RegisterAccountServiceServer(server, &Account{})
	user.RegisterUserServiceServer(server, &User{})
	reflection.Register(server)
	// connection to mongodb
	Client, _ = mongo.NewClient(options.Client().
		ApplyURI(GetEnvVariable("MONGO_ADDRESS",".")))
	MongoCtx, cancel := context.
		WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = Client.Connect(MongoCtx)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(client *mongo.Client, ctx context.Context) {
		err := client.Disconnect(ctx)
		if err != nil {
			fmt.Println(err)
			return
		}
	}(Client, MongoCtx)
	// start listening rpc server
	fmt.Println("Starting server...")
	err = server.Serve(listener)
	if err != nil {
		fmt.Println(err)
		return
	}
}
