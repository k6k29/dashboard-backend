package user

import (
	"dashboard/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func LoadUserRouter(api *gin.RouterGroup) {
	userApi := api.Group("user")
	userApi.Use(jwt.NeedJwtAuth())
	userApi.GET("", GetUserList)
	userApi.GET(":id", GetUser)
	userApi.POST("", CreateUser)
	userApi.PATCH("", UpdateUser)
	userApi.DELETE(":id", DeleteUser)
}
