package methods

import (
	"context"
	"github.com/magiconair/properties/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"project-twit/proto/reactions"
	"testing"
)

var idWithReactions reactions.ReactionUUID

func getRandomReactionsDocument(collection *mongo.Collection) mongoReactionsData {
	pipeline := []bson.D{{{"$sample", bson.D{{"size", 1}}}}}
	cursor, err := collection.Aggregate(context.Background(), pipeline)
	var data mongoReactionsData
	if err != nil {
		log.Fatal(err)
	}
	for cursor.Next(context.Background()) {
		err := cursor.Decode(&data)
		if err != nil {
			log.Fatal(err)
		}
	}
	return data
}

func TestGetTwitLikes(t *testing.T) {
	_, ctx, closeCon := DatabaseInit()
	defer closeCon()
	collection := client.Database(GetEnvVariable("TWIT_DB",".")).
		Collection(GetEnvVariable("TWIT_COLLECTION","."))
	data := getRandomReactionsDocument(collection)
	reactStruct := Like{}
	id := reactions.ReactionUUID{Value: data.ID}
	response, err := reactStruct.GetTwitLikes(ctx, &id)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, data.LikeCounter, response.LikeCounter,
		"response likecounter should be the same")
}

func TestGetTwitRetwits(t *testing.T) {
	_, ctx, closeCon := DatabaseInit()
	defer closeCon()
	collection := client.Database(GetEnvVariable("TWIT_DB",".")).
		Collection(GetEnvVariable("TWIT_COLLECTION","."))
	data := getRandomReactionsDocument(collection)
	reactStruct := Retwit{}
	id := reactions.ReactionUUID{Value: data.ID}
	response, err := reactStruct.GetTwitRetwits(ctx, &id)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, data.RetwitCounter, response.RetwitCounter,
		"response retwitcounter should be the same")
}

func TestRetwitTwit(t *testing.T) {
	_, ctx, closeCon := DatabaseInit()
	defer closeCon()
	reactStruct := Retwit{}
	response, err := reactStruct.RetwitTwit(ctx, &idWithReactions)
	if err != nil {
		log.Fatal(err)
	}
	twitStatus := &reactions.ResponseReaction{Value: "Twit was successfully retwited"}
	assert.Equal(t, twitStatus.Value, response.Value,
		"response retwitcounter should be the same")
}

func TestLikeTwit(t *testing.T) {
	_, ctx, closeCon := DatabaseInit()
	defer closeCon()
	reactStruct := Like{}
	response, err := reactStruct.LikeTwit(ctx, &idWithReactions)
	if err != nil {
		log.Fatal(err)
	}
	twitStatus := &reactions.ResponseReaction{Value: "Twit was successfully liked"}
	assert.Equal(t, twitStatus.Value, response.Value,
		"response likecounter should be the same")
}

func TestUnretwitTwit(t *testing.T) {
	_, ctx, closeCon := DatabaseInit()
	defer closeCon()
	reactStruct := Retwit{}
	response, err := reactStruct.UnretwitTwit(ctx, &idWithReactions)
	if err != nil {
		log.Fatal(err)
	}
	twitStatus := &reactions.ResponseReaction{Value: "Twit was successfully unretwited"}
	assert.Equal(t, twitStatus.Value, response.Value,
		"response retwitcounter should be the same")
}

func TestUnlikeTwit(t *testing.T) {
	_, ctx, closeCon := DatabaseInit()
	defer closeCon()
	reactStruct := Like{}
	response, err := reactStruct.UnlikeTwit(ctx, &idWithReactions)
	if err != nil {
		log.Fatal(err)
	}
	twitStatus := &reactions.ResponseReaction{Value: "Twit was successfully unliked"}
	assert.Equal(t, twitStatus.Value, response.Value,
		"response likecounter should be the same")
}