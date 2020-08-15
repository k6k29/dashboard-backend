package user

import "dashboard/postgresql"

type Serializer struct {
	Id         uint   `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	ActualName string `json:"actual_name"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
}

func (u *User) UserSerializer() Serializer {
	var serializer Serializer
	serializer.Id = u.BaseModel.ID
	serializer.Username = u.Username
	serializer.Password = u.Password
	db := postgresql.GetInstance()
	var profileModel Profile
	if querySet := db.Where("user_id = ?", u.BaseModel.ID).Order("id desc").First(&profileModel); querySet.Error == nil {
		serializer.ActualName = profileModel.ActualName
		serializer.Mobile = profileModel.Mobile
		serializer.Email = profileModel.Email
	}
	return serializer
}

func UserArraySerializers(userArray []User) []Serializer {
	var serializerArray []Serializer
	for k, _ := range userArray {
		serializerArray = append(serializerArray, userArray[k].UserSerializer())
	}
	return serializerArray
}
