package protocol

import (
	"context"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserServices interface {
	CreateUser(ctx context.Context, user User) (User, error)
	GetUser(ctx context.Context, userID uint) (User, error)
	UpdateUser(ctx context.Context, user User) (User, error)
	FindUsers(ctx context.Context) ([]User, error)
	DeleteUser(ctx context.Context, userID uint) error
	DoesUserExist(ctx context.Context, userID uint) (bool, error)
}

// type UserInteractor interface {
// 	CreateUser(context.Context, dto.CreateUserRequest) (dto.CreateUserResponse, error)
// 	UpdateUser(context.Context, dto.UpdateUserRequest) (dto.UpdateUserResponse, error)
// 	FindUsers(context.Context, dto.FindUsersRequest) (dto.FindUsersResponse, error)
// 	FindUser(context.Context, dto.FindUserRequest) (dto.FindUserResponse, error)
// 	DeleteUser(context.Context, dto.DeleteUserRequest) (dto.DeleteUserResponse, error)
// }
