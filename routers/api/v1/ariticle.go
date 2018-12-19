package v1

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/ivanberry/amy-blog/pkg"
	"github.com/ivanberry/amy-blog/pkg/e"
	"github.com/ivanberry/amy-blog/pkg/util"
	"github.com/ivanberry/amy-blog/models"
)

func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0 , 1, "state").Message("状态值只允许1或1")
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId

		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		//
		//data["lists"] = (util.GetPage(c), setting.PageSize, maps)
	}



}

func AddArticle(c *gin.Context) {
	//PASS
}

func EditArticle(c *gin.Context) {
	//PASS
}

func DeleteArticle(c *gin.Context) {
	//PASS
}

