package main

import (
	router "server/pkg/api"
)

func main() {
	router := router.InitRouter()
	router.Run("0.0.0.0:5000")
}
