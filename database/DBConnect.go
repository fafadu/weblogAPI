package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//設定全域變數:在其他專案引用DB功能
var DBconnect *gorm.DB //附上值
var err error

func DD() {

	dsn := "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"
	DBconnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

}
