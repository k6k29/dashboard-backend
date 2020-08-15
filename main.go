package main

import (
	"dashboard/router"
)

func main() {
	r := router.SetupRouter()
	r.Run("0.0.0.0:8080")
}
