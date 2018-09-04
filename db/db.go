package db

import (
	"github.com/globalsign/mgo"
)

func NewDatabaseConnection(url string) (*mgo.Session, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	session.SetMode(mgo.Monotonic, true)

	return session, nil
}
