package methods

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"project-twit/proto/account"
	"project-twit/utils"
)

type Account struct {
	account.UnimplementedAccountServiceServer
}

type streamType interface {
	account.AccountService_GetRetwitedTwitsServer
	account.AccountService_GetLikedTwitsServer
}

func getReactedTwits(action string, Username *account.Username, stream streamType) error {
	var retrievedTwitIds []string
	decodedIds := utils.MongoDecodedIds{}
	data := utils.MongoDecodedTwitData{}
	idCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("USER_REACTED_COLLECTION",".."))
	idCursor, err := idCollection.Find(context.Background(), bson.M{
		"username": Username.Value,
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
	twitCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("TWIT_COLLECTION",".."))
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
		response := &account.AccountTwit{
			Id: id,
			Date: data.Date,
			Text: data.Text,
			Nickname: data.Nickname,
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

func (a *Account) GetLikedTwits(Username *account.Username, stream account.AccountService_GetLikedTwitsServer) error {
	err := getReactedTwits("like", Username, stream)
	if err != nil {
		return err
	}
	return nil
}

func (a *Account) GetRetwitedTwits(Username *account.Username, stream account.AccountService_GetRetwitedTwitsServer) error {
	err := getReactedTwits("retwit", Username, stream)
	if err != nil {
		return err
	}
	return nil
}

func (a *Account) GetUserTwits(Username *account.Username, stream account.AccountService_GetUserTwitsServer) error {
	twitCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("TWIT_COLLECTION",".."))
	data := utils.MongoDecodedTwitData{}
	cursor, err := twitCollection.Find(context.Background(), bson.M{"nickname": Username.Value})
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
		twitID := account.AccountTwitUUID{Value: data.ID}
		protoData := &account.AccountTwit{
			Id: &twitID,
			Date: data.Date,
			Text: data.Text,
			Nickname: data.Nickname}
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


