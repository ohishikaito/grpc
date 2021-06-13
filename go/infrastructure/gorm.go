package infrastructure

import (
	"io"
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

	// ログファイルを開く。CREATE=作成 WRONLYで読み書き APPENDで後ろに追加
	logfile, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err.Error())
	}
	// SetOutPutで出力先を指定 MultiWriterで2つの出力先を指定できる
	log.SetOutput(io.MultiWriter(logfile, os.Stdout))
	// ログを吐いた時間を出す
	log.SetFlags(log.Ldate | log.Ltime)
	// LogMode true でログを吐き出すように
	db.LogMode(true)
	// 標準出力で出力するように。SetLogger先は一つしか選べないっぽい
	db.SetLogger(log.New(os.Stdout, "", 0))
	return db
}
