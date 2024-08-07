package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"task/data"
	"task/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func RegistorUser(ctx *gin.Context){
	var newPersonal models.User

	if err := ctx.ShouldBindJSON(&newPersonal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPersonal.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	newPersonal.Password = string(hashedPassword)

	success :=data.AddToUserDataBase(newPersonal)

	
	if success {

		ctx.JSON(http.StatusCreated, gin.H{"message": "succesfully registered"})
	}else{
	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add  to database"})
	}
}

var JwtSecret = []byte(os.Getenv("SECRET"))

//login

func Login(ctx *gin.Context){
	var personel models.User

	if err := ctx.ShouldBindJSON(&personel); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	

// User login logic
result ,success := data.CheckUser(personel.UserName)

fmt.Println(result,success,"personal:::",personel) 

if success !=nil || bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(personel.Password)) != nil {
	ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
	return
}

// Generate JWT
token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
  "userName": result.UserName,
  "role":   result.Role,
  "exp" : time.Now().Add(time.Hour*24*30).Unix(),
})

jwtToken, err := token.SignedString(JwtSecret)
if err != nil {
  ctx.JSON(500, gin.H{"error": "Internal server error"})
  return
}

 ctx.JSON(200, gin.H{"message": "User logged in successfully", "token": jwtToken})
}


func CreateTask(ctx *gin.Context) {
	
	var newTask models.Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
    newTask.ID = primitive.NewObjectID()

	//Authenthication
	user, _:= ctx.Get("user")
	if !user.(models.User).Role && user.(models.User).UserName != newTask.Owner{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You can't assign task for other person"})
		return

	}

	success :=data.AddToDataBase(newTask)

	
	if success {

		ctx.JSON(http.StatusCreated, gin.H{"message": "Task created"})
		return
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
	//Authenthication
	user, _:= ctx.Get("user")
	if !user.(models.User).Role && user.(models.User).UserName != tasks.Owner{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You can't see other's task"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})
}




func GetTasks(ctx *gin.Context){

	//Authenthication
	user, _:= ctx.Get("user")
	if !user.(models.User).Role{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You can't see other's task"})
		return
	}
	tasks ,err :=data.GetAlldatabase()
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve document from database"})
         return
	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})

}
   
func RemoveTask(ctx *gin.Context){
	task,err :=data.FindTaskbyID(ctx)
	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return 
	}
	//Authenthication
	user, _:= ctx.Get("user")
	if !user.(models.User).Role && user.(models.User).UserName != task.Owner{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You can't remove other's task"})
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
	task,err :=data.FindTaskbyID(ctx)
	if err != nil{
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	//Authenthication
	user, _:= ctx.Get("user")
	if !user.(models.User).Role && user.(models.User).UserName != task.Owner{
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You can't update other's task"})
		return
	}
	err = data.UpdateDatabase(ctx)
	if err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update the task"})
	    return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "Task succesfully updated"})

}


