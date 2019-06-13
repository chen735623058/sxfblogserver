package service

import (
	"github.com/gin-gonic/gin"
	"sxfblogserver/domain"
	"github.com/satori/go.uuid"
	"time"
	"sxfblogserver/dao"
	"net/http"
)

// 添加文章
func AddArticle(c *gin.Context)  {
	var rt = false
	var article domain.Article
	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}else {
		_uuid,err := uuid.NewV4()
		if err != nil{
			c.JSON(200, gin.H{"errcode": 400, "description": "生成UUid出错"})
		}
		article.Id = _uuid.String()
		currentTime:=time.Now()
		article.Createtime = currentTime.String()
		rt = dao.AddArticleDao(article);
	}
	if rt == true {
		c.String(http.StatusOK,"添加文章成功")
	}else{
		c.String(http.StatusOK,"添加文章失败")
	}
}

// 删除文章
func DelArticle(c *gin.Context)  {
	articleid, exist := c.GetQuery("id")
	if !exist {
		c.String(http.StatusOK,"没有找到对应的ID")
		return
	}
	rt := dao.DelArticleDao(articleid)
	if rt == true {
		c.String(http.StatusOK,"删除成功")
	}else {
		c.String(http.StatusOK, "删除失败")
	}
}
// 更新文章
func UpdateArticle(c *gin.Context)  {
	var article domain.Article
	err := c.BindJSON(&article)
	if err !=nil {
		c.JSON(200,gin.H{"errcode":400,"description":"Post Data Err"})
		return
	}
	currentTime:=time.Now()
	article.Updatetime = currentTime.String()
	rt := dao.UpdateArticleDao(article)
	if rt == true {
		c.String(http.StatusOK,"修改成功")
	}else{
		c.String(http.StatusOK,"修改失败")
	}
}


// 获取文章列表
func GetArticleList(c *gin.Context)  {
	rt , err := dao.GetArticleListDao()
	if err != nil {
		c.JSON(200,gin.H{"errcode":400,"description":"Post Data Err"})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"stauts":200,
		"error":nil,
		"data":rt,
	})
}

// 获取文章详情
func GetArticleInfo(c *gin.Context)  {
	articleid, exist := c.GetQuery("id")
	if !exist {
		c.String(http.StatusOK,"没有找到对应的ID")
		return
	}
	rt , err := dao.GetArticleInfoDao(articleid)
	if err != nil {
		c.JSON(200,gin.H{"errcode":400,"description":"Post Data Err"})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"stauts":200,
		"error":nil,
		"data":rt,
	})
}

