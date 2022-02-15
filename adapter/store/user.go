package store

import (
	"context"
	"github.com/gocastsian/adamak/entity"
)

func (m MySQLStore) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	u := mapFromUserEntity(user)

	if err := m.db.WithContext(ctx).Create(&u).Error; err != nil {
		return entity.User{}, err
	}

	return mapToUserEntity(u), nil
}

func (m MySQLStore) GetUser(ctx context.Context, userID uint) (entity.User, error) {
	user := User{}
	if err := m.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		return entity.User{}, nil
	}

	return mapToUserEntity(user), nil
}

func (m MySQLStore) UpdateUser(ctx context.Context, user entity.User) (entity.User, error) {
	u := mapFromUserEntity(user)

	if err := m.db.WithContext(ctx).Save(&u).Error; err != nil {
		return entity.User{}, err
	}

	return mapToUserEntity(u), nil
}

func (m MySQLStore) FindUsers(ctx context.Context) ([]entity.User, error) {
	var users []User

	if err := m.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	userEntities := make([]entity.User, len(users))
	for i := range users {
		userEntities[i] = mapToUserEntity(users[i])
	}

	return userEntities, nil
}

func (m MySQLStore) DeleteUser(ctx context.Context, userID uint) error {
	var user User
	if err := m.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error; err != nil {
		return err
	}

	if err := m.db.WithContext(ctx).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}