package user

import (
	"context"

	"gorm.io/gorm"

	"github.com/gocastsian/adamak/protocol"
)

type storage_MySQL struct {
	db *gorm.DB
}

func (m *storage_MySQL) init(config *protocol.ServiceInitializeConfig) {
	m.db = config.MySQL.DB
	var err = m.db.AutoMigrate(&protocol.User{})
	if err != nil {
		panic("[User Microservice] Failed to auto migrate database!")
	}
}

func (m *storage_MySQL) CreateUser(ctx context.Context, user protocol.User) (protocol.User, error) {
	var err = m.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return protocol.User{}, err
	}
	return user, nil
}

func (m *storage_MySQL) GetUser(ctx context.Context, userID uint) (user protocol.User, err error) {
	err = m.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return protocol.User{}, nil
	}
	return user, nil
}

func (m *storage_MySQL) UpdateUser(ctx context.Context, user protocol.User) (protocol.User, error) {
	var err = m.db.WithContext(ctx).Save(&user).Error
	if err != nil {
		return protocol.User{}, err
	}
	return user, nil
}

func (m *storage_MySQL) FindUsers(ctx context.Context) (users []protocol.User, err error) {
	err = m.db.WithContext(ctx).Find(&users).Error
	return
}

func (m *storage_MySQL) DeleteUser(ctx context.Context, userID uint) error {
	var user protocol.User
	var err = m.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return err
	}

	err = m.db.WithContext(ctx).Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *storage_MySQL) DoesUserExist(ctx context.Context, userID uint) (bool, error) {
	if err := m.db.WithContext(ctx).Where("id = ?", userID).First(&protocol.User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
