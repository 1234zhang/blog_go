package jwt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	"project/pkg/e"
	"project/pkg/util"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := e.SUCCESS
		var data interface{}

		token := c.Query("token")
		if token == "" {
			code = e.INVALID_PARAMS
		}

		if code == e.SUCCESS {
			claims, err := util.ParseToken(token)
			if err == nil && claims.ExpiresAt < time.Now().Unix() {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			}
			if err != nil {
				code = e.ERROR_AUTH_TOKEN_TOKEN_TIMEOUT
				fmt.Sprintf("parse jwt happen some err : %v", err.Error())
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusOK, gin.H{
				"code" : code,
				"msg" : e.GetMsg(code),
				"data" : data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}