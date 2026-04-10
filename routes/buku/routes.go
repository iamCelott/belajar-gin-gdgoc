package buku

import (
	"belajar-gin/routes/buku/routes"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	buku := r.Group("/buku")
	buku.GET("/", routes.Read)
	buku.GET("/:id", routes.Detail)
	buku.POST("/", routes.Create)
	buku.PUT("/:id", routes.Update)
	buku.DELETE("/:id", routes.Delete)
}
