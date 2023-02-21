package social

import (
	"context"
	"time"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/andru100/Graphql-Social-Network/graph/model"
)

func NewComment (input model.SendCmtInput) (*model.MongoFields, error) {
	
	collection := Client.Database("datingapp").Collection("userdata")

	currentDoc := model.MongoFields{}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	err := collection.FindOne(ctx, bson.M{"Username": input.Username}).Decode(&currentDoc)

	currentDoc.LastCommentNum += 1
	
	//initialise empty slice to hold future likes and reply comments
	cmtHolder := []model.MsgCmts{}
	likeHolder := []model.Likes{}

	//make new comment struct: 
	newPost := model.PostData{
		Username:    input.Username,    
		SessionUser: input.SessionUser,
		MainCmt:     input.MainCmt,  
		PostNum:     currentDoc.LastCommentNum,    
		Time:        input.Time,   
		TimeStamp:   input.TimeStamp,    
		Date:        input.Date,    
		Comments:    cmtHolder ,
		Likes:      likeHolder, 
	}

	currentDoc.Posts = append(currentDoc.Posts, newPost)
	filter := bson.M{"Username": input.Username} 

	//add new comment to DB 
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

	//update post index count
	update = bson.D{
		{Updatetype, bson.D{
			{"LastCommentNum", currentDoc.LastCommentNum},
		}},
	}

	//put to db
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		err = errors.New("error when updating post index")
		return nil, err
	}

	// check originating page the request came from and return updated comments
	if input.ReturnPage == "All" {
		result, err1 := GetAllComments(input.Username)
		return result, err1
    } else {
	   result, err1:= GetUserComments (input.Username)
	   return result, err1
    }
}