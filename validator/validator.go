package validator

import (
	"context"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gocastsian/adamak/adapter/store"
	"github.com/gocastsian/adamak/contract"
	"github.com/gocastsian/adamak/dto"
)

func ValidateCreateUser(req dto.CreateUserRequest) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.Name, validation.Required),
		validation.Field(&req.Email, validation.Required, is.Email),
		validation.Field(&req.Password, validation.Required),
		)
}

func ValidateUpdateUser(store store.MySQLStore) contract.ValidateUpdateUser {
	return func (ctx context.Context, req dto.UpdateUserRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.ID, validation.By(doesUserExist(ctx, store))),
			validation.Field(&req.Name, validation.Required),
			)
	}
}

func ValidateFindUser(store store.MySQLStore) contract.ValidateFindUser {
	return func (ctx context.Context, req dto.FindUserRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.ID, validation.By(doesUserExist(ctx, store))),
		)
	}
}

func ValidateDeleteUser(store store.MySQLStore) contract.ValidateDeleteUser {
	return func (ctx context.Context, req dto.DeleteUserRequest) error {
		return validation.ValidateStruct(&req,
			validation.Field(&req.ID, validation.By(doesUserExist(ctx, store))),
		)
	}
}