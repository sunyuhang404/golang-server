package model

import (
	"bs-server/src/util"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func InitModel() {

	var err error

	dsn := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		util.DBUser,
		util.DBPassword,
		util.DBHost,
		util.DBName))

	db, err = gorm.Open(dsn, &gorm.Config{
		SkipDefaultTransaction: true, // 跳过默认事物
	})

	if err != nil {
		log.Printf("error")
	}

}

func CloseDB() {

}
