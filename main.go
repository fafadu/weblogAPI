package main

import (
	"GolangAPI/database"
	_ "GolangAPI/docs" //匯入docs init的部分

	"GolangAPI/src"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           FAFA的SwaggerAPI
// @version         1.0
// @description     FAFA的練習
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8000
// @BasePath  /v1/usersnn

//為主程式進入點。初始化gin Engine，設定Router群組路徑/demo並啟動engine。
func main() {
	router := gin.Default() //設定路由器
	// use ginSwagger middleware to serve the API docs

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/v1")
	src.AddUserRouter(v1) //接受群組?

	go func() { //main要設資料庫,開啟專案時要同步啟動資料庫 DD的func自動連結',附值了可以使用變數database裡面的操作

		database.DD()
	}()

	router.Run(":8000")
}
