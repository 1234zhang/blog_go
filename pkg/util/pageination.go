package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"log"
	"project/pkg/setting"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, err := com.StrTo(c.Query("page")).Int()

	if err != nil {
		log.Fatalf("get page happend some error: %v", err)
	}

	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return result
}
