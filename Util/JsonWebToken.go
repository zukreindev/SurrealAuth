package util

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWTData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignToken(data JWTData) interface{} {

	if data.Username == "" || data.Email == "" || data.Password == "" {
		return "INVALID_DATA"
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": data.Username,
		"email":    data.Email,
		"password": data.Password,
		// "exp": time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(GetConfig("jwt", "secret")))

	if err != nil {
		panic(err)
	}

	return tokenString
}

func VerifyToken(tokenString string) interface{} {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(GetConfig("jwt", "secret")), nil
	})

	if err != nil {
		return err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return JWTData{
			Username: claims["username"].(string),
			Email:    claims["email"].(string),
			Password: claims["password"].(string),
		}
	} else {
		return "INVALID_TOKEN"
	}
}
