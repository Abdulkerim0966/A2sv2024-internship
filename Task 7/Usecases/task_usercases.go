package Usecases

import (
	"context"

	"task/Domain"
)


type TaskUsecase struct {
	taskRepository Domain.TaskRepository
}

func NewTaskUsecase(taskRepository Domain.TaskRepository) Domain.TaskUsecase {
	return &TaskUsecase{
		taskRepository: taskRepository,
	}
}
func (tu *TaskUsecase) CreateTask(ctx context.Context, task *Domain.Task) error {
	return tu.taskRepository.CreateTask(ctx, task)
}
func (tu *TaskUsecase) GetAllTasks() ([]Domain.Task ,error) {
	return tu.taskRepository.GetAllTasks()
}
func (tu *TaskUsecase) GetTaskByID(ctx context.Context, taskid string) (Domain.Task, error) {
	return tu.taskRepository.FindTaskbyID(ctx, taskid)
}
func (tu *TaskUsecase) DeleteTask(ctx context.Context, taskid string) error {
	return tu.taskRepository.DeleteTask(ctx, taskid)
}
func (tu *TaskUsecase) UpdateTask(ctx context.Context, taskid string, task Domain.Task) error {
	return tu.taskRepository.UpdateTask(ctx, taskid, task)
}




















// // Adding data to database
// func AddToDataBase(newTask Domain.Task) bool {
// 	_, err := collectionTask.InsertOne(context.TODO(), newTask)

// 	return err == nil
// }

// // get all batabase


// //func DeleteTask(ctx *gin.Context) error {
// 	id := ctx.Param("id")

// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = collectionTask.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}})
// 	return err
// }

// func UpdateDatabase(ctx *gin.Context) error {
// 	id := ctx.Param("id")
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return err
// 	}
// 	var updatedTask Domain.Task
// 	if err := ctx.ShouldBindJSON(&updatedTask); err != nil {
// 		return err
// 	}
// 	update := bson.D{
// 		{Key: "$set", Value: updatedTask},
// 	}
// 	_, err = collectionTask.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}}, update)
// 	return err

// }
