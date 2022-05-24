package main

import (
	"GolangAPI/database"
	_ "GolangAPI/docs" //匯入docs init的部分
	_ "GolangAPI/settings"
	"GolangAPI/src"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           weblog API
// @version         1.0
// @description     可查詢weblog的資訊
// @license.name  Apache 2.0

//為主程式進入點。初始化gin Engine，設定Router群組路徑/demo並啟動engine。
func main() {
	router := gin.Default() //設定路由器
	// use ginSwagger middleware to serve the API docs
	var webport string = viper.GetString("web.port")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log := router.Group("log")
	src.AddUserRouter(log) //接受群組?

	go func() { //main要設資料庫,開啟專案時要同步啟動資料庫 DD的func自動連結',附值了可以使用變數database裡面的操作

		database.DD()
	}()
	router.Run(":" + webport)
}
