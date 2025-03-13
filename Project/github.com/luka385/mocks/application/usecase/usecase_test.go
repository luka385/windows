package usecase

import (
	"errors"
	"primer-api/domain"
	mocks "primer-api/mock"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepositoryPort(ctrl)
	usecase := NewUseCase(mockRepo)

	user := &domain.User{
		ID:       "1",
		Email:    "lucasbrandan@gmail.com",
		Password: "3858",
	}

	mockRepo.EXPECT().CreateUser(user).Return(nil)

	err := usecase.CreateUser(user)

	assert.NoError(t, err)
}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepositoryPort(ctrl)
	usecase := NewUseCase(mockRepo)

	expectedUser := &domain.User{
		ID:       "1",
		Email:    "lucasbrandan@gmail.com",
		Password: "3858",
	}

	mockRepo.EXPECT().GetUserByID("1").Return(expectedUser, nil)

	user, err := usecase.GetUserByID("1")

	//POSIBLES CASOS DE ERROR

	//aqui testeo que no de error
	assert.NoError(t, err)

	//Aqui testeo si el usuario que me devuelve no es nulo
	assert.NotNil(t, user)

	//Aqui testeo si el usuario que me devuelve(user) sea igual que el usuario que espero(expectedUser)
	assert.Equal(t, expectedUser, user)
}

func TestCreateUser_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepositoryPort(ctrl)
	usecase := NewUseCase(mockRepo)

	user := &domain.User{
		ID:       "1",
		Email:    "lucasbrandan@gmail.com",
		Password: "3858",
	}

	mockRepo.EXPECT().CreateUser(user).Return(errors.New("ERROR USER NOT CREATED"))

	err := usecase.CreateUser(user)

	//CASOS DE ERROR

	// Valido que sea error
	assert.Error(t, err)

	//me aseguro que el error que me devolvio sea igual al error que yo espero
	assert.Equal(t, "ERROR USER NOT CREATED", err.Error())
}

func TestGetUserByID_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockRepositoryPort(ctrl)
	usecase := NewUseCase(mockRepo)
	mockRepo.EXPECT().GetUserByID("1").Return(nil, errors.New("USER NOT FOUND"))

	user, err := usecase.GetUserByID("1")

	//CASOS DE ERROR

	//Valido que sea error
	assert.Error(t, err)

	//verifico que el usuario sea nulo
	assert.Nil(t, user)

	// me aseguro q el mensaje que me devuelve sea igual al que yo espero
	assert.Equal(t, "USER NOT FOUND", err.Error())
}
