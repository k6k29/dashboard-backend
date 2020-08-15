package user

import "dashboard/model"

type User struct {
	model.BaseModel
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null" json:"password"`
}

type Profile struct {
	model.BaseModel
	UserId     uint   `json:"user_id"`
	ActualName string `json:"actual_name"`
	Mobile     string `json:"mobile"`
	Email      string `json:"email"`
}
