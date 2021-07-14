package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"project/models"
	"project/pkg/e"
	"project/pkg/util"
)

type auth struct {
	Username string `valid:"Required; Maxsize(50)"`
	Password string `valid:"Required; Maxsize(50)"`
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	code := e.SUCCESS

	valid := validation.Validation{}
	ok, _ := valid.Valid(&auth{Username: username, Password: password})

	if !ok {
		for _, err := range valid.Errors{
			log.Fatal(err.Key, err.Message)
		}
		c.JSON(http.StatusOK, gin.H{
			"code" : e.INVALID_PARAMS,
			"msg" : e.GetMsg(e.INVALID_PARAMS),
			"data" : "",
		})
	}

	data := make(map[string]string)
	isExist := models.CheckAuth(username, password)
	if isExist {
		token, err := util.GenerateToken(username, password)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		} else {
			data["token"] = token
			code = e.SUCCESS
		}
	} else {
		code = e.ERROR_AUTH_TOKEN
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}