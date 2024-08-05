package main

import (
	"fmt"
	"task/data"
	"task/router"
)


func main() {

	data.ConnectTodatabase()
	
	fmt.Println("=======================================")
	fmt.Println("|                                     |")
	fmt.Println("|     WELL COME TO TASK MANAGMENT     |")
	fmt.Println(`|                                     |`)
	fmt.Println("=======================================")

	

	router.Router()
	data.DisconnectToDatabase()
	
	
}
