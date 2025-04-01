package models

import (
	"gorm.io/gorm"
	"instant_messaging/utils"
	"time"
)

// models/users.go

// UserBasic 用户信息
// @Description 用户基本信息结构体
type UserBasic struct {
	gorm.Model
	Name          string // 用户名
	Password      string // 密码
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     *time.Time
	HeartbeatTime *time.Time
	LoginOutTime  *time.Time
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	//for _, v := range data {
	//	fmt.Println(v)
	//}
	return data
}

func CreateUser(user UserBasic) *gorm.DB {

	return utils.DB.Create(&user)
}

func DeleteUserByID(id uint) error {
	result := utils.DB.Delete(&UserBasic{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Updates(&user)
}

func FindUserByName(name string) (UserBasic, int64) {
	var user UserBasic
	result := utils.DB.Where("name = ?", name).First(&user)
	return user, result.RowsAffected
}

func FindUserByNameAndPwd(name, password string) (UserBasic, int64) {
	var user UserBasic
	result := utils.DB.Where("name = ? AND password = ?", name, password).First(&user)
	return user, result.RowsAffected
}

func FindUserByPhone(phone string) (UserBasic, int64) {
	user := UserBasic{}
	result := utils.DB.Where("phone = ?", phone).First(&user)
	return user, result.RowsAffected
}

func FindUserByEmail(email string) (UserBasic, int64) {
	user := UserBasic{}
	result := utils.DB.Where("email = ?", email).First(&user)
	return user, result.RowsAffected
}
