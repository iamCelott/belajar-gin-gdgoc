package routes

import (
	"strconv"

	"belajar-gin/models"
	"belajar-gin/utils"

	"github.com/gin-gonic/gin"
)

func Read(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	search := ctx.DefaultQuery("search", "")
	db := utils.DB().Preload("File", "type = ?", "file").Where("name LIKE ?", "%"+search+"%").Order("name ASC")
	data, pagination, err := utils.Paginate[models.Book](db, &models.Book{}, page, pageSize)
	if err != nil {
		ctx.Error(err)
		utils.RespondingInternalError(ctx, &models.BaseResponse{})
		return
	}
	utils.Responding(ctx, &models.BaseResponse{
		Data: data,
		Meta: map[string]interface{}{"pagination": pagination},
	})
}
