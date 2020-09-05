package login

import (
	"dashboard/error/errorResponse"
	"dashboard/middleware/jwt"
	"dashboard/model/user"
	"dashboard/postgresql"
	"dashboard/util/password"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	var request Request
	if err := json.NewDecoder(c.Request.Body).Decode(&request); err != nil {
		e := errorResponse.Response{ErrorCode: "参数错误"}
		c.JSON(http.StatusBadRequest, &e)
		return
	}
	db := postgresql.GetInstance()
	var userModel user.User
	if querySet := db.Table("users").Where("deleted_at is null AND username = ?", request.Username).Order("id desc").First(&userModel); querySet.Error == gorm.ErrRecordNotFound {
		e := errorResponse.Response{ErrorCode: "用户名错误或用户不存在"}
		c.JSON(http.StatusBadRequest, &e)
		return
	} else if querySet.Error != nil {
		e := errorResponse.Response{ErrorCode: querySet.Error.Error()}
		c.JSON(http.StatusBadRequest, &e)
		panic(querySet.Error.Error())
	} else {
		encryptPassword, err := password.EncryptPassword(request.Password)
		if err != nil {
			e := errorResponse.Response{ErrorCode: err.Error()}
			c.JSON(http.StatusBadRequest, &e)
			panic(err.Error())
		}
		if encryptPassword == userModel.Password {
			serializer := userModel.Serializer()
			var response = Response{
				Token:      jwt.GenerateToken(c, serializer.Id),
				Id:         serializer.Id,
				Username:   serializer.Username,
				ActualName: serializer.ActualName,
				Mobile:     serializer.Mobile,
				Email:      serializer.Email,
			}
			c.JSON(http.StatusOK, &response)
		} else {
			e := errorResponse.Response{ErrorCode: "账号或密码错误"}
			c.JSON(http.StatusBadRequest, &e)
			return
		}
	}
}

