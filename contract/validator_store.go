package contract

import "context"

type ValidatorStore interface {
	DoesUserExist(ctx context.Context, userID uint) (bool, error)
}
