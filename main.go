package main

import (
	"belajar-gin/routes"
	"belajar-gin/utils"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(".env file not found")
	}
	utils.AutoMigrate()
	r := gin.Default()
	routes.Routes(r)
	port := os.Getenv("PORT")
	if port == "" {
		panic("Port is not specify")
	}
	r.Run(fmt.Sprintf(":%v", port))
}
