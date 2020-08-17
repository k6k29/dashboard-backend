package user

import (
	"dashboard/postgresql"
	"dashboard/util/password"
	"errors"
)

type Serializer struct {
	Id         uint   `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	ProfileId  uint   `json:"profile_id"`
	ActualName string `json:"actual_name"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
}

func (u *User) Serializer() Serializer {
	var serializer Serializer
	serializer.Id = u.BaseModel.ID
	serializer.Username = u.Username
	serializer.Password = u.Password
	db := postgresql.GetInstance()
	var profileModel Profile
	if querySet := db.Where("user_id = ?", u.BaseModel.ID).Order("id desc").First(&profileModel); querySet.Error == nil {
		serializer.ProfileId = profileModel.BaseModel.ID
		serializer.ActualName = profileModel.ActualName
		serializer.Mobile = profileModel.Mobile
		serializer.Email = profileModel.Email
	}
	return serializer
}

func ArraySerializers(userArray []User) []Serializer {
	var serializerArray []Serializer
	for k, _ := range userArray {
		serializerArray = append(serializerArray, userArray[k].Serializer())
	}
	return serializerArray
}

func (s *Serializer) Save() error {
	db := postgresql.GetInstance()
	var userModel User
	if s.Id == 0 {
		if querySet := db.Find(&userModel, s.Id); querySet.Error != nil {
			return querySet.Error
		}
	}
	var userCount int
	db.Table("users").Where("username = ? AND id != ?", s.Username, userModel.BaseModel.ID).Count(&userCount)
	if userCount > 0 {
		return errors.New("用户名已存在")
	}
	userModel.Username = s.Username
	if userModel.Password != s.Password {
		encryptPassword, err := password.EncryptPassword(s.Password)
		if err != nil {
			return err
		}
		userModel.Password = encryptPassword
	}
	if s.Id == 0 {
		if querySet := db.Create(&userModel); querySet.Error != nil {
			return querySet.Error
		}
	} else {
		if querySet := db.Save(&userModel); querySet.Error != nil {
			return querySet.Error
		}
	}
	var profileModel Profile
	if s.ProfileId != 0 {
		if querySet := db.Find(&profileModel, s.ProfileId); querySet.Error != nil {
			return querySet.Error
		}
	}
	profileModel.UserId = userModel.BaseModel.ID
	profileModel.ActualName = s.ActualName
	profileModel.Email = s.Email
	profileModel.Mobile = s.Mobile
	if s.ProfileId == 0 {
		if querySet := db.Create(&profileModel); querySet.Error != nil {
			return querySet.Error
		}
	} else {
		if querySet := db.Save(&profileModel); querySet.Error != nil {
			return querySet.Error
		}
	}
	return nil
}
