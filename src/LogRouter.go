package src

import (
	"GolangAPI/service"

	"github.com/gin-gonic/gin"
)

func AddUserRouter(r *gin.RouterGroup) { //裡面要放的參數是 router的群組
	log := r.Group("/logs") //練立user選用群組 Group是user的地址
	log.GET("/", service.FindALLlogs)

	log.GET("/:id", service.FindBylogId)

	log.POST("/", service.CreatLog)

	/*		user.DELETE("/:id", service.DeleteUser)
			user.PUT("/:id", service.PutUser)
	*/
}
