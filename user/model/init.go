package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func  Database(connString string)  {
	db,err := gorm.Open(mysql.Open(connString),&gorm.Config{})
	if err != nil{
		panic(err)
	}
	DB = db
	migration()

}