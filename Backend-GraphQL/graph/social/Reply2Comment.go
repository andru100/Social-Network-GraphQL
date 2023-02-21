package social

import (
	"context"
	"errors"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/andru100/Graphql-Social-Network/graph/model"
)


func Reply2Comment(input model.ReplyCommentInput) (*model.MongoFields, error) { // adds replies to users comments mongo doc

	collection := Client.Database("datingapp").Collection("userdata")

	currentDoc := model.MongoFields{}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	// Find the document that mathes the id 
	err := collection.FindOne(ctx, bson.M{"Username": input.AuthorUsername}).Decode(&currentDoc)

	//Find the comment being replied do by index and add it
	//TODO make each sub filed a mongo doc so we can search by ID and save looping through all comments
	for i := 0; i < len(currentDoc.Posts); i++ {
		if currentDoc.Posts[i].PostNum == input.PostIndx {
			reply := model.MsgCmts{
				Username: input.ReplyUsername ,
				Comment:  input.ReplyComment , 
				Profpic:  input.ReplyProfpic ,
			}
			currentDoc.Posts[i].Comments = append(currentDoc.Posts[i].Comments, reply) // add reply to post
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
				err = errors.New("error when adding comment to DB")
				return nil, err
			}                             
		}
	}

	
	if input.ReturnPage == "All" {
		result, err1 := GetAllComments(input.ReplyUsername)
		return result, err1
    } else {
	   result, err1:= GetUserComments (input.AuthorUsername)
	   return result, err1
    }
}
