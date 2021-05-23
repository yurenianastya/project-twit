package grpc

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"project-twit/proto/reactions"
)

var reactioncol *mongo.Collection

type mongoReactionsData struct {
	ID            string    `json:"_id" bson:"_id"`
	LikeCounter   int32       `json:"likes" bson:"likes"`
	RetwitCounter int32      `json:"retwits" bson:"retwits"`
}

type Like struct {
	reactions.UnimplementedLikeServiceServer
}

type Retwit struct {
	reactions.UnimplementedRetwitServiceServer
}

func (* Like) GetTwitLikes(c context.Context, id *reactions.ReactionUUID) (*reactions.Like, error) {
	reactioncol = client.Database("twit").Collection("reactions")
	var data = mongoReactionsData{}
	result := reactioncol.FindOne(mongoCtx, bson.M{"_id": id.Value})
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not find twitreactions with Object ID %s: %v", id.Value, err))
	}
	response := &reactions.Like{TwitId: id, LikeCounter: data.LikeCounter}
	return response, nil
}

func (* Like) LikeTwit(c context.Context, id *reactions.ReactionUUID) (*reactions.Like, error){
	reactioncol = client.Database("twit").Collection("reactions")
	var data = mongoReactionsData{}
	result := reactioncol.FindOne(mongoCtx, bson.M{"_id": id.Value})
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not find twitreactions with Object ID %s: %v", id.Value, err))
	}
	data.LikeCounter++
	opts := options.Update().SetUpsert(true)
	update := bson.M{"$set": bson.M{"likes": data.LikeCounter}}
	_, err := reactioncol.UpdateOne(mongoCtx, bson.M{"_id": id.Value}, update, opts)
	if err != nil {
		return nil, err
	}
	return &reactions.Like{TwitId: id, LikeCounter: data.LikeCounter}, nil
}

func (* Like) UnlikeTwit (c context.Context, id *reactions.ReactionUUID) (*reactions.Like, error) {
	reactioncol = client.Database("twit").Collection("reactions")
	var data = mongoReactionsData{}
	result := reactioncol.FindOne(mongoCtx, bson.M{"_id": id.Value})
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not find twitreactions with Object ID %s: %v", id.Value, err))
	}
	data.LikeCounter--
	opts := options.Update().SetUpsert(true)
	update := bson.M{"$set": bson.M{"likes": data.LikeCounter}}
	_, err := reactioncol.UpdateOne(mongoCtx, bson.M{"_id": id.Value}, update, opts)
	if err != nil {
		return nil, err
	}
	return &reactions.Like{TwitId: id, LikeCounter: data.LikeCounter}, nil
}

func (* Retwit) GetTwitRetwits(c context.Context, id *reactions.ReactionUUID) (*reactions.Retwit, error){
	reactioncol = client.Database("twit").Collection("reactions")
	var data = mongoReactionsData{}
	result := reactioncol.FindOne(mongoCtx, bson.M{"_id": id.Value})
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not find twitreactions with Object ID %s: %v", id.Value, err))
	}
	response := &reactions.Retwit{TwitId: id, RetwitCounter: data.RetwitCounter}
	return response, nil
}

func (* Retwit) RetwitTwit(c context.Context, id *reactions.ReactionUUID) (*reactions.Retwit, error){
	reactioncol = client.Database("twit").Collection("reactions")
	var data = mongoReactionsData{}
	result := reactioncol.FindOne(mongoCtx, bson.M{"_id": id.Value})
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not find twitreactions with Object ID %s: %v", id.Value, err))
	}
	data.RetwitCounter++
	opts := options.Update().SetUpsert(true)
	update := bson.M{"$set": bson.M{"retwits": data.RetwitCounter}}
	_, err := reactioncol.UpdateOne(mongoCtx, bson.M{"_id": id.Value}, update, opts)
	if err != nil {
		return nil, err
	}
	return &reactions.Retwit{TwitId: id, RetwitCounter: data.RetwitCounter}, nil
}

func (* Retwit) UnretwitTwit(c context.Context, id *reactions.ReactionUUID) (*reactions.Retwit, error){
	reactioncol = client.Database("twit").Collection("reactions")
	var data = mongoReactionsData{}
	result := reactioncol.FindOne(mongoCtx, bson.M{"_id": id.Value})
	if err := result.Decode(&data); err != nil {
		return nil, status.Errorf(codes.NotFound,
			fmt.Sprintf("Could not find twitreactions with Object ID %s: %v", id.Value, err))
	}
	data.RetwitCounter--
	opts := options.Update().SetUpsert(true)
	update := bson.M{"$set": bson.M{"retwits": data.RetwitCounter}}
	_, err := reactioncol.UpdateOne(mongoCtx, bson.M{"_id": id.Value}, update, opts)
	if err != nil {
		return nil, err
	}
	return &reactions.Retwit{TwitId: id, RetwitCounter: data.RetwitCounter}, nil
}