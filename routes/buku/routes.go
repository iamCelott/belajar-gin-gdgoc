package buku

import (
	"belajar-gin/routes/buku/routes"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	buku := r.Group("/buku")
	buku.GET("/", routes.Read)
	buku.POST("/")
	buku.PUT("/:id")
	buku.DELETE("/:id")
}
