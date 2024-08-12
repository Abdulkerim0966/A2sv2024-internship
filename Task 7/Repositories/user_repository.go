package repositories

import (
	"context"
	"task/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	database   mongo.Database
	collection string
}


func NewUserRepository(db mongo.Database, user string) Domain.UserRepository {
	
	return &UserRepository{
		database:   db,
		collection: user,
	}
}




func (ur *UserRepository) RegistorUser(ctx context.Context, user *Domain.User) error {
	collection := ur.database.Collection(ur.collection)
	_, err := collection.InsertOne(ctx, user)
	
	return err
}

// Login implements Domain.UserRepository.
func (ur *UserRepository) GetUserByUserName(ctx context.Context, userName string) (Domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user Domain.User
	err := collection.FindOne(ctx, bson.M{"username" :userName}).Decode(&user)
	return user, err
}

