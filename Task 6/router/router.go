package router

import (
	"task/controllers"
	"task/middleware"

	"github.com/gin-gonic/gin"
)


func Router(){


var Router = gin.Default()
Router.POST("/register",controllers.RegistorUser)
Router.POST("/login",controllers.Login)


Router.GET("/tasks", middleware.AuthMidddleware(),controllers.GetTasks)

Router.GET("/tasks/:id",middleware.AuthMidddleware(),controllers.GetTaskById)

Router.POST("/tasks",middleware.AuthMidddleware(), controllers.CreateTask)

Router.DELETE("/tasks/:id",middleware.AuthMidddleware(),controllers.RemoveTask)

Router.PUT("/tasks/:id",middleware.AuthMidddleware(), controllers.UpdatedTask)

Router.Run()

}


