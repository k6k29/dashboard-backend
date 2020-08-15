package login

import "github.com/gin-gonic/gin"

func LoadLoginRouter(api *gin.RouterGroup) {
	api.POST("login", Login)
}