package main

import (
	"log"

	"github.com/globalsign/mgo"
)

func GetConnection() (session *mgo.Session, err error) {

	session, err = mgo.Dial(MONGO_HOST)

	if err != nil {
		log.Printf("[ERROR] could not connect to mongo, because: %v", err)
		return
	}
	return
}
