package service

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"instant_messaging/models"
	"instant_messaging/utils"
	"strconv"
)

// GetUserList 获取用户列表
// @Tags 用户模块
// @Summary 用户列表
// @Success 200 {object} map[string]interface{}
// @Router /users [get]
func GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(200, gin.H{"code": 1, "data": data})
}

// CreateUser 创建用户
// @Tags 用户模块
// @Summary 新增用户
// @Accept json
// @Produce json
// @Param user body models.UserBasic true "用户信息"
// @Success 200 {string} string "新增用户成功！"
// @Router /user [post]
func CreateUser(c *gin.Context) {
	var user models.UserBasic
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "参数解析失败"})
		return
	}
	sameUser, count := models.FindUserByName(user.Name)
	if count > 0 {
		fmt.Println(sameUser.ID)
		c.JSON(200, gin.H{"code": 0, "message": "用户名重复"})
		return
	}

	user.Password = utils.HashPassword(user.Password)
	models.CreateUser(user)
	c.JSON(200, gin.H{"code": 1, "message": "新增用户成功！"})
}

// DeleteUser 删除用户
// @Tags 用户模块
// @Summary 删除用户
// @Param id path int true "用户ID"
// @Success 200 {string} string "删除用户成功！"
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{"error": "ID 不可用"})
		return
	}

	if err := models.DeleteUserByID(uint(id)); err != nil {
		c.JSON(200, gin.H{"code": 0, "message": "删除用户失败"})
		return
	}

	c.JSON(200, gin.H{"code": 1, "message": "删除用户成功！"})
}

// UpdateUser 更新用户
// @Tags 用户模块
// @Summary 更新用户信息
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param user body models.UserBasic true "用户信息"
// @Success 200 {string} string "修改成功！"
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(400, gin.H{"error": "ID 不可用"})
		return
	}

	var user models.UserBasic
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "参数解析失败"})
		return
	}

	user.ID = uint(id)
	_, err = govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{"code": 0, "message": err.Error()})
		return
	}
	models.UpdateUser(user)
	c.JSON(200, gin.H{"code": 1, "message": "修改成功！"})
}

// Login 用户登陆
// @Tags 用户模块
// @Summary 登陆
// @Param user body models.UserBasic true "用户信息"
// @Success 200 {string} string "登陆成功"
// @Router /login [post]
func Login(c *gin.Context) {
	var user models.UserBasic
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "参数解析失败"})
		return
	}
	data, row := models.FindUserByName(user.Name)
	if row <= 0 {
		c.JSON(200, gin.H{"code": 0, "message": "用户名或密码错误"})
		return
	}
	success := utils.VerifyPassword(user.Password, data.Password)
	token, err := utils.GenerateJWT(int(data.ID))
	if err != nil {
		c.JSON(200, gin.H{"code": 0, "message": "生成token失败"})
		return
	}
	if !success {
		c.JSON(200, gin.H{"code": 0, "message": "登陆失败"})
		return
	}
	rs := map[string]string{
		"accessToken": token,
	}
	c.JSON(200, gin.H{"code": 1, "message": "登陆成功", "data": rs})
}
