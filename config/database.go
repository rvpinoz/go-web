package config

import (
	"go-web/tools"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"time"
)

var DB *gorm.DB
var err error

func Mysql() {
	start := time.Now()

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	dsn := user + ":" + pass + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		go tools.Panic(err.Error(), "Connection to Database Mysql Error", start)
	}

	go tools.Info("Connected", "Connection to Database Mysql Success", start)
}
