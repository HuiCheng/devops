package include

import (
	"github.com/HuiCheng/devops/backend/api/v1/configuration"
	"github.com/jinzhu/gorm"

	// sqlite3
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Var
var (
	DB *gorm.DB
)

// InitDB Func
func InitDB(t string) {
	var err error
	var db *gorm.DB
	if gormType == "sqlite" {
		db, err = gorm.Open("sqlite3", gormHost)
		if err != nil {
			panic("failed to connect database: " + err.Error())
		}
	}
	if t == "test" {
		db.LogMode(true)
	}
	db.AutoMigrate(&configuration.Key{})
	db.AutoMigrate(&configuration.Namespace{})
	db.AutoMigrate(&configuration.Value{})
	db.AutoMigrate(&configuration.Config{})
	DB = db
}
