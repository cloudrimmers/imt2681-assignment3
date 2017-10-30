package database

import (
	"os"

	"gopkg.in/mgo.v2"
)

var mongoURI string = os.Getenv("MONGODB_URI")
var mongoDB string = os.Getenv("MONGODB_NAME")
var session *mgo.Session
var err error

// Open ...
func Open() (*mgo.Database, error) {
	database := &mgo.Database{}
	session, err = mgo.Dial(mongoURI)
	if err == nil {
		database = session.DB(mongoDB)
	}
	return database, err
}

// Close ...
func Close() {
	session.Close()
}