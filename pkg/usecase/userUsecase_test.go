package usecase

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	"github.com/akshayur04/project-ecommerce/pkg/repository/mockRepo"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"golang.org/x/crypto/bcrypt"
)

type eqCreateUserParamsMatcher struct {
	arg      helperStruct.UserReq
	password string
}

func (e eqCreateUserParamsMatcher) Matches(x interface{}) bool {
	arg, ok := x.(helperStruct.UserReq)
	if !ok {
		return false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(arg.Password), []byte(e.password)); err != nil {
		return false
	}
	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamsMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParams(arg helperStruct.UserReq, password string) gomock.Matcher {
	return eqCreateUserParamsMatcher{arg, password}
}

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
				userRepo.EXPECT().UserSignUp(gomock.Any(),
					EqCreateUserParams(helperStruct.UserReq{
						Name:     "akshay",
						Email:    "akshay@gmail.com",
						Mobile:   "9072001341",
						Password: "123456789",
					},
						"123456789")).
					Times(1).
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
		{
			name: "alredy exits",
			input: helperStruct.UserReq{
				Name:     "akshay",
				Email:    "akshay@gmail.com",
				Mobile:   "9072001341",
				Password: "123456789",
			},
			buildStub: func(userRepo mockRepo.MockUserRepository) {
				userRepo.EXPECT().UserSignUp(gomock.Any(),
					EqCreateUserParams(helperStruct.UserReq{
						Name:     "akshay",
						Email:    "akshay@gmail.com",
						Mobile:   "9072001341",
						Password: "123456789",
					},
						"123456789")).
					Times(1).
					Return(response.UserData{},
						errors.New("user alredy exits"))
			},
			expectedOutput: response.UserData{},
			expectedError:  errors.New("user alredy exits"),
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
