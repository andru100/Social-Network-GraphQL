package social

import (
	"context"
	"errors"
	"time"
	"github.com/andru100/Graphql-Social-Network/graph/model"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateBio(request model.UpdateBioInput) (*model.MongoFields, error) { // updates user bio section
	
	collection := Client.Database("datingapp").Collection("userdata")

	filter := bson.M{"Username": request.Username}

	Updatetype := "$set"
	Key2updt := "Bio"

	update := bson.D{
		{Updatetype, bson.D{
			{Key2updt, request.Bio},
		}},
	}

	//put to db
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		err = errors.New("error when updating to DB")
		return nil, err
	}

	currentDoc := model.MongoFields{}

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	err = collection.FindOne(ctx, bson.M{"Username": request.Username}).Decode(&currentDoc)

	return &currentDoc, err
}
