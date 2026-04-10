package routes

import (
	"belajar-gin/models"
	"belajar-gin/utils"

	"github.com/gin-gonic/gin"
)

func Delete(ctx *gin.Context) {
	tx := utils.DB().Begin()
	var data models.Book
	id := ctx.Param("id")
	// Find the client by ID
	if result := tx.First(&data, "id = ?", id); result.Error != nil {
		tx.Rollback()
		utils.RespondingNotFound(ctx, &models.BaseResponse{})
		return
	}
	// Create the client in the database
	if result := tx.Delete(&data); result.Error != nil {
		tx.Rollback()
		ctx.Error(result.Error)
		utils.RespondingInternalError(ctx, &models.BaseResponse{})
		return
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		ctx.Error(err)
		utils.RespondingInternalError(ctx, &models.BaseResponse{})
		return
	}
	utils.Responding(ctx, &models.BaseResponse{Message: "Buku berhasil dihapus", Data: data})
}
