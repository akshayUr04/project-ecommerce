package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserAuthorizationMiddleware(c *gin.Context) {
	// s := c.Request.Header.Get("Authorization")
	tokenString, err := c.Cookie("UserAuth")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err := ValidateToken(tokenString); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func ValidateToken(token string) error {
	Tokenvalue, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if claims, ok := Tokenvalue.Claims.(jwt.MapClaims); ok && Tokenvalue.Valid {
		//Check the expir
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return fmt.Errorf("token expired")
		}
		// fmt.Println(claims["exp"])
	}

	return err
}
