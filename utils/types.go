package utils

import "google.golang.org/protobuf/types/known/timestamppb"

// MongoReactionsData struct for decoding reactions data from mongodb, has the same fields as in database
type MongoReactionsData struct {
	ID            string `json:"_id" bson:"_id"`
	LikeCounter   int32  `json:"likes" bson:"likes"`
	RetwitCounter int32  `json:"retwits" bson:"retwits"`
}

// MongoDecodedTwitData struct for decoding twit data from mongodb, has the same fields as in database
type MongoDecodedTwitData struct {
	ID       string                 `json:"_id" bson:"_id"`
	Date     *timestamppb.Timestamp `json:"date" bson:"date"`
	Text     string                 `json:"text" bson:"text"`
	Nickname string                 `json:"nickname" bson:"nickname"`
}

type MongoDecodedNames struct {
	Name []string `json:"users" bson:"users"`
}

type MongoDecodedIds struct {
	TwitId string `json:"twit_id" bson:"twit_id"`
}
type MongoDecodedList struct {
	ID        string   `json:"_id" bson:"_id"`
	Name      string   `json:"name" bson:"name"`
	Usernames []string `json:"users" bson:"users"`
}