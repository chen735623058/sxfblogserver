package dao

import (
	"gopkg.in/mgo.v2"
	"sxfblogserver/domain"

	"gopkg.in/mgo.v2/bson"
	"log"
)


func  getCol(session *mgo.Session) *mgo.Collection {
	return getCollect(session, "Anthology")
}

func AddanthologyDao(anthology domain.Anthology ) bool {
	mongo, err := mgo.Dial("127.0.0.1") // 建立连接
	defer mongo.Close()
	if err != nil {
		return false
	}
	client := getCol(mongo)
	//创建数据
	data := anthology
	//插入数据
	cErr := client.Insert(&data)
	if cErr != nil {
		return false
	}
	return true
}

func FindanthologylistDao() (t []domain.Anthology, err error) {
	mongo,err := mgo.Dial("127.0.0.1")
	defer mongo.Close()
	if err != nil{
		return
	}
	client := getCol(mongo)
	client.Find(nil).All(&t)
	return
}

func UpdateanthologylistDao(anthology domain.Anthology) bool {
	mongo,err := mgo.Dial("127.0.0.1")
	defer mongo.Close()
	if err != nil{
		return false
	}
	client := getCol(mongo)
	err = client.Update(bson.M{"id":anthology.Id},bson.M{"$set":bson.M{"name":anthology.Name}})
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func DeleteanthologylistDao(_id string) bool  {
	mongo, err := mgo.Dial("127.0.0.1")
	defer mongo.Close()
	if err != nil {
		return false
	}
	client := getCol(mongo)
	err = client.Remove(bson.M{"id":_id})
	if err != nil {
		return  false
	}
	return  true

}