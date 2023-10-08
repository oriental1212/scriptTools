package sqlTools

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
 * 连接配置
 */
var username = "db_admin"
var password = "Rk2z6mFg9NUjv8h7"
var url = "mysql.dev.myyixue.com:3306"
var database = "sc_irs_center"

func ConnectDevMySQL() (db *gorm.DB) {
	dsn := username + ":" + password + "@tcp(" + url + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
