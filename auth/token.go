package auth

import (
	"github.com/dgrijalva/jwt-go"
)

var secret []byte = []byte("MySECRET")

func CreateTokenString(u UserClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &u)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetToken(tokenString string) *jwt.Token {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	return token
}

func GetRoleFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return "", err
	}

	role := token.Claims.(jwt.MapClaims)["role"]

	return role.(string), nil
}

func GetUserIdFromToken(tokenString string) (float64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return 0, err
	}

	userId := token.Claims.(jwt.MapClaims)["_id"]

	return userId.(float64), nil
}
