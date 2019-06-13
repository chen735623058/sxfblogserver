package dao

import (
	"gopkg.in/mgo.v2"
	"sxfblogserver/domain"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func getColArticle(session *mgo.Session) *mgo.Collection  {
	return  getCollect(session,"Article")
}

func AddArticleDao(article domain.Article) bool  {
	mongo, err := mgo.Dial("127.0.0.1")
	defer mongo.Close()
	if err != nil {
		return false
	}
	client := getColArticle(mongo)
	// 创建数据
	data := article
	// 插入数据
	cErr := client.Insert(&data)
	if cErr != nil {
		return false
	}
	return  true

}

func DelArticleDao( _id string) bool  {
	mongo, err := mgo.Dial("127.0.0.1")
	defer mongo.Close()
	if err != nil {
		return false
	}
	client := getColArticle(mongo)
	err = client.Remove(bson.M{"id":_id})
	if err != nil {
		return  false
	}
	return  true
}

func UpdateArticleDao(article domain.Article) bool  {
	mongo,err := mgo.Dial("127.0.0.1")
	defer mongo.Close()
	if err != nil{
		return false
	}
	client := getColArticle(mongo)
	err = client.Update(
		bson.M{"id":article.Id},
		article,
	)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}



func GetArticleListDao() (t []domain.Article, err error) {
	mongo,err := mgo.Dial("127.0.0.1")
	defer mongo.Close()
	if err != nil{
		return
	}
	client := getColArticle(mongo)
	client.Find(nil).Select(bson.M{"content": 0}).All(&t)
	return
}

func GetArticleInfoDao( _id string) (t domain.Article, err error)  {
	mongo, err := mgo.Dial("127.0.0.1")
	defer mongo.Close()
	if err != nil {
		return
	}
	client := getColArticle(mongo)
	err = client.Find(bson.M{"id": _id}).One(&t)

	return
}


