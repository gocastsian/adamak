package contract

import (
	"context"
	"github.com/gocastsian/adamak/entity"
)

//go:generate mockgen -destination=../mocks/store/user_store_mock.go -package=store_mock  -self_package=github.com/gocastsian/adamak/contract github.com/gocastsian/adamak/contract UserStore

type UserStore interface {
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
	GetUser(ctx context.Context, userID uint) (entity.User, error)
	UpdateUser(ctx context.Context, user entity.User) (entity.User, error)
	FindUsers(ctx context.Context) ([]entity.User, error)
	DeleteUser(ctx context.Context, userID uint) error
}