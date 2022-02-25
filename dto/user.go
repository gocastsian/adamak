package dto

import "github.com/gocastsian/adamak/entity"

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	User entity.User `json:"user"`
}

type UpdateUserRequest struct {
	ID uint `json:"id"`
	Name  string `json:"name"`
}

type UpdateUserResponse struct {
	User entity.User `json:"user"`
}

type FindUsersRequest struct {}

type FindUsersResponse struct {
	Users []entity.User `json:"users"`
}

type FindUserRequest struct {
	ID uint `json:"id"`
}

type FindUserResponse struct {
	User entity.User `json:"user"`
}

type DeleteUserRequest struct {
	ID uint `json:"id"`
}

type DeleteUserResponse struct {}


