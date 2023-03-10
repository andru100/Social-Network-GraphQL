// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	//"github.com/99designs/gqlgen/graphql"
	jwt "github.com/dgrijalva/jwt-go"
)

type Authd struct {
	AuthdUser string `json:"AuthdUser"`
}

type Claims struct {
	Username string `json:"Username"`
	jwt.StandardClaims
}

type GetUserCmts struct {
	Page     string `json:"Page"`
	UserName string `json:"UserName"`
}

type Jwtdata struct {
	Token string `json:"Token"`
}

type JwtdataInput struct {
	Token string `json:"Token"`
}

type Likes struct {
	Username string `bson:"Username" json:"Username"`
	Profpic  string `bson:"Profpic" json:"Profpic"`
}

type MongoFields struct {
	Key            string     `json:"Key"`
	ID 			   string 	  `bson:"_id" json:"ID"`
	Username       string     `json:"Username"`
	Password       string     `json:"Password"`
	Email          string     `json:"Email"`
	Bio            string     `json:"Bio"`
	Profpic        string     `json:"Profpic"`
	Photos         []string   `json:"Photos"`
	LastCommentNum int        `json:"LastCommentNum"`
	Posts          []PostData `json:"Posts"`
}

type MsgCmts struct {
	Username string `bson:"Username" json:"Username"`
	Comment  string `bson:"Comment" json:"Comment"`
	Profpic  string `bson:"Profpic json:"Profpic"`
}

type MutationInput struct {
	Name string `json:"name"`
	ID   string `json:"ID"`
}

type NewUserDataInput struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
	Email    string `json:"Email"`
}

type PostData struct {
	Username    string    `bson:"Username" json:"Username"`
	SessionUser string    `bson:"SessionUser" json:"SessionUser"`
	MainCmt     string    `bson:"MainCmt" json:"MainCmt"`
	PostNum     int       `bson:"PostNum" json:"PostNum"`
	Time        string    `bson:"Time" json:"Time"`
	TimeStamp   int64     `bson:"TimeStamp" json:"TimeStamp"`
	Date        string    `bson:"Date" json:"Date"`
	Comments    []MsgCmts `bson:"Comments" json:"Comments"`
	Likes       []Likes   `bson:"Likes" json:"Likes"`
}

type ReplyCommentInput struct {
	AuthorUsername string `json:"AuthorUsername"`
	ReplyUsername  string `json:"ReplyUsername"`
	ReplyComment   string `json:"ReplyComment"`
	ReplyProfpic   string `json:"ReplyProfpic"`
	PostIndx       int    `json:"PostIndx"`
	ReturnPage     string `json:"ReturnPage"`
}

type SendCmtInput struct {
	Username    string `json:"Username"`
	SessionUser string `json:"SessionUser"`
	MainCmt     string `json:"MainCmt"`
	Time        string `json:"Time"`
	TimeStamp   int64    `json:"TimeStamp"`
	Date        string `json:"Date"`
	ReturnPage  string `json:"ReturnPage"`
}

type SendLikeInput struct {
	Username   string `json:"Username"`
	LikedBy    string `json:"LikedBy"`
	LikeByPic  string `json:"LikeByPic"`
	PostIndx   int    `json:"PostIndx"`
	ReturnPage string `json:"ReturnPage"`
}

type UpdateBioInput struct {
	Username string `json:"Username"`
	Bio      string `json:"Bio"`
}

// type UploadInput struct {
// 	//Username string         `json:"Username"`
// 	File     graphql.Upload `json:"File"`
// 	//Type     string         `json:"Type"`
// }

type Usrsignin struct {
	Username       string     `bson:"Username" json:"Username"`
	Password       string     `bson:"Password" json:"Password"`
	Email          string     `bson:"Email" json:"Email"`
	Bio            string     `bson:"Bio" json:"Bio"`
	Photos         []string   `bson:"Photos" json:"Photos"`
	LastCommentNum int        `bson:"LastCommentNum" json:"LastCommentNum"`
	Posts          []PostData `bson:"Posts" json:"Posts"`
}

type UsrsigninInput struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}
