package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"instant_messaging/models"
	"log"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/im?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// 迁移 schema
	//db.AutoMigrate(&models.UserBasic{})

	// 创建一个新的用户实例
	user := &models.UserBasic{}
	user.Name = "test3"
	user.Password = "123456"
	db.Create(user)

	var foundUser models.UserBasic
	fmt.Println(db.First(&foundUser, 1)) // 根据主键查找
	fmt.Println(foundUser)

	//db.Model(&user).Update("Password", "123456")

}
