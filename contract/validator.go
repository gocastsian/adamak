package contract

import (
	"context"
	"github.com/gocastsian/adamak/dto"
)

type (
	ValidateCreateUser func (req dto.CreateUserRequest) error
	ValidateUpdateUser func (ctx context.Context, req dto.UpdateUserRequest) error
	ValidateFindUser func (ctx context.Context, req dto.FindUserRequest) error
	ValidateDeleteUser func (ctx context.Context, req dto.DeleteUserRequest) error
)

