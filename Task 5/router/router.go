package router

import (
	"task/controllers"

	"github.com/gin-gonic/gin"

)


func Router(){



var Router = gin.Default()

Router.GET("/tasks", controllers.GetTasks)

Router.GET("/tasks/:id",controllers.GetTaskById)

Router.POST("/tasks", controllers.CreateTask)

Router.DELETE("/tasks/:id",controllers.RemoveTask)

Router.PUT("/tasks/:id", controllers.UpdatedTask)

Router.Run()

}


