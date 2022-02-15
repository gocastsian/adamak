package store

import (
	"context"
	"gorm.io/gorm"
)

func (m MySQLStore) DoesUserExist(ctx context.Context, userID uint) (bool, error) {
	if err := m.db.WithContext(ctx).Where("id = ?", userID).First(&User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
