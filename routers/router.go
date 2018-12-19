package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ivanberry/amy-blog/pkg"
	"github.com/ivanberry/amy-blog/routers/api/v1"



	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/ivanberry/amy-blog/docs"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)


		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticles)

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	return r
}