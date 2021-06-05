package methods

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"project-twit/proto/reactions"
	"project-twit/utils"
)

type Reaction struct {
	reactions.UnimplementedReactionServiceServer
}

func (r *Reaction) GetTwitReactions(c context.Context, id *reactions.ReactionUUID) (*reactions.ActionResult, error) {
	reactionsCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("REACTIONS_COLLECTION",".."))
	data := utils.MongoReactionsData{}
	result := reactionsCollection.FindOne(MongoCtx, bson.M{"_id": id.Value})
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not find twit with Object ID %s: %v", id.Value, err))
	}
	response := &reactions.ActionResult{
		TwitId: id,
		LikeCounter: data.LikeCounter,
		RetwitCounter: data.RetwitCounter,
	}
	return response, nil
}

func (r *Reaction) ReactToTwit(c context.Context, action *reactions.UsersAction) (*reactions.ResponseReaction, error) {
	reactedIdCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("USER_REACTED_COLLECTION",".."))
	_, err := reactedIdCollection.InsertOne(MongoCtx, bson.M{
		"_id": uuid.New().String(),
		"twit_id": action.TwitId.Value,
		"username": action.User,
		"action": action.Action,
	})
	if err != nil {
		return nil, err
	}
	reactionsCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("REACTIONS_COLLECTION",".."))
	opts := options.Update().SetUpsert(true)
	switch action.Action {
	case "like":
		_, err := reactionsCollection.UpdateOne(MongoCtx, bson.M{"_id": action.TwitId.Value},
			bson.M{"$inc": bson.M{"likes": 1}}, opts)
		if err != nil {
			return nil, err
		}
	case "retwit":
		_, err := reactionsCollection.UpdateOne(MongoCtx, bson.M{"_id": action.TwitId.Value},
			bson.M{"$inc": bson.M{"retwits": 1}}, opts)
		if err != nil {
			return nil, err
		}
	}
	response := &reactions.ResponseReaction{Value: "Reacted to twit successfully"}
	return response, nil
}

func (r *Reaction) UnreactToTwit(c context.Context, action *reactions.UsersAction) (*reactions.ResponseReaction, error) {
	reactedIdCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("USER_REACTED_COLLECTION",".."))
	_, err := reactedIdCollection.DeleteOne(MongoCtx, bson.M{
		"twit_id": action.TwitId.Value,
	})
	if err != nil {
		return nil, err
	}
	reactionsCollection := Client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("REACTIONS_COLLECTION",".."))
	opts := options.Update().SetUpsert(true)
	switch action.Action {
	case "unlike":
		_, err := reactionsCollection.UpdateOne(MongoCtx, bson.M{"_id": action.TwitId.Value},
			bson.M{"$inc": bson.M{"likes": -1}}, opts)
		if err != nil {
			return nil, err
		}
	case "unretwit":
		_, err := reactionsCollection.UpdateOne(MongoCtx, bson.M{"_id": action.TwitId.Value},
			bson.M{"$inc": bson.M{"retwits": -1}}, opts)
		if err != nil {
			return nil, err
		}
	}
	response := &reactions.ResponseReaction{Value: "Unreacted to twit successfully"}
	return response, nil
}