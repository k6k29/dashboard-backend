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
	testDockerCloudConnApi := api.Group("docker-cloud-test-conn")
	testDockerCloudConnApi.Use(jwt.NeedJwtAuth())
	testDockerCloudConnApi.POST("", TestDockerCloudConn)
	DockerCloudContainerApi := api.Group("docker-cloud-container")
	DockerCloudContainerApi.Use(jwt.NeedJwtAuth())
	{
		DockerCloudContainerApi.GET("", ListDockerCloudContainer)
	}
	DockerCloudImageApi := api.Group("list-image-cloud")
	DockerCloudImageApi.Use(jwt.NeedJwtAuth())
	{
		DockerCloudImageApi.GET("", ListDockerCloudImage)
	}
}
