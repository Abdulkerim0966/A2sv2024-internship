package Usecases

import (
	"context"
	"task/Domain"
	Infrastructure "task/Infrastructure"

)

type userUsecase struct {
	userRepository Domain.UserRepository
}
func NewUserUsecase(userRepository Domain.UserRepository) Domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}
func (uc *userUsecase) RegistorUser(ctx context.Context, user *Domain.User) error {
	return uc.userRepository.RegistorUser(ctx, user)
}

func (lu *userUsecase) GetUserByUserName(ctx context.Context, userName string) (Domain.User, error) {
	return lu.userRepository.GetUserByUserName(ctx, userName)
}

func (lu *userUsecase) CreateAccessToken(user *Domain.User ,secret string) (accessToken string,err error) {
   return Infrastructure.CreateAccessToken(user,secret)
}












// var collectionUser *mongo.Collection

// // add user to database
// func AddToUserDataBase(newUser Domain.User) bool {
// 	_, err := collectionUser.InsertOne(context.TODO(), newUser)
// 	return err == nil
// }

// // check the user validation
// func CheckUser(username string) (Domain.User, error) {

// 	filter := bson.M{"username": username}
// 	var result Domain.User

// 	err := collectionUser.FindOne(context.TODO(), filter).Decode(&result)

// 	return result, err

// }
