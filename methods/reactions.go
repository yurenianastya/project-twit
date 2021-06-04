package methods

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"project-twit/proto/reactions"
)

const increment = "increment"
const decrement = "decrement"

// struct for decoding reactions data from mongodb, has the same fields as in database
type mongoReactionsData struct {
	ID            string     `json:"_id" bson:"_id"`
	LikeCounter   int32      `json:"likes" bson:"likes"`
	RetwitCounter int32      `json:"retwits" bson:"retwits"`
}

type Like struct {
	reactions.UnimplementedLikeServiceServer
}

type Retwit struct {
	reactions.UnimplementedRetwitServiceServer
}

func get(id *reactions.ReactionUUID) mongoReactionsData {
	reactionCollection := client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("REACTIONS_COLLECTION",".."))
	var data = mongoReactionsData{}
	result := reactionCollection.FindOne(mongoCtx, bson.M{"_id": id.Value})
	err := result.Decode(&data)
	if err != nil {
		log.Fatal("error while decoding ", err)
		return data
	}
	return data
}

func update(field string, method string, id *reactions.ReactionUUID) {
	reactionCollection := client.Database(GetEnvVariable("TWIT_DB","..")).
		Collection(GetEnvVariable("REACTIONS_COLLECTION",".."))
	opts := options.Update().SetUpsert(true)
	switch method {
	case decrement:
		_, err := reactionCollection.UpdateOne(mongoCtx, bson.M{"_id": id.Value},
			bson.M{"$inc": bson.M{field: -1}}, opts)
		if err != nil {
			log.Fatal(err)
		}
	case increment:
		_, err := reactionCollection.UpdateOne(mongoCtx, bson.M{"_id": id.Value},
			bson.M{"$inc": bson.M{field: 1}}, opts)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (* Like) GetTwitLikes(c context.Context, id *reactions.ReactionUUID) (*reactions.Like, error) {
	data := get(id)
	response := &reactions.Like{TwitId: id, LikeCounter: data.LikeCounter}
	return response, nil
}

func (* Like) LikeTwit(c context.Context, id *reactions.ReactionUUID) (*reactions.ResponseReaction, error){
	update("likes", increment, id)
	twitStatus := &reactions.ResponseReaction{Value: "Twit was successfully liked"}
	return twitStatus, nil
}

func (* Like) UnlikeTwit (c context.Context, id *reactions.ReactionUUID) (*reactions.ResponseReaction, error) {
	update("likes", decrement, id)
	twitStatus := &reactions.ResponseReaction{Value: "Twit was successfully unliked"}
	return twitStatus, nil
}

func (* Retwit) GetTwitRetwits(c context.Context, id *reactions.ReactionUUID) (*reactions.Retwit, error){
	data := get(id)
	response := &reactions.Retwit{TwitId: id, RetwitCounter: data.RetwitCounter}
	return response, nil
}

func (* Retwit) RetwitTwit(c context.Context, id *reactions.ReactionUUID) (*reactions.ResponseReaction, error){
	update("retwits", increment, id)
	twitStatus := &reactions.ResponseReaction{Value: "Twit was successfully retwited"}
	return twitStatus, nil
}

func (* Retwit) UnretwitTwit(c context.Context, id *reactions.ReactionUUID) (*reactions.ResponseReaction, error){
	update("retwits", decrement, id)
	twitStatus := &reactions.ResponseReaction{Value: "Twit was successfully unretwited"}
	return twitStatus, nil
}