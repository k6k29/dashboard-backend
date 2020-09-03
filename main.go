package main

import (
	"dashboard/router"
	"dashboard/util/password"
	"log"
)

func main() {
	x,_ := password.EncryptPassword("123456")
	log.Println(x)
	Migrate()
	r := router.SetupRouter()
	r.Run("0.0.0.0:8080")
}
