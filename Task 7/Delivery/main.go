package main

import (
	"context"
	"fmt"
	"log"

	"task/Delivery/router"
	// "task/Usecases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	fmt.Println("=======================================")
	fmt.Println("|                                     |")
	fmt.Println("|     WELL COME TO TASK MANAGMENT     |")
	fmt.Println(`|                                     |`)
	fmt.Println("=======================================")

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}

	// Connect to MongoDB

	ClientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	db, err := mongo.Connect(context.TODO(), ClientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = db.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	// db :=Client.Database(os.Getenv("DatabaseName"))

	gin := gin.Default()

	router.Setup(*db.Database("Task_manager"), gin)
	gin.Run()

	// var ClientOptions *options.ClientOptions
	// var Client *mongo.Client

	// var collectionTask *mongo.Collection

	// func ConnectTodatabase() {

	// 	var ClientOptions = options.Client().ApplyURI("mongodb://localhost:27017")
	// 	Client, err := mongo.Connect(context.TODO(), ClientOptions)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	// Check the connection
	// 	err = Client.Ping(context.TODO(), nil)

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	collectionTask = Client.Database("Task_manager").Collection("Task")
	// 	collectionUser = Client.Database("Task_manager").Collection("User")

	// }

	// // disconnect it from database
	// func DisconnectToDatabase() {
	// 	err := Client.Disconnect(context.TODO())

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// }

	// Usecases.ConnectTodatabase()

	// router.Router()
	// Usecases.DisconnectToDatabase()

}
