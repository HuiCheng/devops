package include

import (
	"fmt"

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
func InitDB() {
	var err error
	var db *gorm.DB
	if gormType == "sqlite" {
		db, err = gorm.Open("sqlite3", gormHost)
		if err != nil {
			panic("failed to connect database: " + err.Error())
		}
	}
	fmt.Println("InitDB")
	db.AutoMigrate(&configuration.Namespace{})
	DB = db
}
