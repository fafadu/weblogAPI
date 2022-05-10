package src

import (
	"GolangAPI/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) { //裡面要放的參數是 router的群組
	user := r.Group("/users") //練立user選用群組 Group是user的地址
	user.GET("/", service.FindALLUsers)
	user.GET("/:id", service.FindByUserId)
	user.POST("/", service.PostUser)
	user.DELETE("/:id", service.DeleteUser)
	user.PUT("/:id", service.PutUser)
}
