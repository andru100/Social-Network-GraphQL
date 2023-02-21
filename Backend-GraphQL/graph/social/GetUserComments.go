package social

import (
	"context"
	"errors"
	"sort"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/andru100/Graphql-Social-Network/graph/model"
)

func GetUserComments(username string)(*model.MongoFields, error){
	
	collection := Client.Database("datingapp").Collection("userdata") // connect to db and collection.
	currentDoc := model.MongoFields{}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	err := collection.FindOne(ctx, bson.M{"Username": username}).Decode(&currentDoc)
	if err != nil {
		err = errors.New("unable to find users data")
		return nil, err
	}

	sort.Slice(currentDoc.Posts, func(i, j int) bool { // needs to be done on adding post and remove this
		return currentDoc.Posts[i].TimeStamp > currentDoc.Posts[j].TimeStamp
	})

	return &currentDoc, err
}