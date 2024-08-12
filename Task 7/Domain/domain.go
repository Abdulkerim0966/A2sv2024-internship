package Domain

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionTask = "task"
	CollectionUser = "user"
)

type Task struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Owner       string             `json:"owner"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	DueDate     time.Time          `json:"due_date"`
	Status      string             `json:"status"`
}
type User struct {
	UserName string `json:"username"`
	Role     bool   `json:"role"`
	Password string `json:"password"`
}
type TaskRepository interface {
	CreateTask(ctx context.Context, task *Task) error
	GetAllTasks() ([]Task, error)
	FindTaskbyID(ctx context.Context, taskid string) (Task, error)
	DeleteTask(ctx context.Context, taskid string) error
	UpdateTask(ctx context.Context, taskid string, task Task) error
	
	 
}
type TaskUsecase interface {
	CreateTask(ctx context.Context, task *Task) error
	GetAllTasks() ([]Task, error)
	GetTaskByID(ctx context.Context, taskid string) (Task, error)
	DeleteTask(ctx context.Context, taskid string) error
	UpdateTask(ctx context.Context, taskid string, task Task) error
}
type UserRepository interface {
	RegistorUser(ctx context.Context, user *User) error
	GetUserByUserName(ctx context.Context, userName string) (User, error)
	
}
type UserUsecase interface {
	RegistorUser(ctx context.Context, user *User) error
	GetUserByUserName(ctx context.Context, userName string) (User, error)
	CreateAccessToken(user *User, secret string) (accessToken string, err error)
}
type JwrCustonClaims struct {
	UserName string `json:"username"`
	Role     bool   `json:"role"`
	jwt.StandardClaims
}
