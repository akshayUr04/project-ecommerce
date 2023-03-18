package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/akshayur04/project-ecommerce/pkg/common/helperStruct"
	"github.com/akshayur04/project-ecommerce/pkg/common/response"
	interfaces "github.com/akshayur04/project-ecommerce/pkg/repository/interface"
	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

// --------------------------------------------------UserSignUp------------------------------------------------------------

func (c *userUseCase) UserSignUp(ctx context.Context, user helperStruct.UserReq) (response.UserData, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return response.UserData{}, err
	}
	user.Password = string(hash)
	userData, err := c.userRepo.UserSignUp(ctx, user)
	return userData, err
}

// --------------------------------------------------UserLogin------------------------------------------------------------

func (c *userUseCase) UserLogin(ctx context.Context, user helperStruct.LoginReq) (string, error) {
	userData, err := c.userRepo.UserLogin(ctx, user.Email)
	if err != nil {
		return "", err
	}

	if user.Email == "" {
		return "", fmt.Errorf("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	if userData.IsBlocked {
		return "", fmt.Errorf("user is blocked")
	}

	claims := jwt.MapClaims{
		"id":  userData.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (c *userUseCase) IsSignIn(phno string) (bool, error) {
	isSignin, err := c.userRepo.IsSignIn(phno)
	return isSignin, err
}

func (c *userUseCase) OtpLogin(phno string) (string, error) {
	id, err := c.userRepo.OtpLogin(phno)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (c *userUseCase) AddAddress(id int, address helperStruct.Address) error {
	err := c.userRepo.AddAddress(id, address)
	return err
}

func (c *userUseCase) UpdateAddress(id, addressId int, address helperStruct.Address) error {
	err := c.userRepo.UpdateAddress(id, addressId, address)
	return err
}

func (c *userUseCase) Viewprfile(id int) (response.UserData, error) {
	profile, err := c.userRepo.Viewprfile(id)
	return profile, err
}

func (c *userUseCase) UserEditProfile(id int, updatingDetails helperStruct.UserReq) (response.UserData, error) {
	updatedProfile, err := c.userRepo.UserEditProfile(id, updatingDetails)
	return updatedProfile, err
}

func (c *userUseCase) UpdatePassword(id int, Passwords helperStruct.UpdatePassword) error {

	orginalPassword, err := c.userRepo.FindPassword(id)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(orginalPassword), []byte(Passwords.OldPassword))
	if err != nil {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(Passwords.NewPasswoed), 10)
	if err != nil {
		return err
	}
	newPassword := string(hash)

	err = c.userRepo.UpdatePassword(id, newPassword)
	return err
}
