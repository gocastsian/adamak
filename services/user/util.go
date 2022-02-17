package user

import (
	"context"
	"fmt"
)

func doesUserExist(ctx context.Context, userID uint) error {
	ok, err := storage.DoesUserExist(ctx, userID)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("user: %d does not exist", userID)
	}
	return nil
}
