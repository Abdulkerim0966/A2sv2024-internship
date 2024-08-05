package data

import (
	"context"
	"fmt"

	"log"

	"task/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var ClientOptions *options.ClientOptions
var Client *mongo.Client

var collection *mongo.Collection


func ConnectTodatabase(){
	
	var ClientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	Client,err := mongo.Connect(context.TODO(), ClientOptions)

	
	if err != nil {
		log.Fatal(err)
	}
	
	// Check the connection
	err = Client.Ping(context.TODO(), nil)
	
	if err != nil {
		log.Fatal(err)
	}
	collection =Client.Database("Task_manager").Collection("Task")

}

// disconnect it from database
func DisconnectToDatabase(){
	err := Client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	}


}


func FindTaskbyID(ctx *gin.Context) (models.Task ,error){
	id := ctx.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
    
	if err != nil{
		return models.Task{},err
	}


	filter := bson.D{{Key: "_id", Value: objectID}}
 
	var result models.Task

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	
	return result ,err

}



// Adding data to database
func AddToDataBase(newTask models.Task) bool{
	_,err := collection.InsertOne(context.TODO(),newTask)

	return err == nil 
}

// get all batabase
func GetAlldatabase()( []*models.Task , error){
	findOptions := options.Find()
    var results []*models.Task

	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}


for cur.Next(context.TODO()) {
    
    
    var elem models.Task
    err := cur.Decode(&elem)
    if err != nil {
		fmt.Println("here")
        log.Fatal(err)
    }

    results = append(results, &elem)
}

if err := cur.Err(); err != nil {
    log.Fatal(err)
}

cur.Close(context.TODO())
return results,err
}


func DeleteTask(ctx *gin.Context) error{
	id := ctx.Param("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objectID,}})
return err
}


func UpdateDatabase(ctx *gin.Context) error{
	id := ctx.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	var updatedTask models.Task
	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
		return err
	}
	update := bson.D{
		{Key: "$set", Value: updatedTask},
	}
	_,err = collection.UpdateOne(context.TODO(),bson.D{{Key: "_id", Value: objectID}},update)
	return err

}



