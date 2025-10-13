package db

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB

func init() {
	data, e := os.ReadFile("./dbConfig.txt")
	if e != nil {
		panic(e)
	}
	dsn := string(data)
	var err error
	Database, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
}
