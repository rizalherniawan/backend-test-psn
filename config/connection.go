package config

import (
	"fmt"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	val := url.Values{}
	val.Add("parseTime", "True")
	val.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?%v",
		Configuration().dbUsername, Configuration().dbPassword, Configuration().dbHost,
		Configuration().dbPort, Configuration().dbName, val.Encode())

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
