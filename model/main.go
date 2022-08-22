package model

import (
	"gin-react-template/common"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
)

var DB *gorm.DB

func createAdminAccount() {
	var user User
	DB.Where(User{Role: common.RoleAdminUser}).Attrs(User{
		Username:    "admin",
		Password:    "123456",
		Role:        common.RoleAdminUser,
		Status:      common.UserStatusEnabled,
		DisplayName: "Administrator",
	}).FirstOrCreate(&user)
}

func CountTable(tableName string) (num int) {
	DB.Table(tableName).Count(&num)
	return
}

func InitDB() (db *gorm.DB, err error) {
	if os.Getenv("SQL_DSN") != "" {
		// Use MySQL
		db, err = gorm.Open("mysql", os.Getenv("SQL_DSN"))
	} else {
		// Use SQLite
		db, err = gorm.Open("sqlite3", common.SQLitePath)
	}
	if err == nil {
		DB = db
		db.AutoMigrate(&File{})
		db.AutoMigrate(&User{})
		db.AutoMigrate(&Option{})
		createAdminAccount()
		return DB, err
	} else {
		log.Fatal(err)
	}
	return nil, err
}
