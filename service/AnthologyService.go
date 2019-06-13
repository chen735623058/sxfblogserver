package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sxfblogserver/dao"
	"sxfblogserver/domain"
	"github.com/satori/go.uuid"
)

// 添加文集
func AddAnthology(c *gin.Context)  {
	var rt = false;
	var anthology domain.Anthology
	err := c.BindJSON(&anthology)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	} else {
		_uuid,err := uuid.NewV4()
		if err != nil{
			c.JSON(200, gin.H{"errcode": 400, "description": "生成UUid出错"})
		}
		anthology.Id = _uuid.String()
		rt = dao.AddanthologyDao(anthology)
	}
	if rt == true {
		c.String(http.StatusOK,"添加成功")
	}else {
		c.String(http.StatusOK,"添加失败")
	}
}
// 删除文集
func DelAnthology(c *gin.Context)  {
	//Anthologyid := c.Query("id")
	Anthologyid, exist := c.GetQuery("id")
	if !exist {
		c.String(http.StatusOK,"没有找到相应的ID")
		return
	}
	rt := dao.DeleteanthologylistDao(Anthologyid)
	if rt == true {
		c.String(http.StatusOK,"删除成功")
	}else{
		c.String(http.StatusOK,"删除失败")
	}
}
// 修改文集
func UpdateAnthology(c *gin.Context)  {
	var anthology domain.Anthology
	err := c.BindJSON(&anthology)
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}
	rt  := dao.UpdateanthologylistDao(anthology)
	if rt == true {
		c.String(http.StatusOK,"修改成功")
	}else {
		c.String(http.StatusOK,"修改失败")
	}
}
// 查询文集
func GetAnthologys(c *gin.Context)  {

	rt , err := dao.FindanthologylistDao()
	if err != nil {
		c.JSON(200, gin.H{"errcode": 400, "description": "Post Data Err"})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"stauts":200,
		"error":nil,
		"data":rt,
	})
}
