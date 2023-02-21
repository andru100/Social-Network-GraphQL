package social

import (
	"context"
	"errors"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/andru100/Graphql-Social-Network/graph/model"
)

func SignUp(newUserData *model.NewUserDataInput) (string, error) { // takes id and sets up bucket and mongodb

	var token string

	collection := Client.Database("datingapp").Collection("userdata") // connect to db and collection.

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	// search for duplicate username 
	//TODO change this to a map rather than search all docs
	result := model.MongoFields{}

	err := collection.FindOne(ctx, bson.M{"Username": newUserData.Username}).Decode(&result)

	if err == nil {
		err = errors.New("username in use")
		return token, err
	}

	createuser := model.Usrsignin{Username: newUserData.Username, Email: newUserData.Email, Password: newUserData.Password, Photos: []string{}, LastCommentNum: 0, Posts: []model.PostData{} }

	//username not in use so add new userdata struct
	_, err = collection.InsertOne(context.TODO(), createuser)
	if err != nil {
		err = errors.New("problem creating user")
		return token, err
	}


	Createbucket(newUserData.Username) // create bucket to store users files

	token, err = MakeJwt(&newUserData.Username, true) // make jwt with user id and auth true

	return token, err
}
