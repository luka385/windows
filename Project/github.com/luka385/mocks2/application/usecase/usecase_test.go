package usecase_test

import (
	"errors"
	"testing"

	"Project/github.com/luka385/mocks2/application/usecase"
	"Project/github.com/luka385/mocks2/domain"
	mocks "Project/github.com/luka385/mocks2/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserUseCase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockRepo := mocks.NewMockRepositoryPort(ctrl)
	userUseCase := usecase.NewUserUseCase(MockRepo)

	test := []struct {
		name       string       // nombre del caso
		method     string       // nombre del metodo que se usa
		userID     string       // user id para cada caso donde solo se usa en UPDATE, GETBYID Y DELETE
		inputUser  *domain.User // este field se utiliza para enviar un usuario con el caso de CREATE Y UPDATE
		mockreturn interface{}  //se utiliza interface para almacenar valor de cualquier tipo
		mockErr    error        // el tipo de error de cada caso
		expectErr  bool         // se utiliza para verificar si el caso sera con error o no
	}{
		//CreateUser
		{
			name:       "CreateUser - usuario creado con exito",
			method:     "CreateUser",
			inputUser:  &domain.User{Id: "2", Email: "pancho@gmail.com", Password: "3858"},
			mockreturn: nil,
			mockErr:    nil,
		},
		//CreateUser_Error
		{
			name:       "CreateUser - usuario no creado con exito",
			method:     "CreateUser",
			inputUser:  &domain.User{Id: "FailID", Email: "Fail@gmail.com", Password: "FailPassword"},
			mockreturn: nil,
			mockErr:    errors.New("Error Usuario no creado"),
			expectErr:  true,
		},
		//GetUsers
		{
			name:   "GetUsers - usuarios encontrados",
			method: "GetUsers",
			mockreturn: []*domain.User{
				{Id: "1", Email: "lucas@gmail.com", Password: "123"},
				{Id: "2", Email: "pancho@gmail.com", Password: "321"},
			},
			mockErr: nil,
		},
		//GetUser-Error
		{
			name:       "GetUsers - usuarios no encontrados",
			method:     "GetUsers",
			mockreturn: nil,
			mockErr:    errors.New("Usuarios no encontrados"),
			expectErr:  true,
		},
		// GetUserByID
		{
			name:       "GetUserByID - usuario encontrado",
			method:     "GetUserByID",
			userID:     "1",
			mockreturn: &domain.User{Id: "1", Email: "lucas@gmail.com", Password: "123"},
			mockErr:    nil,
		},
		//GetUserByID-Error
		{
			name:       "GetUserByID - usuario no encontrado",
			method:     "GetUserByID",
			userID:     "99",
			mockreturn: nil,
			mockErr:    errors.New("usuario no encontrado"),
			expectErr:  true,
		},

		//UpdateUser
		{
			name:       "UpdateUser - usuario actualizado",
			method:     "UpdateUser",
			userID:     "2",
			inputUser:  &domain.User{Email: "pancho2@gmail.com", Password: "3221"},
			mockreturn: nil,
			mockErr:    nil,
		},
		//UpdateUser_Error
		{
			name:       "UpdateUser - usuario no actualizado",
			method:     "UpdateUser",
			userID:     "2",
			inputUser:  &domain.User{Email: "errorUpdate", Password: "errorUpdate"},
			mockreturn: nil,
			mockErr:    errors.New("Usuario no actualizado"),
			expectErr:  true,
		},
		//DeleteUser
		{
			name:       "DeleteUser - usuario borrado",
			method:     "DeleteUser",
			userID:     "2",
			mockreturn: nil,
			mockErr:    nil,
		},
		//DeleteUser_Error
		{
			name:       "DeleteUser - usuario no borrado",
			method:     "DeleteUser",
			userID:     "2",
			mockreturn: nil,
			mockErr:    errors.New("Usuario no borrado correctamente"),
			expectErr:  true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.method {
			case "CreateUser":
				MockRepo.EXPECT().CreateUser(tt.inputUser).Return(tt.mockErr)

				err := userUseCase.CreateUser(tt.inputUser)

				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}

			case "GetUsers":

				var usersMock []*domain.User
				if tt.mockreturn != nil {
					usersMock = tt.mockreturn.([]*domain.User)
				}

				MockRepo.EXPECT().GetUsers().Return(usersMock, tt.mockErr)

				users, err := userUseCase.GetUsers()

				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.mockreturn, users)
				}

			case "GetUserByID":
				var userMock *domain.User
				if tt.mockreturn != nil {
					userMock = tt.mockreturn.(*domain.User)
				}

				MockRepo.EXPECT().GetUserById(tt.userID).Return(userMock, tt.mockErr)

				user, err := userUseCase.GetUserById(tt.userID)

				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, tt.mockreturn, user)
				}

			case "UpdateUser":
				MockRepo.EXPECT().UpdateUser(tt.userID, tt.inputUser).Return(tt.mockErr)

				err := userUseCase.UpdateUser(tt.userID, tt.inputUser)

				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}

			case "DeleteUser":
				MockRepo.EXPECT().DeleteUser(tt.userID).Return(tt.mockErr)

				err := userUseCase.DeleteUser(tt.userID)

				if tt.expectErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
				}
			}
		})
	}
}
