package methods

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"project-twit/proto/lists"
	"project-twit/utils"
	"strings"
)

type List struct {
	lists.UnimplementedListServiceServer
}

func (l *List) CreateCustomList(c context.Context, name *lists.ListName) (*lists.ListResponse, error) {
	listCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("CUSTOM_LISTS",".."))
	_, err := listCollection.InsertOne(MongoCtx, bson.M{
		"_id": uuid.New().String(),
		"name": name.Value,
		"users": []string{},
	})
	if err != nil {
		return nil, status.Errorf( codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	response := &lists.ListResponse{Value: "Created a custom list"}
	return response, nil
}

func (l *List) GetListUsers(c context.Context, id *lists.ListUUID) (*lists.ListResponse, error) {
	var userList []string
	data := utils.MongoDecodedList{}
	listCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("CUSTOM_LISTS",".."))
	result := listCollection.FindOne(MongoCtx, bson.M{"_id": id.Value})
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not find twit with Object ID %s: %v", id.Value, err))
	}
	userList = append(userList, data.Usernames...)
	response := &lists.ListResponse{Value: strings.Join(userList,",")}
	return response, nil
}

func (l *List) AddUserToCustomList(c context.Context, user *lists.User) (*lists.ListResponse, error) {
	listCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("CUSTOM_LISTS",".."))
	_, err := listCollection.UpdateOne(MongoCtx, bson.M{"_id": user.ListUUID},
	bson.M{"$push": bson.M{"users": user.Username}})
	if err != nil {
		return nil, err
	}
	response := &lists.ListResponse{Value: "Added user to a list"}
	return response, nil
}

func (l *List) RemoveUserFromCustomList(c context.Context, user *lists.User) (*lists.ListResponse, error) {
	listCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("CUSTOM_LISTS",".."))
	_, err := listCollection.UpdateOne(MongoCtx, bson.M{"_id": user.ListUUID},
		bson.M{"$pull": bson.M{"users": user.Username}})
	if err != nil {
		return nil, err
	}
	response := &lists.ListResponse{Value: "Removed user from list"}
	return response, nil
}

func (l *List) GetUsersTwitsFromCustomList(id *lists.ListUUID, stream lists.ListService_GetUsersTwitsFromCustomListServer) error {
	listCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("CUSTOM_LISTS",".."))
	namesData := utils.MongoDecodedNames{}
	twitData := utils.MongoDecodedTwitData{}
	result := listCollection.FindOne(MongoCtx, bson.M{"_id": id.Value})
	if err := result.Decode(&namesData); err != nil {
		return status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not find twit with Object ID %s: %v", id.Value, err))
	}
	twitCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("TWIT_COLLECTION",".."))
	cursor, err := twitCollection.Find(MongoCtx, bson.M{"nickname": bson.M{"$in":namesData.Name}})
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
		err := cursor.Decode(&twitData)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		twitID := lists.ListUUID{Value: twitData.ID}
		response := &lists.ListTwit{
			Id: &twitID,
			Date: twitData.Date,
			Text: twitData.Text,
			Nickname: twitData.Nickname}
		err = stream.Send(response)
		if err != nil {
			return err
		}
	}
	return nil
}

func (l *List) DeleteCustomList(c context.Context, id *lists.ListUUID) (*lists.ListResponse, error) {
	listCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("CUSTOM_LISTS",".."))
	_, err := listCollection.DeleteOne(MongoCtx, bson.M{"_id": id.Value})
	if err != nil {
		return nil, status.Errorf( codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	response := &lists.ListResponse{Value: "Deleted a custom list"}
	return response, nil
}