package utils

import (
	"net/http"

	"belajar-gin/models"

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
func RespondingSuccess(ctx *gin.Context, res *models.BaseResponse, claim ...string) {
	var message = "Sukses"
	if res.Message != "" {
		message = res.Message
	}
	Responding(ctx, &models.BaseResponse{Status: http.StatusOK, Message: message, Data: res.Data}, claim...)
}
func RespondingInternalError(ctx *gin.Context, res *models.BaseResponse, claim ...string) {
	var message = "Terjadi kesalahan di server"
	if res.Message != "" {
		message = res.Message
	}
	Responding(ctx, &models.BaseResponse{Status: http.StatusInternalServerError, Message: message, Data: res.Data}, claim...)
}
func RespondingMethodNotAllowed(ctx *gin.Context, res *models.BaseResponse, claim ...string) {
	var message = "Method tidak di izinkan"
	if res.Message != "" {
		message = res.Message
	}
	Responding(ctx, &models.BaseResponse{Status: http.StatusMethodNotAllowed, Message: message, Data: res.Data}, claim...)
}
func RespondingUnauthorized(ctx *gin.Context, res *models.BaseResponse, claim ...string) {
	var message = "Anda tidak memilik hak akses"
	if res.Message != "" {
		message = res.Message
	}
	Responding(ctx, &models.BaseResponse{Status: http.StatusUnauthorized, Message: message, Data: res.Data}, claim...)
}
func RespondingNotFound(ctx *gin.Context, res *models.BaseResponse, claim ...string) {
	var message = "Data Tidak Ditemukan"
	if res.Message != "" {
		message = res.Message
	}
	Responding(ctx, &models.BaseResponse{Status: http.StatusNotFound, Message: message, Data: res.Data}, claim...)
}
func RespondingUnprocceable(ctx *gin.Context, res *models.BaseResponse, claim ...string) {
	var message = "Permintaan tidak dapat dilakukan"
	if res.Message != "" {
		message = res.Message
	}
	Responding(ctx, &models.BaseResponse{Status: http.StatusUnprocessableEntity, Message: message, Data: res.Data}, claim...)
}
func RespondingBadRequest(ctx *gin.Context, res *models.BaseResponse, claim ...string) {
	var message = "Format request tidak valid"
	if res.Message != "" {
		message = res.Message
	}
	Responding(ctx, &models.BaseResponse{Status: http.StatusBadRequest, Message: message, Data: res.Data, Meta: res.Meta}, claim...)
}
func RespondingConflict(ctx *gin.Context, res *models.BaseResponse, claim ...string) {
	var message = "Data entity konflik"
	if res.Message != "" {
		message = res.Message
	}
	Responding(ctx, &models.BaseResponse{Status: http.StatusConflict, Message: message, Data: res.Data}, claim...)
}
func RespondingTokenExpired(ctx *gin.Context, res *models.BaseResponse, claim ...string) {
	var message = "Sesi anda telah berakhir, silahkan login kembali"
	if res.Message != "" {
		message = res.Message
	}
	Responding(ctx, &models.BaseResponse{Status: http.StatusExpectationFailed, Message: message, Data: res.Data}, claim...)
}
