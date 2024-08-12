package controllers

import (
	
	"net/http"
	"os"

	"task/Domain"
	"task/Infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskController struct {
	TaskUsecase Domain.TaskUsecase
}

         // creat task
func (tc *TaskController) CreateTask(ctx *gin.Context) {

	var newTask Domain.Task

	if err := ctx.ShouldBindJSON(&newTask); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = primitive.NewObjectID()

	

	//Authenthication

	claim := ctx.MustGet("claims").(Domain.JwrCustonClaims)

	if !claim.Role && claim.UserName != newTask.Owner {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You can't assign task for other person"})
		return

	}
	
	err := tc.TaskUsecase.CreateTask(ctx, &newTask)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add document to database"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
	// }
}

func (tc *TaskController) GetTasks(ctx *gin.Context) {

		//Authenthication
		
		claim := ctx.MustGet("claims").(Domain.JwrCustonClaims)
		
		if !claim.Role{
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You don't have  authorization to see all tasks"})
			return
		}
		tasks ,err :=tc.TaskUsecase.GetAllTasks()
		if err!=nil{
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve document from database"})
	         return
		}
		ctx.JSON(http.StatusOK, gin.H{"tasks": tasks})

	}

func (tc *TaskController) GetTaskById(ctx *gin.Context) {
	id := ctx.Param("id")
	
	task, err := tc.TaskUsecase.GetTaskByID(ctx, id)
	if err != nil {
	
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	//Authenthication

    claim := ctx.MustGet("claims").(Domain.JwrCustonClaims)

	if !claim.Role && claim.UserName != task.Owner {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You can't see other's task"})
		return

	}
	ctx.JSON(http.StatusOK, gin.H{"tasks": task})

}
	

// // delete

func (tc *TaskController) RemoveTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := tc.TaskUsecase.GetTaskByID(ctx,id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	//Authenthication
	claim := ctx.MustGet("claims").(Domain.JwrCustonClaims)

	if !claim.Role && claim.UserName != task.Owner {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You can't remove other's task"})
		return
	}

	err = tc.TaskUsecase.DeleteTask(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete the document from database"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task removed"})

}

// // update
func (tc *TaskController) UpdatedTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := tc.TaskUsecase.GetTaskByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
		return
	}
	//Authenthication
	claim := ctx.MustGet("claims").(Domain.JwrCustonClaims)
	if !claim.Role && claim.UserName != task.Owner {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "You can't update other's task"})
		return
	}
	err = tc.TaskUsecase.UpdateTask(ctx, id, task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update the task"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Task succesfully updated"})
}






type UserController struct {
	UserUsecase Domain.UserUsecase
}

func (uc *UserController) RegistorUser(ctx *gin.Context) {
	var newPersonal Domain.User

	if err := ctx.ShouldBindJSON(&newPersonal); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := Infrastructure.HashPassword(newPersonal.Password)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	newPersonal.Password = string(hashedPassword)

	err = uc.UserUsecase.RegistorUser(ctx, &newPersonal)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add  to database"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "succesfully registered"})

}

func (uc *UserController) Login(ctx *gin.Context) {
	var request Domain.User

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	user, err := uc.UserUsecase.GetUserByUserName(ctx, request.UserName)

	
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid  username"})
		return
	}
	if Infrastructure.ComparePassword(user.Password, request.Password) != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}


	accessToken, err := Infrastructure.CreateAccessToken(&request, os.Getenv("SECRET"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": accessToken})

}



