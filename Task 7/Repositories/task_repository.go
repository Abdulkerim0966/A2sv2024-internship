package repositories

import (
	"context"

	"task/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type taskRepository struct {
	database   mongo.Database
	collection string
}

func NewTaskRepository(db mongo.Database, collection string) Domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

func (tr *taskRepository) CreateTask(ctx context.Context, task *Domain.Task) error {
	collection := tr.database.Collection(tr.collection)
	_, err := collection.InsertOne(ctx, task)

	return err
}
func (tr *taskRepository) GetAllTasks() ([]Domain.Task, error) {
	collection := tr.database.Collection(tr.collection)
	tasks := []Domain.Task{}

	// Query the database for all tasks.
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	// Iterate over the cursor and decode each task into a Task struct.
	err = cursor.All(context.Background(), &tasks)
	return tasks, err
}

func (tr *taskRepository) FindTaskbyID(ctx context.Context, taskid string) (Domain.Task, error) {
	collection := tr.database.Collection(tr.collection)
	objectID, err := primitive.ObjectIDFromHex(taskid)

	if err != nil {

		return Domain.Task{}, err
	}
	filter := bson.D{{Key: "_id", Value: objectID}}

	var result Domain.Task

	err = collection.FindOne(context.TODO(), filter).Decode(&result)

	return result, err

}
func (tr *taskRepository) DeleteTask(ctx context.Context, taskid string) error {
	collection := tr.database.Collection(tr.collection)
	objectID, err := primitive.ObjectIDFromHex(taskid)

	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}})
	return err
}
func (tr *taskRepository) UpdateTask(ctx context.Context, taskid string, task Domain.Task) error {
	collection := tr.database.Collection(tr.collection)
	objectID, err := primitive.ObjectIDFromHex(taskid)

	if err != nil {
		return err
	}
	_, err = collection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: objectID}}, bson.D{{Key: "$set", Value: task}})
	return err
}
