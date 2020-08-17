package user

import (
	"dashboard/error/errorResponse"
	"dashboard/model/user"
	"dashboard/postgresql"
	"dashboard/response"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetUserList(c *gin.Context) {
	db := postgresql.GetInstance()
	userDb := db.Table("users").Where("deleted_at is null")
	if username := c.DefaultQuery("username", ""); username != "" {
		userDb = userDb.Where("username like ?", "%"+username+"%")
	}
	if actualName := c.DefaultQuery("actual_name", ""); actualName != "" {
		var userIdArray []uint
		db.Table("profiles").Where("actual_name like ?", "%"+actualName+"%").Pluck("id", &userIdArray)
		userDb = userDb.Where("id in (?)", userIdArray)
	}
	var userModelArray []user.User
	if page := c.DefaultQuery("page", ""); page != "" {
		pageInt, _ := strconv.Atoi(page)
		if querySet := userDb.Limit(20).Offset((pageInt - 1) * 20).Order("id desc").Find(&userModelArray); querySet.Error != nil {
			e := errorResponse.Response{ErrorCode: querySet.Error.Error()}
			c.JSON(http.StatusBadRequest, &e)
			panic(querySet.Error)
		}
		var pageResponse response.PageResponse
		pageResponse.Results = user.ArraySerializers(userModelArray)
		userDb.Count(&pageResponse.Count)
		c.JSON(http.StatusOK, &pageResponse)
	} else {
		if querySet := userDb.Order("id desc").Find(&userModelArray); querySet.Error != nil {
			e := errorResponse.Response{ErrorCode: querySet.Error.Error()}
			c.JSON(http.StatusBadRequest, &e)
			panic(querySet.Error)
		}
		userSerializer := user.ArraySerializers(userModelArray)
		c.JSON(http.StatusOK, &userSerializer)
	}
}

func GetUser(c *gin.Context) {
	db := postgresql.GetInstance()
	userID := c.Param("id")
	var userModel user.User
	if querySet := db.Where("deleted_at is null").Find(&userModel, userID); querySet.Error != gorm.ErrRecordNotFound {
		e := errorResponse.Response{ErrorCode: "用户不存在"}
		c.JSON(http.StatusBadRequest, &e)
		return
	} else if querySet.Error != nil {
		e := errorResponse.Response{ErrorCode: querySet.Error.Error()}
		c.JSON(http.StatusBadRequest, &e)
		panic(querySet.Error)
	} else {
		serializer := userModel.Serializer()
		c.JSON(http.StatusOK, &serializer)
	}
}

func CreateUser(c *gin.Context) {
	var serializer user.Serializer
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

func UpdateUser(c *gin.Context) {
	var serializer user.Serializer
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

func DeleteUser(c *gin.Context) {
	db := postgresql.GetInstance()
	userID := c.Param("id")
	var userModel user.User
	if querySet := db.Where("deleted_at is null").Find(&userModel, userID); querySet.Error != gorm.ErrRecordNotFound {
		e := errorResponse.Response{ErrorCode: "用户不存在"}
		c.JSON(http.StatusBadRequest, &e)
		return
	} else if querySet.Error != nil {
		e := errorResponse.Response{ErrorCode: querySet.Error.Error()}
		c.JSON(http.StatusBadRequest, &e)
		panic(querySet.Error)
	} else {
		now := time.Now()
		userModel.BaseModel.DeletedAt = now
		if querySet := db.Save(&userModel); querySet.Error != nil {
			e := errorResponse.Response{ErrorCode: querySet.Error.Error()}
			c.JSON(http.StatusBadRequest, &e)
			panic(querySet.Error)
		}
		c.JSON(http.StatusNoContent, "")
	}
}
