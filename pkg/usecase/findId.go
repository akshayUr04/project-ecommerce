package usecase

import (
	"fmt"

	services "github.com/akshayur04/project-ecommerce/pkg/usecase/interface"
	"github.com/golang-jwt/jwt"
)

type FindIdUseCase struct {
}

func NewFindIdUseCase() services.FindIdUseCase {
	return &FindIdUseCase{}
}

func (c *FindIdUseCase) FindId(cookie string) (int, error) {
	Tokenvalue, err := jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("secret"), nil
	})
	if err != nil {
		return 0, err
	}
	var id interface{}
	if claims, ok := Tokenvalue.Claims.(jwt.MapClaims); ok && Tokenvalue.Valid {
		id = claims["id"]
	}

	n, ok := id.(float64)
	if !ok {
		return 0, fmt.Errorf("expected an int value, but got %T", id)
	}

	v := int(n)

	return v, nil
}
