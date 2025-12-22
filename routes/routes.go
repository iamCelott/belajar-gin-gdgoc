package routes

import (
	"belajar-gin/routes/buku"

	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	buku.Routes(r)
}
