package methods

import (
	"context"
	"fmt"
	"github.com/yurenianastya/project-twit/proto/account"
	"github.com/yurenianastya/project-twit/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Account struct {
	account.UnimplementedAccountServiceServer
}

type streamType interface {
	account.AccountService_GetRetwitedTwitsServer
	account.AccountService_GetLikedTwitsServer
}

func getReactedTwits(action string, Username *account.AccountUUID, stream streamType) error {
	idCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("USER_REACTED_TWITS",".."))
	twitCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("TWIT_COLLECTION",".."))
	userCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("USER",".."))
	var retrievedTwitIds []string
	decodedIds := utils.MongoDecodedTwitIds{}
	data := utils.MongoDecodedTwitData{}
	usernameData := utils.MongoDecodedUsernames{}
	idCursor, err := idCollection.Find(context.Background(), bson.M{
		"user_id": Username.Value,
		"action": action,
	})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	for idCursor.Next(context.Background()) {
		err := idCursor.Decode(&decodedIds)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		retrievedTwitIds = append(retrievedTwitIds, decodedIds.TwitId)
	}
	err = idCursor.Close(context.Background())
	if err != nil {
		return err
	}
	if retrievedTwitIds == nil {
		return nil
	}
	twitCursor, err := twitCollection.Find(context.Background(), bson.M{"_id": bson.M{"$in":retrievedTwitIds}})
	if err != nil {
		return err
	}
	for twitCursor.Next(context.Background()) {
		err = twitCursor.Decode(&data)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		id := &account.AccountTwitUUID{Value: data.ID}
		result := userCollection.FindOne(MongoCtx, bson.M{"_id": data.User})
		if err := result.Decode(&usernameData); err != nil {
			return status.Errorf(codes.NotFound,
				fmt.Sprintf("Could not find user with Object ID %s: %v", data.User, err))
		}
		response := &account.AccountTwit{
			Id: id,
			Date: data.Date,
			Text: data.Text,
			Nickname: usernameData.Username,
		}
		err = stream.Send(response)
		if err != nil {
			return err
		}
	}
	if err := twitCursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown twitCursor error: %v", err))
	}
	return nil
}

func (a *Account) GetLikedTwits(userId *account.AccountUUID, stream account.AccountService_GetLikedTwitsServer) error {
	err := getReactedTwits("like", userId, stream)
	if err != nil {
		return err
	}
	return nil
}

func (a *Account) GetRetwitedTwits(userId *account.AccountUUID, stream account.AccountService_GetRetwitedTwitsServer) error {
	err := getReactedTwits("retwit", userId, stream)
	if err != nil {
		return err
	}
	return nil
}

func (a *Account) GetUserTwits(userId *account.AccountUUID, stream account.AccountService_GetUserTwitsServer) error {
	twitCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("TWIT_COLLECTION",".."))
	userCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("USER",".."))
	data := utils.MongoDecodedTwitData{}
	usernameData := utils.MongoDecodedUsernames{}
	cursor, err := twitCollection.Find(context.Background(), bson.M{"user_id": userId.Value})
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
		result := userCollection.FindOne(MongoCtx, bson.M{"_id": data.User})
		if err := result.Decode(&usernameData); err != nil {
			return status.Errorf(codes.NotFound,
				fmt.Sprintf("Could not find user with Object ID %s: %v", data.User, err))
		}
		twitID := account.AccountTwitUUID{Value: data.ID}
		protoData := &account.AccountTwit{
			Id: &twitID,
			Date: data.Date,
			Text: data.Text,
			Nickname: usernameData.Username}
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


