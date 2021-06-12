package infrastructure

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewGormConnect() *gorm.DB {
	var database string
	// NOTE: railsのDBと合わせるため、loc=Asia%2FTokyoをつけない。(DB保存時はUTC、取り出す時にJTCにする)
	// https://github.com/go-sql-driver/mysql
	switch os.Getenv("environment") {
	case "production":
		database = "root:finder0501@tcp(" + os.Getenv("DB_HOST") + ")/grpc-db?charset=utf8&parseTime=true"
	default:
		database = "root:finder0501@tcp(api_db_1)/grpc_development?charset=utf8&parseTime=true"
	}

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
