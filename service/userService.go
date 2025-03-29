package service

import (
	"github.com/gin-gonic/gin"
	"instant_messaging/models"
)

// GetUserList
// @Tags 用户列表
// @Success 200 {string} json{"code", "data"}
// @Router /user/info [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(200, gin.H{"data": data})
}
