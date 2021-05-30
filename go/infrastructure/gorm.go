package infrastructure

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewGormConnect() *gorm.DB {
	database := "root:finder0501@tcp(go_db_1)/grpc_development?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open("mysql", database)
	if err != nil {
		log.Fatal(err.Error())
	}
	file, err := os.OpenFile("sql_log_file", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	// logにファイルを出力するようにする
	log.SetOutput(file)
	db.LogMode(true)
	db.SetLogger(log.New(file, "", 0))
	return db
}
