package service

import (
	"ldchat/models"

	"github.com/gin-gonic/gin"
)

func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func GetUserList(c *gin.Context) {
	data, err := models.GetUserList()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  data,
		})
	}
}
