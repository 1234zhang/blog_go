package v1

import (
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
	"project/models"
	"project/pkg/e"
	"project/pkg/setting"
	"project/pkg/util"
)

//获取文章tag
func GetTag(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}

	var status int = -1
	if params := c.Query("status"); params != "" {
		status = com.StrTo(params).MustInt()
		maps["status"] = status
	}

	code := e.SUCCESS
	data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}

// 添加文章tag
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人字符不能超过100")
	valid.Range(state, 0, 1, "state").Message("状态只能为0和1")

	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg": e.GetMsg(e.INVALID_PARAMS),
			"data": fmt.Sprintf("%v", valid.Errors),
		})
		return
	}
	code := e.SUCCESS
	if models.ExistTagByName(name) {
		code = e.ERROR_EXIST_TAG
	}
	models.AddTag(name, state, createdBy)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 修改文章tag
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("state 只能是0或者1")
	}
	valid.Required(id, "id").Message("id不能为空")
	valid.Required(name, "name").Message("标签名不能为空")
	valid.MaxSize(name, 100, "name").Message("标签最长不能超过100")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人名字长度不能超过100")

	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg": e.GetMsg(e.INVALID_PARAMS),
			"data": fmt.Sprintf("%v", valid.Errors),
		})
		return
	}

	code := e.SUCCESS
	isExist := models.ExistTagById(id)
	if  isExist{
		data := make(map[string]interface{})

		data["modified_by"] = modifiedBy
		data["name"] = name
		data["state"] = state
		models.EditTag(id, data)
	}

	if !isExist {
		code = e.ERROR_NOT_EXIST_TAG
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 删除文章tag
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}

	valid.Min(id, 1, "id").Message("id 必须大于0")
	if valid.HasErrors() {
		c.JSON(http.StatusOK, gin.H{
			"code": e.INVALID_PARAMS,
			"msg": e.GetMsg(e.INVALID_PARAMS),
			"data": fmt.Sprintf("%v", valid.Errors),
		})
		return
	}

	code := e.SUCCESS
	isExist := models.ExistTagById(id)
	if isExist{
		models.DeletedTag(id)
	}
	if !isExist {
		code = e.ERROR_NOT_EXIST_TAG
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": make(map[string]string),
	})
}