package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"mxshop_srvs/user_srv/model"
)

func main() {
	dsn := "root:123456@tcp(192.168.202.170:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 为了使生成的表不带S
		},
	})

	_ = db.AutoMigrate(&model.User{})

}
