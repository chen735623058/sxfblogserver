package main
import (
	"github.com/gin-gonic/gin"
	"sxfblogserver/service"
)
func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 文集管理
	router.GET("/getAnthology", service.GetAnthologys)
	router.GET("/delAnthology", service.DelAnthology)
	router.POST("/updateAnthology", service.UpdateAnthology)
	router.POST("/addAnthology", service.AddAnthology)
	// 文章管理

	router.GET("/delArticle", service.DelArticle)
	router.POST("/updateArticle", service.UpdateArticle)
	router.POST("/addArticle", service.AddArticle)
	router.GET("/getArticleList", service.GetArticleList)
	router.GET("/getArticleInfo", service.GetArticleInfo)

	router.Run(":9999")
}
