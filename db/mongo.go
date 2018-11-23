package db

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
)

// MongoDB 's structure.
type MongoDB struct {
	session Session
	dbname  string
}

// Session of MongoDB.
type Session struct {
	session *mgo.Session
}

// NewMongoDB creates MongoDB instance.
func NewMongoDB(dbserver, dbname, dbuser, dbpass string, timeout int) (*MongoDB, error) {
	mongodb := new(MongoDB)
	mongodb.dbname = dbname

	// Setup MongoDB
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{dbserver},
		Timeout:  time.Duration(timeout) * time.Second,
		Database: dbname,
		Username: dbuser,
		Password: dbpass,
	}
	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return nil, fmt.Errorf("Unable to dial Mongo DB: %v", err)
	}
	s := &Session{session}
	mongodb.SetSession(s)
	return mongodb, nil
}

// SetSession alters the session reside inside the MongoDB instance.
func (m *MongoDB) SetSession(session *Session) {
	m.session = *session
}
