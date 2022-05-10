package service

import (
	"GolangAPI/pojo"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary FindALLUsers
// @Tags Users
// @Produce application/json
// @Router / [GET]
func FindALLUsers(c *gin.Context) {
	users := pojo.FindALLUsers() //回傳世user陣列
	c.JSON(http.StatusOK, users)
}

// @Summary FindByUserId
// @Tags Users
// @Produce application/json
// @param id path int true "id"
// @Router /{id} [GET]
func FindByUserId(c *gin.Context) {
	user := pojo.FindByUserId(c.Param("id"))
	if user.Id == 0 { //找不到
		c.JSON(http.StatusNotFound, "Error沒有此id或格式錯誤")
		return
	}
	log.Println("User -->", user)
	c.JSON(http.StatusOK, user)
}

// @Summary ADDuser
// @Tags Users
// @Produce application/json
// @param data body pojo.User true "data"
// @Router / [POST]
func PostUser(c *gin.Context) {
	user := pojo.User{}      //定義user變數是pojo的user //把post的資料傳給變數 傳進來是json格式 用指標
	err := c.BindJSON(&user) //bind json 如果沒成功 回傳是error
	if err != nil {
		c.JSON(http.StatusNotAcceptable, "Error :"+err.Error()) //如果出現錯誤 無法建立 會回傳回去 json可以直接寫訊息是什麼
		return
	}
	newUser := pojo.CreatUser(user)
	c.JSON(http.StatusOK, newUser)
}

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
	user := pojo.User{}
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
