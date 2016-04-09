package controllers

import (
	"log"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var songs *mgo.Collection
var users *mgo.Collection

func init() {
	var err error
	session, err = mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	songs = session.DB("bins").C("songs")
	users = session.DB("bins").C("users")
}
