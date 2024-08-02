package controllers

import (
	"task/data"
	"task/router"
)

func Controller() {
	router.Router.GET("/tasks", data.GetTasks)

	router.Router.GET("/tasks/:id", data.GetTask)

	router.Router.POST("/tasks", data.CreateTask)

	router.Router.DELETE("/tasks/:id", data.RemoveTask)

	router.Router.PUT("/tasks/:id", data.UpdatedTask)

	router.Router.Run()

}
