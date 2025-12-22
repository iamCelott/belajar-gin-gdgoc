package main

import (
	"belajar-gin/routes"
	"belajar-gin/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(".env file not found")
	}
	r := gin.Default()
	routes.Routes(r)
	utils.AutoMigrate()
	r.Run()
}
