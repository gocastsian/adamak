package store

import "github.com/gocastsian/adamak/entity"

type User struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func mapFromUserEntity(user entity.User) User {
	return User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}

func mapToUserEntity(user User) entity.User {
	return entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
