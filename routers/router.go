package routers

import (
	"github.com/gin-gonic/gin"
	"project/middleware/jwt"
	"project/pkg/setting"
	"project/routers/api"
	v1 "project/routers/v1"
)

func InitRoute() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/auth", api.GetAuth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		// 获取文章tag
		apiV1.GET("/tags", v1.GetTag)
		// 添加文章tag
		apiV1.POST("/tags", v1.AddTag)
		// 修改文章tag
		apiV1.PUT("/tags/:id", v1.EditTag)
		// 删除文章tag
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
		// 获取指定文章
		apiV1.GET("/article/:id", v1.GetArticle)
		// 获取文章列表
		apiV1.GET("/articles", v1.GetArticles)
		// 添加文章
		apiV1.POST("/article", v1.AddArticle)
		// 修改文章
		apiV1.PUT("/article/:id", v1.EditArticle)
		// 删除文章
		apiV1.DELETE("/article/:id", v1.DeletedArticle)
	}
	return r
}
