package validator

import (
	"context"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gocastsian/adamak/contract"
)

func doesUserExist(ctx context.Context, store contract.ValidatorStore) validation.RuleFunc {
	return func(value interface{}) error {
		userID := value.(uint)

		ok, err := store.DoesUserExist(ctx, userID)
		if err != nil {
			return fmt.Errorf("%v", err)
		}

		if !ok {
			return fmt.Errorf("user: %d does not exist", userID)
		}

		return nil
	}
}

