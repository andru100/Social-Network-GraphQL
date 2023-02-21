package social

import (
	"context"
	"time"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/andru100/Graphql-Social-Network/graph/model"
)


func LikeCmt  (input model.SendLikeInput) (*model.MongoFields, error) {

	collection := Client.Database("datingapp").Collection("userdata")

	currentDoc := model.MongoFields{}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	err := collection.FindOne(ctx, bson.M{"Username": input.Username}).Decode(&currentDoc)
	
	//Find the comment being liked
	//TODO make each sub field a mongo doc so that comments can be searched by ID of doc and save on looping through data
	for i := 0; i < len(currentDoc.Posts); i++ {
		if currentDoc.Posts[i].PostNum == input.PostIndx {
			likesent := model.Likes{
				Username: input.LikedBy ,
				Profpic:  input.LikeByPic,
			}
			currentDoc.Posts[i].Likes = append(currentDoc.Posts[i].Likes, likesent) // add like to post
			filter := bson.M{"Username": currentDoc.Posts[i].Username}   
			Updatetype := "$set"
			Key2updt := "Posts"
			update := bson.D{
				{Updatetype, bson.D{
					{Key2updt, currentDoc.Posts},
				}},
			}
		
			//put to db
			_, err = collection.UpdateOne(context.TODO(), filter, update)
			if err != nil {
				err = errors.New("error when adding Like to DB")
				return nil, err
			}
		}
	}

	
	// check originating page request came from and return updated comments to save an extra api call on react refresh component
	if input.ReturnPage == "All" {
		result, err1 := GetAllComments(input.Username)
		return result, err1
    } else {
	   result, err1:= GetUserComments (input.Username)
	   return result, err1
    }
}