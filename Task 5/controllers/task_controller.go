package controllers

import (
	"net/http"

	"task/data"
	"task/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


func CreateTask(ctx *gin.Context) {
	
	var newTask models.Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
    newTask.ID = primitive.NewObjectID()
	success :=data.AddToDataBase(newTask)

	
	if success {

		ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
	}else{
	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add document to database"})
	}
	}
	
func GetTaskById(ctx *gin.Context) {
	tasks,err := data.FindTaskbyID(ctx)
	
	if err !=nil{
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}




func GetTasks(ctx *gin.Context){
	tasks ,err :=data.GetAlldatabase()
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve document from database"})
         return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})

}
   
func RemoveTask(ctx *gin.Context){
	_,err :=data.FindTaskbyID(ctx)
	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return 
	}

	err =data.DeleteTask(ctx)
	if err!= nil {
	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete the document from database"})
    return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})

}


// update
func UpdatedTask(ctx *gin.Context){
	_,err :=data.FindTaskbyID(ctx)
	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	err = data.UpdateDatabase(ctx)
	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update the task"})
	    return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Task succesfully updated"})

}


