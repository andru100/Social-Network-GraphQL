package social

import (
	"context"
	"log"
	"sort"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/andru100/Graphql-Social-Network/graph/model"
)

func GetAllComments(username string) (*model.MongoFields, error) { // gets comments for all friends/users, for the home page feed

	collection := Client.Database("datingapp").Collection("userdata") // connect to db and collection.

	currentDoc := model.MongoFields{}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	var allPosts []model.PostData

	findOptions := options.Find()
	findOptions.SetLimit(2)

	var results []*model.MongoFields

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterating through cursor decode documents one at a time
	for cur.Next(context.TODO()) {

		var elem model.MongoFields
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	for _, record := range results {
		for _, posts := range record.Posts {
			allPosts = append(allPosts, posts)
		}
	}

	//Sort posts or comments by time descending
	//TODO sort posts as they are added to record to save on constantly resorting
	sort.Slice(allPosts, func(i, j int) bool { 
		return allPosts[i].TimeStamp > allPosts[j].TimeStamp
	})

	var json2send model.MongoFields
	json2send.Posts = allPosts
	err = collection.FindOne(ctx, bson.M{"Username": username}).Decode(&currentDoc)
	json2send.Profpic = currentDoc.Profpic
	json2send.Bio = currentDoc.Bio
	json2send.Photos = currentDoc.Photos

	return &json2send, err
	

}
