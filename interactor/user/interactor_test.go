package user_test

import (
	"context"
	"fmt"
	"github.com/bxcodec/faker/v3"
	"github.com/gocastsian/adamak/contract"
	"github.com/gocastsian/adamak/dto"
	"github.com/gocastsian/adamak/entity"
	"github.com/gocastsian/adamak/interactor/user"
	storemock "github.com/gocastsian/adamak/mocks/store"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func setup(t *testing.T) (contract.UserInteractor, *storemock.MockUserStore, func()) {

	mockCtrl := gomock.NewController(t)
	mockUserStore := storemock.NewMockUserStore(mockCtrl)

	service := user.New(mockUserStore)

	return service, mockUserStore, func(){
		mockCtrl.Finish()
	}
}

func TestCreateUser(t *testing.T) {
	t.Run("fail on store creation error", func(t *testing.T) {
		interactor, mockUserStore, teardown := setup(t)
		defer teardown()

		req := dto.CreateUserRequest{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: faker.Password(),
		}

		ctx := context.Background()

		mockUserStore.EXPECT().CreateUser(ctx, gomock.Any()).Return(entity.User{}, fmt.Errorf(""))

		_, err := interactor.CreateUser(ctx, req)
		assert.NotNil(t, err)
	})

	t.Run("successful", func(t *testing.T) {
		interactor, mockUserStore, teardown := setup(t)
		defer teardown()

		req := dto.CreateUserRequest{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: faker.Password(),
		}

		ctx := context.Background()

		passedUser := entity.User{
			ID:       0,
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
		}

		returnedUser := passedUser
		returnedUser.ID = uint(rand.Uint64())

		mockUserStore.EXPECT().CreateUser(ctx, passedUser).Return(returnedUser, nil)

		resp, err := interactor.CreateUser(ctx, req)
		assert.Nil(t, err)
		assert.EqualValues(t, returnedUser, resp.User)
	})
}

func TestFindUser(t *testing.T) {
	t.Run("fail on store get user error", func(t *testing.T) {
		interactor, mockUserStore, teardown := setup(t)
		defer teardown()

		req := dto.FindUserRequest{
			ID: 0,
		}

		ctx := context.Background()

		mockUserStore.EXPECT().GetUser(ctx, gomock.Any()).Return(entity.User{}, fmt.Errorf(""))

		_, err := interactor.FindUser(ctx, req)
		assert.NotNil(t, err)
	})

	t.Run("successful", func(t *testing.T) {
		interactor, mockUserStore, teardown := setup(t)
		defer teardown()

		req := dto.FindUserRequest{
			ID: uint(rand.Uint64()),
		}

		ctx := context.Background()

		returnedUser := entity.User{
			ID:       req.ID,
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: faker.Password(),
		}

		mockUserStore.EXPECT().GetUser(ctx, req.ID).Return(returnedUser, nil)

		resp, err := interactor.FindUser(ctx, req)
		assert.Nil(t, err)
		assert.EqualValues(t, returnedUser, resp.User)
	})
}

func TestDeleteUser(t *testing.T) {
	t.Run("fail on store delete user error", func(t *testing.T) {
		interactor, mockUserStore, teardown := setup(t)
		defer teardown()

		req := dto.DeleteUserRequest{
			ID: 0,
		}

		ctx := context.Background()

		mockUserStore.EXPECT().DeleteUser(ctx, gomock.Any()).Return(fmt.Errorf(""))

		_, err := interactor.DeleteUser(ctx, req)
		assert.NotNil(t, err)
	})

	t.Run("successful", func(t *testing.T) {
		interactor, mockUserStore, teardown := setup(t)
		defer teardown()

		req := dto.DeleteUserRequest{
			ID: uint(rand.Uint64()),
		}

		ctx := context.Background()

		mockUserStore.EXPECT().DeleteUser(ctx, req.ID).Return(nil)

		_, err := interactor.DeleteUser(ctx, req)
		assert.Nil(t, err)
	})
}