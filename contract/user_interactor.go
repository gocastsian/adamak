package contract

import (
	"context"
	"github.com/gocastsian/adamak/dto"
)

type UserInteractor interface {
	CreateUser(context.Context, dto.CreateUserRequest) (dto.CreateUserResponse, error)
	UpdateUser(context.Context, dto.UpdateUserRequest) (dto.UpdateUserResponse, error)
	FindUsers(context.Context, dto.FindUsersRequest) (dto.FindUsersResponse, error)
	FindUser(context.Context, dto.FindUserRequest) (dto.FindUserResponse, error)
	DeleteUser(context.Context, dto.DeleteUserRequest) (dto.DeleteUserResponse, error)
}
