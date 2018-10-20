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

func (m *mongoAuthRepo) SchoolRegistration(s *models.School) (*models.School, error) {
	err := m.UserCollection.Insert(&s)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (m *mongoAuthRepo) Login(email, password string) (*models.User, string, error) {
	u := models.User{}
	err := m.UserCollection.Find(bson.M{"email": email}).One(&u)

	if err != nil {
		return nil, "", errors.New("User with this email not found")
	}

	// check the password hash
	isSame := auth.CheckPasswordHash(password, u.Password)
	if !isSame {
		return nil, "", errors.New("Wrong password")
	}

	userClaim := auth.UserClaim{
		Email: email,
		Role:  u.Role,
	}
	token, err := auth.CreateTokenString(userClaim)
	if err != nil {
		return &u, "", err
	}
	return &u, token, nil
}

func (m *mongoAuthRepo) Signup(u *models.User) (*models.User, error) {
	// check if user already exists with this email
	user := models.User{}
	err := m.UserCollection.Find(bson.M{"email": u.Email}).One(&user)
	if err == nil {
		return nil, errors.New("User already registered")
	}

	hashedPassword, err := auth.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	u.Password = hashedPassword
	u.UserID = auth.GenerateUserID()

	err = m.UserCollection.Insert(&u)

	if err != nil {
		return nil, err
	}

	return u, nil
}
