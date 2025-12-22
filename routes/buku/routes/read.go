package routes

import (
	"belajar-gin/models"
	"belajar-gin/utils"

	"github.com/gin-gonic/gin"
)

func Read(ctx *gin.Context) {
	var data []models.Book
	result := utils.DB().Find(&data)
	if result.Error != nil {
		utils.Responding(ctx, &models.BaseResponse{
			Message: "Error",
		})
		return
	}
	utils.Responding(ctx, &models.BaseResponse{Data: data})
}
