package user

import "dashboard/model"


type User struct {
	model.BaseModel
	Username string `json:"username"`
	Password string `json:"password"`
}

type Profile struct {
	model.BaseModel
	UserId uint `json:"user_id"`
	ActualName string `json:"actual_name"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
}
