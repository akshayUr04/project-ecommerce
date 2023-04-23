package usecase

import (
	"context"
	"testing"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	"github.com/akshayur04/project-ecommerce/pkg/repository/mockRepo"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestUserSignup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepo := mockRepo.NewMockUserRepository(ctrl)
	userUseCase := NewUserUseCase(userRepo)
	testData := []struct {
		name           string
		input          helperStruct.UserReq
		buildStub      func(userRepo mockRepo.MockUserRepository)
		expectedOutput response.UserData
		expectedError  error
	}{
		{
			name: "new user",
			input: helperStruct.UserReq{
				Name:     "akshay",
				Email:    "akshay@gmail.com",
				Mobile:   "9072001341",
				Password: "123456789",
			},
			buildStub: func(userRepo mockRepo.MockUserRepository) {
				userRepo.EXPECT().UserSignUp(gomock.Any(), helperStruct.UserReq{
					Name:     "akshay",
					Email:    "akshay@gmail.com",
					Mobile:   "9072001341",
					Password: "123456789",
				}).Times(1).
					Return(response.UserData{
						Id:     1,
						Name:   "akshay",
						Email:  "akshay@gmail.com",
						Mobile: "9072001341",
					}, nil)
			},
			expectedOutput: response.UserData{
				Id:     1,
				Name:   "akshay",
				Email:  "akshay@gmail.com",
				Mobile: "9072001341",
			},
			expectedError: nil,
		},
	}
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*userRepo)
			actualUser, err := userUseCase.UserSignUp(context.TODO(), tt.input)
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, actualUser, tt.expectedOutput)
		})
	}

}
