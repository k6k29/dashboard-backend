package dockerCloud

import (
	"dashboard/error/errorResponse"
	"dashboard/model/dockerCloud"
	"dashboard/postgresql"
	"dashboard/response"
	"dashboard/util/dockerTool"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

func GetDockerCloudList(c *gin.Context) {
	db := postgresql.GetInstance()
	dockerCloudDb := db.Table("docker_clouds").Where("deleted_at is null")
	if name := c.DefaultQuery("name", ""); name != "" {
		dockerCloudDb = dockerCloudDb.Where("name like ?", "%"+name+"%")
	}
	if host := c.DefaultQuery("host", ""); host != "" {
		dockerCloudDb = dockerCloudDb.Where("host like ?", "%"+host+"%")
	}
	var dockerCloudModelArray []dockerCloud.DockerCloud
	if page := c.DefaultQuery("page", ""); page != "" {
		pageInt, _ := strconv.Atoi(page)
		if querySet := dockerCloudDb.Limit(20).Offset((pageInt - 1) * 20).Order("id desc").Find(&dockerCloudModelArray); querySet.Error != nil {
			e := errorResponse.Response{ErrorCode: querySet.Error.Error()}
			c.JSON(http.StatusBadRequest, &e)
			panic(querySet.Error)
		}
		var pageResponse response.PageResponse
		pageResponse.Results = dockerCloud.ArraySerializers(dockerCloudModelArray)
		dockerCloudDb.Count(&pageResponse.Count)
		c.JSON(http.StatusOK, &pageResponse)
	} else {
		if querySet := dockerCloudDb.Order("id desc").Find(&dockerCloudModelArray); querySet.Error != nil {
			e := errorResponse.Response{ErrorCode: querySet.Error.Error()}
			c.JSON(http.StatusBadRequest, &e)
			panic(querySet.Error)
		}
		dockerCloudSerializer := dockerCloud.ArraySerializers(dockerCloudModelArray)
		c.JSON(http.StatusOK, &dockerCloudSerializer)
	}
}

func GetDockerCloud(c *gin.Context) {
	db := postgresql.GetInstance()
	dockerCloudID := c.Param("id")
	var dockerCloudModel dockerCloud.DockerCloud
	if querySet := db.Where("deleted_at is null").Find(&dockerCloudModel, dockerCloudID); querySet.Error != gorm.ErrRecordNotFound {
		e := errorResponse.Response{ErrorCode: "DockerCloud不存在"}
		c.JSON(http.StatusBadRequest, &e)
		return
	} else if querySet.Error != nil {
		e := errorResponse.Response{ErrorCode: querySet.Error.Error()}
		c.JSON(http.StatusBadRequest, &e)
		panic(querySet.Error)
	} else {
		serializer := dockerCloudModel.Serializer()
		c.JSON(http.StatusOK, &serializer)
	}
}

func CreateDockerCloud(c *gin.Context) {
	var serializer dockerCloud.Serializer
	if err := json.NewDecoder(c.Request.Body).Decode(&serializer); err != nil {
		e := errorResponse.Response{ErrorCode: "参数错误"}
		c.JSON(http.StatusBadRequest, &e)
		return
	}
	if err := serializer.Save(); err != nil {
		e := errorResponse.Response{ErrorCode: err.Error()}
		c.JSON(http.StatusBadRequest, &e)
		panic(err.Error())
	}
	c.JSON(http.StatusCreated, "")
}

func UpdateDockerCloud(c *gin.Context) {
	var serializer dockerCloud.Serializer
	if err := json.NewDecoder(c.Request.Body).Decode(&serializer); err != nil {
		e := errorResponse.Response{ErrorCode: "参数错误"}
		c.JSON(http.StatusBadRequest, &e)
		return
	}
	if err := serializer.Save(); err != nil {
		e := errorResponse.Response{ErrorCode: err.Error()}
		c.JSON(http.StatusBadRequest, &e)
		panic(err.Error())
	}
	c.JSON(http.StatusAccepted, "")
}

func DeleteDockerCloud(c *gin.Context) {
	db := postgresql.GetInstance()
	dockerCloudId := c.Param("id")
	var dockerCloudModel dockerCloud.DockerCloud
	if querySet := db.Where("deleted_at is null").Find(&dockerCloudModel, dockerCloudId); querySet.Error != gorm.ErrRecordNotFound {
		e := errorResponse.Response{ErrorCode: "DockerCloud不存在"}
		c.JSON(http.StatusBadRequest, &e)
		return
	} else if querySet.Error != nil {
		e := errorResponse.Response{ErrorCode: querySet.Error.Error()}
		c.JSON(http.StatusBadRequest, &e)
		panic(querySet.Error)
	} else {
		now := time.Now()
		dockerCloudModel.BaseModel.DeletedAt = now
		if querySet := db.Save(&dockerCloudModel); querySet.Error != nil {
			e := errorResponse.Response{ErrorCode: querySet.Error.Error()}
			c.JSON(http.StatusBadRequest, &e)
			panic(querySet.Error)
		}
		c.JSON(http.StatusNoContent, "")
	}
}

func TestDockerCloudConn(c *gin.Context) {
	var testConnRequest TestConnRequest
	if err := json.NewDecoder(c.Request.Body).Decode(&testConnRequest); err != nil {
		e := errorResponse.Response{ErrorCode: "参数错误"}
		c.JSON(http.StatusBadRequest, &e)
		return
	}
	_, cli, info, err := dockerTool.NewClient(testConnRequest.Host, strconv.Itoa(testConnRequest.Port), testConnRequest.UseTLS, testConnRequest.TLSCaCert, testConnRequest.TLSCert, testConnRequest.TLSKey)
	if err != nil {
		e := errorResponse.Response{ErrorCode: err.Error()}
		c.JSON(http.StatusBadRequest, &e)
		return
	}
	cli.Close()
	c.JSON(http.StatusOK, &info)
}

func ListDockerCloudContainer(c *gin.Context) {
	db := postgresql.GetInstance()
	dockerCloudID := c.DefaultQuery("docker_cloud_id","")
	var dockerCloudModel dockerCloud.DockerCloud
	if querySet := db.Where("deleted_at is null").Find(&dockerCloudModel, dockerCloudID); querySet.Error != gorm.ErrRecordNotFound {
		e := errorResponse.Response{ErrorCode: "DockerCloud不存在"}
		c.JSON(http.StatusBadRequest, &e)
		return
	} else if querySet.Error != nil {
		e := errorResponse.Response{ErrorCode: querySet.Error.Error()}
		c.JSON(http.StatusBadRequest, &e)
		panic(querySet.Error)
	} else {
		serializer := dockerCloudModel.Serializer()
		ctx, cli, _, err := dockerTool.NewClient(serializer.Host, strconv.Itoa(serializer.Port), serializer.UseTLS, serializer.TLSCaCert, serializer.TLSCert, serializer.TLSKey)
		if err != nil {
			e := errorResponse.Response{ErrorCode: err.Error()}
			c.JSON(http.StatusBadRequest, &e)
			return
		}
		defer cli.Close()
		// 搜索
		var containerListOptions types.ContainerListOptions
		args := filters.NewArgs()
		argsMap := make(map[string]string)
		if status := c.DefaultQuery("status", ""); status != "" {
			argsMap["status"] = status
		}
		if image := c.DefaultQuery("image", ""); image != "" {
			argsMap["image"] = ".*" + image + ".*"
		}
		if name := c.DefaultQuery("name", ""); name != "" {
			argsMap["name"] = ".*" + name + ".*"
		}
		args.MatchKVList("label", argsMap)
		containerListOptions.Filters = args
		containers, err := cli.ContainerList(ctx, containerListOptions)
		c.JSON(http.StatusOK, &containers)
	}
}

func ListDockerCloudImage(c *gin.Context) {
	db := postgresql.GetInstance()
	dockerCloudID := c.DefaultQuery("docker_cloud_id","")
	var dockerCloudModel dockerCloud.DockerCloud
	if querySet := db.Where("deleted_at is null").Find(&dockerCloudModel, dockerCloudID); querySet.Error != gorm.ErrRecordNotFound {
		e := errorResponse.Response{ErrorCode: "DockerCloud不存在"}
		c.JSON(http.StatusBadRequest, &e)
		return
	} else if querySet.Error != nil {
		e := errorResponse.Response{ErrorCode: querySet.Error.Error()}
		c.JSON(http.StatusBadRequest, &e)
		panic(querySet.Error)
	} else {
		serializer := dockerCloudModel.Serializer()
		ctx, cli, _, err := dockerTool.NewClient(serializer.Host, strconv.Itoa(serializer.Port), serializer.UseTLS, serializer.TLSCaCert, serializer.TLSCert, serializer.TLSKey)
		if err != nil {
			e := errorResponse.Response{ErrorCode: err.Error()}
			c.JSON(http.StatusBadRequest, &e)
			return
		}
		defer cli.Close()
		// 搜索
		var imageListOptions types.ImageListOptions
		args := filters.NewArgs()
		argsMap := make(map[string]string)
		if name := c.DefaultQuery("name", ""); name != "" {
			argsMap["name"] = ".*" + name + ".*"
		}
		args.MatchKVList("label", argsMap)
		imageListOptions.Filters = args
		images, err := cli.ImageList(ctx, imageListOptions)
		c.JSON(http.StatusOK, &images)
	}
}

