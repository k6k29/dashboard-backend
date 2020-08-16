package main

import (
	"dashboard/model/user"
	"dashboard/postgresql"
)

func Migrate() {
	db := postgresql.GetInstance()
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&user.Profile{})
}
