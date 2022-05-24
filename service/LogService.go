package service

import (
	"GolangAPI/pojo"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 查詢所有weblog資訊
// @Tags Get weblog.查詢
// @Produce application/json
// @Router /log/logs [GET]
func FindALLlogs(c *gin.Context) {
	logs := pojo.FindALLlogs() //回傳世user陣列
	c.JSON(http.StatusOK, logs)
}

// @Summary 以Id搜尋welog資訊
// @Tags Get weblog.查詢
// @Produce application/json
// @param id path int true "id"
// @Router /log/logs/{id} [GET]
func FindBylogId(c *gin.Context) {
	log := pojo.FindBylogId(c.Param("id"))
	if log.Id == 0 {
		c.JSON(http.StatusNotFound, "Error沒有此id或格式錯誤")
		return
	}
	fmt.Println("log -->", log)
	c.JSON(http.StatusOK, log)
}

// @Summary ADD新增
// @Tags 新增
// @Produce application/json
// @param data body pojo.Log true "data"
// @Router /log/logs [POST]
func CreatLog(c *gin.Context) {
	log := pojo.Log{}
	err := c.BindJSON(&log)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error :"+err.Error())
		return
	}
	newLog := pojo.CreatLog(log)
	c.JSON(http.StatusOK, newLog)
}

/*

// @Summary Delete
// @Tags Users
// @Produce application/json
// @param id path int true "id"
// @Router /{id} [DELETE]
func DeleteUser(c *gin.Context) {
	user := pojo.DeleteUser(c.Param("id"))
	if !user { //這個東西不存在
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, "Successfully")
}

// @Summary UPDATE
// @Tags Users
// @Produce application/json
// @param id path int true "id"
// @param data body pojo.User true "data"
// @Router /{id} [PUT]

func PutUser(c *gin.Context) {
	user := pojo.logId{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Error")
		return
	} //要傳入要修改的id資料
	user = pojo.UpdateUser(c.Param("id"), user)
	if user.Id == 0 {
		c.JSON(http.StatusNotFound, "Error")
		return
	}
	c.JSON(http.StatusOK, user)
}
*/
