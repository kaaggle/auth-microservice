package authorization

import "github.com/dgrijalva/jwt-go"

var secret = []byte("MySECRET")

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
