package user

import (
	"context"
	"github.com/gocastsian/adamak/contract"
	"github.com/gocastsian/adamak/dto"
	"github.com/gocastsian/adamak/entity"
)

type Interactor struct {
	store contract.UserStore
}

func New(store contract.UserStore) Interactor {
	return Interactor{store: store}
}

func (i Interactor) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	createdUser, err := i.store.CreateUser(ctx, user)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	return dto.CreateUserResponse{User: createdUser}, nil
}

func (i Interactor) UpdateUser(ctx context.Context, req dto.UpdateUserRequest) (dto.UpdateUserResponse, error) {
	user, err := i.store.GetUser(ctx, req.ID)
	if err != nil {
		return dto.UpdateUserResponse{}, err
	}

	user.Name = req.Name

	updatedUser, err := i.store.UpdateUser(ctx, user)
	if err != nil {
		return dto.UpdateUserResponse{}, err
	}

	return dto.UpdateUserResponse{User: updatedUser}, nil
}

func (i Interactor) FindUsers(ctx context.Context, _ dto.FindUsersRequest) (dto.FindUsersResponse, error) {
	users, err := i.store.FindUsers(ctx)
	if err != nil {
		return dto.FindUsersResponse{}, err
	}

	return dto.FindUsersResponse{Users: users}, nil
}

func (i Interactor) FindUser(ctx context.Context, req dto.FindUserRequest) (dto.FindUserResponse, error) {
	user, err := i.store.GetUser(ctx, req.ID)
	if err != nil {
		return dto.FindUserResponse{}, err
	}

	return dto.FindUserResponse{User: user}, nil
}

func (i Interactor) DeleteUser(ctx context.Context, req dto.DeleteUserRequest) (dto.DeleteUserResponse, error) {
	err := i.store.DeleteUser(ctx, req.ID)
	if err != nil {
		return dto.DeleteUserResponse{}, err
	}

	return dto.DeleteUserResponse{}, nil
}