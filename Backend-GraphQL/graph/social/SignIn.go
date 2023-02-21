package social

import (
	"context"
	"time"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/andru100/Graphql-Social-Network/graph/model"
)

func Signin(userid string, password string) (*string, error) { // takes id and sets up bucket and mongodb

	collection := Client.Database("datingapp").Collection("userdata") // connect to db and collection

	result := model.MongoFields{}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	err := collection.FindOne(ctx, bson.M{"Username": userid}).Decode(&result)
	
	if err != nil {
		return nil, errors.New("username not found")
	}

	if result.Password == password {
		token, err := MakeJwt(&userid, true)
		return &token, err
	} else {
		return nil, errors.New("password does not match")
	}

}
