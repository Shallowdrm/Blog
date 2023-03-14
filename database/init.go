package database

import (
	"blog/model"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	RS *dbserver
)

type dbserver struct {
	DB *gorm.DB
}

func Init() {
	dsn := "root:13545193002@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB = db
	RS = &dbserver{
		DB,
	}

	if err != nil {
		panic("连接数据库失败")
	} else {
		log.Println("连接数据库成功")
	}

}

func (RS *dbserver) CheckByUsername(user *model.RegisterInfo) bool {
	var user1 model.RegisterInfo
	err := RS.DB.Where("username = ?", user.Username).First(&user1).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func (RS *dbserver) CreatNewUser(user *model.RegisterInfo) error {
	return RS.DB.Create(&user).Error
}

func (RS *dbserver) FindUserById(username string) (user *model.RegisterInfo, boo bool) {
	err := RS.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, false
	} else {
		return user, true
	}
}
