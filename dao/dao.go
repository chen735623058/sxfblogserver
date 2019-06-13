package dao
import (
	"gopkg.in/mgo.v2"
)

func getCollect(session *mgo.Session, collectionName string) *mgo.Collection {
	return session.DB("sxfblog").C(collectionName)
}