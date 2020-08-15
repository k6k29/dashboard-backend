package user

import (
	"dashboard/error/errorResponse"
	"dashboard/model/user"
	"dashboard/postgresql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUserList(c *gin.Context) {
	db := postgresql.GetInstance()
	userDb := db.Table("users")
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
		var pageResponse PageResponse
		pageResponse.Results = user.UserArraySerializers(userModelArray)
		userDb.Count(&pageResponse.Count)
		c.JSON(http.StatusOK, &pageResponse)
	} else {
		if querySet := userDb.Order("id desc").Find(&userModelArray); querySet.Error != nil {
			e := errorResponse.Response{ErrorCode: querySet.Error.Error()}
			c.JSON(http.StatusBadRequest, &e)
			panic(querySet.Error)
		}
		userSerializer := user.UserArraySerializers(userModelArray)
		c.JSON(http.StatusOK, &userSerializer)
	}
}

func GetUser(c *gin.Context){

}