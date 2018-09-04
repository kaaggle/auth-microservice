package repository

import (
	"errors"
	"schoolsystem/auth-microservice/auth"
	"schoolsystem/auth-microservice/models"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type mongoAuthRepo struct {
	Session        *mgo.Session
	UserCollection *mgo.Collection
}

func NewMongoAuthRepository(session *mgo.Session) auth.AuthRepository {

	return &mongoAuthRepo{
		Session:        session,
		UserCollection: session.DB("school-system").C("users"),
	}
}

func (m *mongoAuthRepo) Login(email, password string) (bool, error) {
	u := models.User{}
	err := m.UserCollection.Find(bson.M{"email": email}).One(&u)

	if err != nil {
		return false, errors.New("User with this email not found")
	}

	// check the password hash
	isSame := auth.CheckPasswordHash(u.Password, password)

	if !isSame {
		return false, errors.New("Wrong password")
	}

	return true, nil
}

func (m *mongoAuthRepo) Signup(u *models.User) (*models.User, error) {
	hashedPassword, err := auth.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = hashedPassword
	err = m.UserCollection.Insert(&u)

	if err != nil {
		return nil, err
	}

	return u, nil
}
