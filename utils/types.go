package utils

import "google.golang.org/protobuf/types/known/timestamppb"

type MongoReactionsData struct {
	ID            string `json:"_id" bson:"_id"`
	LikeCounter   int32  `json:"likes" bson:"likes"`
	RetwitCounter int32  `json:"retwits" bson:"retwits"`
}

type MongoDecodedTwitData struct {
	ID   string                 `json:"_id" bson:"_id"`
	Date *timestamppb.Timestamp `json:"date" bson:"date"`
	Text string                 `json:"text" bson:"text"`
	User string                 `json:"user_id" bson:"user_id"`
}

type MongoDecodedUsersIds struct {
	Ids    []string `json:"users" bson:"users"`
}

type MongoDecodedListsNames struct {
	Name      string `json:"name" bson:"name"`
	Creator   string `json:"creator_id" bson:"creator_id"`
}

type MongoDecodedUsernames struct {
	Username string `json:"username" bson:"username"`
}

type MongoDecodedTwitIds struct {
	TwitId string `json:"twit_id" bson:"twit_id"`
}
type MongoDecodedList struct {
	ID      string   `json:"_id" bson:"_id"`
	Name    string   `json:"name" bson:"name"`
	UserIds []string `json:"users" bson:"users"`
}

type MongoDecodedUser struct {
	ID       string `json:"_id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}