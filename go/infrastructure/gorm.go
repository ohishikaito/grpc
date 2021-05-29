package infrastructure

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewGormConnect() *gorm.DB {
	database := "root:finder0501@tcp(go_db_1)/grpc_development?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open("mysql", database)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
