package router

import (
	"dashboard/api/dockerCloud"
	"dashboard/api/login"
	"dashboard/api/user"
	"dashboard/middleware/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.CorsHandler())
	r.Use(gin.Recovery())
	api := r.Group("/api")
	login.LoadLoginRouter(api)
	user.LoadUserRouter(api)
	dockerCloud.LoadDockerCloudRouter(api)
	return r
}
