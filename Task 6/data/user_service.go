package data

import (
	"context"
	"task/models"

	"go.mongodb.org/mongo-driver/bson"
)

// add user to database
func AddToUserDataBase(newUser models.User) bool{
	_,err := collectionUser.InsertOne(context.TODO(),newUser)
	return err == nil 
}
// check the user validation
func CheckUser(username string) (models.User ,error){
	

	filter := bson.M{"username": username }
	var result models.User

	err := collectionUser.FindOne(context.TODO(), filter).Decode(&result)
	
	return result ,err

}

