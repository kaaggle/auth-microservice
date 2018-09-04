package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/globalsign/mgo/bson"
)

type UserClaim struct {
	ID       bson.ObjectId `json:"_id"`
	Email    string        `json:"email"`
	Role     string        `json:"role"`
	Password string        `json:"password"`
	jwt.StandardClaims
}
