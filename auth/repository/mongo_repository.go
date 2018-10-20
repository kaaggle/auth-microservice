package repository

import (
	"church-adoration/auth"
	"church-adoration/models"
	"errors"

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
		UserCollection: session.DB("church-adoration").C("users"),
	}
}

func (m *mongoAuthRepo) Login(email, password string) (*models.User, string, error) {
	u := models.User{}
	err := m.UserCollection.Find(bson.M{"email": email}).One(&u)

	if err != nil {
		return nil, "", errors.New("User with this email not found.")
	}

	if !u.Verified {
		return nil, "", errors.New("User is not verified.")
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

func (m *mongoAuthRepo) Signup(u *models.User) (*models.User, string, error) {
	// check if user already exists with this email
	user := models.User{}
	err := m.UserCollection.Find(bson.M{"email": u.Email}).One(&user)
	if err == nil || user.Email != "" {
		return nil, "", errors.New("User already registered")
	}



	hashedPassword, err := auth.HashPassword(u.Password)
	if err != nil {
		return nil, "", err
	}
	u.Password = hashedPassword

	upsertedUser, err := m.UserCollection.Upsert(nil, u)

	if err != nil {
		return nil, "", err
	}

	return u, upsertedUser.UpsertedId.(bson.ObjectId).Hex(), nil
}

func (m *mongoAuthRepo) ActivateUser(id string) error {
	err := m.UserCollection.Update(bson.M{"_id": bson.ObjectIdHex(id)}, bson.M{"$set": bson.M{"verified": true}})

	if err != nil {
		return err
	}

	return nil
}
