package router

import (
	"dashboard/api/login"
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
	return r
}
