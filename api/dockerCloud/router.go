package dockerCloud

import (
	"dashboard/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func LoadDockerCloudRouter(api *gin.RouterGroup) {
	dockerCloudApi := api.Group("docker-cloud")
	dockerCloudApi.Use(jwt.NeedJwtAuth())
	dockerCloudApi.GET("", GetDockerCloudList)
	dockerCloudApi.GET(":id", GetDockerCloud)
	dockerCloudApi.POST("", CreateDockerCloud)
	dockerCloudApi.PATCH("", UpdateDockerCloud)
	dockerCloudApi.DELETE(":id", DeleteDockerCloud)
}
