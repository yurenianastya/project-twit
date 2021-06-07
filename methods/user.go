package methods

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/yurenianastya/project-twit/proto/user"
	"github.com/yurenianastya/project-twit/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type User struct {
	user.UnimplementedUserServiceServer
}

func (u *User) CreateUser(c context.Context, input *user.User) (*user.UserResponse, error) {
	userCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("USER",".."))
	_, err := userCollection.InsertOne(MongoCtx, bson.M{
		"_id": uuid.New().String(),
		"username": input.Username,
		"password": input.Password,
		"email": input.Email,
	})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	userStatus := &user.UserResponse{Value: "Successfully added new user"}
	return userStatus, nil
}

func (u *User) GetUser(c context.Context, id *user.UserUUID) (*user.User, error) {
	userCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("USER",".."))
	decodedUser := utils.MongoDecodedUser{}
	result := userCollection.FindOne(MongoCtx, bson.M{"_id": id.Value})
	if err := result.Decode(&decodedUser); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not find twit with Object ID %s: %v", id.Value, err))
	}
	Id := &user.UserUUID{Value: decodedUser.ID}
	response := &user.User{
		Id: Id,
		Username: decodedUser.Username,
		Password: decodedUser.Password,
		Email: decodedUser.Email,
	}
	return response, nil
}

func (u *User) GetUsers(n *emptypb.Empty, stream user.UserService_GetUsersServer) error {
	userCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("USER",".."))
	decodedUser := utils.MongoDecodedUser{}
	cursor, err := userCollection.Find(context.Background(), bson.M{})
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
		err := cursor.Decode(&decodedUser)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		userId := &user.UserUUID{Value: decodedUser.ID}
		response := &user.User{
			Id: userId,
			Username: decodedUser.Username,
			Password: decodedUser.Password,
			Email: decodedUser.Email,
		}
		err = stream.Send(response)
		if err != nil {
			return err
		}
	}
	return nil
}

func (u *User) UpdateUser(c context.Context, newUser *user.User) (*user.User, error) {
	userCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("USER",".."))
	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": newUser.Id.Value}
	_, err := userCollection.UpdateOne(MongoCtx, filter,
		bson.M{ "$set": bson.M {
			"username": newUser.Username,
			"password": newUser.Password,
			"email": newUser.Email,
		}}, opts)
	if err != nil {
		return nil, err
	}
	response := &user.User{
		Id: newUser.Id,
		Username: newUser.Username,
		Password: newUser.Password,
		Email: newUser.Email,
	}
	return response, nil
}

func (u *User) DeleteUser(c context.Context, userid *user.UserUUID) (*user.UserResponse, error) {
	userCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("USER",".."))
	reactedTwitsCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("USER_REACTED_TWITS",".."))
	twitCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("TWIT_COLLECTION",".."))
	listCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("CUSTOM_LISTS",".."))
	_, err := userCollection.DeleteOne(MongoCtx, bson.M{"_id": userid.Value})
	if err != nil {
		return nil, err
	}
	_, err = twitCollection.DeleteMany(MongoCtx,bson.M{"user_id": userid.Value})
	if err != nil {
		return nil, err
	}
	_, err = reactedTwitsCollection.DeleteMany(MongoCtx, bson.M{"user_id": userid.Value})
	if err != nil {
		return nil, err
	}
	_, err = listCollection.UpdateOne(MongoCtx,bson.M{"_id": userid.Value},
		bson.M{"$pull": bson.M{"users": userid.Value}})
	if err != nil {
		return nil, err
	}
	_, err = listCollection.DeleteMany(MongoCtx, bson.M{"creator_id": userid.Value})
	if err != nil {
		return nil, err
	}
	response := &user.UserResponse{Value: "User was successfully deleted"}
	return response, nil
}


