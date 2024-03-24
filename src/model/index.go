package model

import (
	"bs-server/src/util"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int       `gorm:"primary_key" json:"id"`
	CreatedOn  time.Time `gorm:"autoCreateTime" json:"created_on"`
	ModifiedOn time.Time `gorm:"autoUpdateTime:milli" json:"modified_on"`
	DeleteOn   time.Time `gorm:"autoUpdateTime:milli" json:"delete_on"`
}

func InitModel() {

	var err error

	dsn := mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		util.DBUser,
		util.DBPassword,
		util.DBHost,
		util.DBName))

	db, err = gorm.Open(dsn, &gorm.Config{
		//SkipDefaultTransaction: true,                                       // 跳过默认事物
		NamingStrategy: schema.NamingStrategy{SingularTable: true}, // 表名不要默认追加s
	})

	if err != nil {
		log.Printf("error")
	}

	fmt.Println("init model success")
}

func CloseDB() {

}
