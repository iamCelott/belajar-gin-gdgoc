package routes

import (
	"belajar-gin/models"
	"belajar-gin/utils"

	"github.com/gin-gonic/gin"
)

func Detail(ctx *gin.Context) {
	var data models.Book
	id := ctx.Param("id")
	if result := utils.DB().Preload("Category").Preload("File", "type = ?", "file").First(&data, "id = ?", id); result.Error != nil {
		utils.RespondingNotFound(ctx, &models.BaseResponse{})
		return
	}
	utils.Responding(ctx, &models.BaseResponse{Data: data})
}
