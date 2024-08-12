package router

import (
	"os"
	"task/Delivery/controllers"
	"task/Domain" // Add this import statement
	"task/Infrastructure"
	repositories "task/Repositories"
	"task/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(db mongo.Database, gin *gin.Engine) {

	var PublicRouter = gin.Group("")

	newUserRouter(db, PublicRouter)

	var ProtectedRouter = gin.Group("")
	ProtectedRouter.Use(Infrastructure.AuthMidddleware(os.Getenv("SECRET")))

	newTaskRouter(db, ProtectedRouter)
	// newPeblicRouter =(client,PublicRouter)

	// Router.Run()

}

func newUserRouter(db mongo.Database, Group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(db, Domain.CollectionUser)
	uc := &controllers.UserController{
		UserUsecase: Usecases.NewUserUsecase(ur),
	}
	Group.POST("/register", uc.RegistorUser)
	Group.POST("/login", uc.Login)
}

func newTaskRouter(db mongo.Database, Group *gin.RouterGroup) {
	tr := repositories.NewTaskRepository(db, Domain.CollectionTask)
	tc := &controllers.TaskController{
		TaskUsecase: Usecases.NewTaskUsecase(tr),
	}
	Group.GET("/tasks", Infrastructure.AuthMidddleware(os.Getenv("SECRET")), tc.GetTasks)

	Group.GET("/tasks/:id", Infrastructure.AuthMidddleware(os.Getenv("SECRET")), tc.GetTaskById)

	Group.POST("/tasks", Infrastructure.AuthMidddleware(os.Getenv("SECRET")), tc.CreateTask)

	Group.DELETE("/tasks/:id", Infrastructure.AuthMidddleware(os.Getenv("SECRET")), tc.RemoveTask)

	Group.PUT("/tasks/:id", Infrastructure.AuthMidddleware(os.Getenv("SECRET")), tc.UpdatedTask)

}
