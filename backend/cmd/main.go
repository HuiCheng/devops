package main

import (
	"flag"

	"github.com/HuiCheng/devops/backend/include"
	"github.com/HuiCheng/devops/backend/router"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	bind     string
	gormType string
	gormHost string
)

func initFlag() {
	flag.StringVar(&bind, "bind", ":8081", "For example, 127.0.0.1:8081 0.0.0.0:8081")
	flag.StringVar(&gormType, "gormType", "sqlite", "For example, sqlite | mysql")
	flag.StringVar(&gormHost, "gormHost", "./db.sqlite", "For example, ./db.sqlite | <username>:<password>@tcp(<host>:<port>)/<dbname>?charset=utf8")
	flag.Parse()
}

func initDB() {
	if gormType == "sqlite" {
		db, err := gorm.Open("sqlite3", gormHost)
		if err != nil {
			panic("failed to connect database: " + err.Error())
		}
		include.DB = db
	}
}

func main() {
	initFlag()
	initDB()

	router.MKRouter().Run(bind)
}
