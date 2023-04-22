package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	"github.com/akshayur04/project-ecommerce/pkg/usecase/mockUsecase"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserSignUp(t *testing.T) {
	ctrl := gomock.NewController(t)
	userUseCase := mockUsecase.NewMockUserUseCase(ctrl)
	cartUseCase := mockUsecase.NewMockCartUsecase(ctrl)
	UserHandler := NewUserHandler(userUseCase, cartUseCase)

	testData := []struct {
		name             string
		userData         helperStruct.UserReq
		buildStub        func(userUsecase mockUsecase.MockUserUseCase)
		expectedCode     int
		expectedResponse response.Response
		expectedData     response.UserData
		expectedError    error
	}{
		{
			name: "new user",
			userData: helperStruct.UserReq{
				Name:     "akshay",
				Email:    "akshay@gmail.com",
				Mobile:   "9072001341",
				Password: "123456789",
			},
			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
				userUsecase.EXPECT().UserSignUp(gomock.Any(), helperStruct.UserReq{
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
				cartUseCase.EXPECT().CreateCart(1).Times(1).Return(nil)
			},
			expectedCode: 201,
			expectedResponse: response.Response{
				StatusCode: 201,
				Message:    "user signup Successfully",
				Data: response.UserData{
					Id:     1,
					Name:   "akshay",
					Email:  "akshay@gmail.com",
					Mobile: "9072001341",
				},
				Errors: nil,
			},
			expectedData: response.UserData{
				Id:     1,
				Name:   "akshay",
				Email:  "akshay@gmail.com",
				Mobile: "9072001341",
			},
			expectedError: nil,
		},
		{
			name: "duplicate user",
			userData: helperStruct.UserReq{
				Name:     "akshay",
				Email:    "akshay@gmail.com",
				Mobile:   "9072001341",
				Password: "123456789",
			},
			buildStub: func(userUsecase mockUsecase.MockUserUseCase) {
				userUsecase.EXPECT().UserSignUp(gomock.Any(), helperStruct.UserReq{
					Name:     "akshay",
					Email:    "akshay@gmail.com",
					Mobile:   "9072001341",
					Password: "123456789",
				}).Times(1).
					Return(
						response.UserData{},
						errors.New("user already exists"),
					)
				// cartUseCase.EXPECT().CreateCart(1).Times(1).Return(nil)
			},
			expectedCode: 400,
			expectedResponse: response.Response{
				StatusCode: 400,
				Message:    "unable signup",
				Data:       response.UserData{},
				Errors:     "user already exits",
			},
			expectedData:  response.UserData{},
			expectedError: errors.New("user already exists"),
		},
	}
	for _, tt := range testData {
		t.Run(tt.name, func(t *testing.T) {
			tt.buildStub(*userUseCase)
			engine := gin.Default()
			recorder := httptest.NewRecorder()
			engine.POST("/user/signup", UserHandler.UserSignUp)
			var body []byte
			fmt.Println(tt.userData)
			body, err := json.Marshal(tt.userData)
			assert.NoError(t, err)
			url := "/user/signup"
			req := httptest.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			engine.ServeHTTP(recorder, req)
			var actual response.Response
			err = json.Unmarshal(recorder.Body.Bytes(), &actual)
			assert.NoError(t, err)
			assert.Equal(t, tt.expectedCode, recorder.Code)
			assert.Equal(t, tt.expectedResponse.Message, actual.Message)

			fmt.Printf("type of actual data %t \n %v", actual.Data, actual.Data)
			data, ok := actual.Data.(map[string]interface{})
			if ok {
				userData := response.UserData{
					Id:     int(data["id"].(float64)),
					Name:   data["name"].(string),
					Email:  data["email"].(string),
					Mobile: data["mobile"].(string),
				}
				if !reflect.DeepEqual(tt.expectedData, userData) {
					t.Errorf("got %q, but want %q", userData, tt.expectedData)
				}
			} else {
				t.Errorf("actual.Data is not of type map[string]interface{}")
			}

		})
	}
}
