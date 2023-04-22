package usecase

import (
	"context"
	"testing"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	mock "github.com/akshayur04/project-ecommerce/pkg/repository/mockRepo"
	"github.com/golang/mock/gomock"
)

func TestUserSignUp(t *testing.T) {
	// Create a new controller for the mock objects
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a new mock repository and use it to create a new use case
	mockRepo := mock.NewMockUserRepository(ctrl)
	usecase := NewUserUseCase(mockRepo)

	// Create a new context and user request with all required fields
	ctx := context.Background()
	req := helperStruct.UserReq{
		Name:     "John Doe",
		Email:    "johndoe@example.com",
		Mobile:   "1234567890",
		Password: "password123",
	}

	// Create a new user data object with an ID of 1 and the same fields as the request
	expectedUserData := response.UserData{
		Id:     1,
		Name:   req.Name,
		Email:  req.Email,
		Mobile: req.Mobile,
	}

	// Expect the repository's UserSignUp method to be called with the given arguments
	mockRepo.EXPECT().UserSignUp(ctx, req).Return(expectedUserData, nil)

	// Call the UserSignUp method on the use case
	userData, err := usecase.UserSignUp(ctx, req)

	// Check that the returned data matches the expected data
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if userData != expectedUserData {
		t.Errorf("Unexpected user data. Expected: %v, but got: %v", expectedUserData, userData)
	}

	// // Try calling UserSignUp with an empty name field and verify that it returns an error
	// reqEmptyName := helperStruct.UserReq{
	// 	Name:     "John Doe",
	// 	Email:    "johndoe@example.com",
	// 	Mobile:   "1234567890",
	// 	Password: "password123",
	// }
	// _, err = usecase.UserSignUp(ctx, reqEmptyName)
	// if err == nil {
	// 	t.Error("Expected error due to empty name field, but got nil")
	// }

	// // Try calling UserSignUp with an invalid email address and verify that it returns an error
	// reqInvalidEmail := helperStruct.UserReq{
	// 	Name:     "John Doe",
	// 	Email:    "invalidemail",
	// 	Mobile:   "1234567890",
	// 	Password: "password123",
	// }
	// _, err = usecase.UserSignUp(ctx, reqInvalidEmail)
	// if err == nil {
	// 	t.Error("Expected error due to invalid email address, but got nil")
	// }

	// // Try calling UserSignUp with an empty mobile field and verify that it returns an error
	// reqEmptyMobile := helperStruct.UserReq{
	// 	Name:     "John Doe",
	// 	Email:    "johndoe@example.com",
	// 	Mobile:   "",
	// 	Password: "password123",
	// }
	// _, err = usecase.UserSignUp(ctx, reqEmptyMobile)
	// if err == nil {
	// 	t.Error("Expected error due to empty mobile field, but got nil")
	// }
}
