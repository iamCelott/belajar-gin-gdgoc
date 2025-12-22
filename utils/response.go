package utils

import (
	"belajar-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Responding(ctx *gin.Context, res *models.BaseResponse, claim ...string) {
	response := models.BaseResponse{Status: http.StatusOK, Message: "Succes", Data: nil}
	if res.Status != 0 {
		response.Status = res.Status
	}
	if res.Message != "" {
		response.Message = res.Message
	}
	if res.Data != nil {
		response.Data = res.Data
	}
	if res.Meta != nil {
		response.Meta = res.Meta
	}
	ctx.JSON(response.Status, &response)
}