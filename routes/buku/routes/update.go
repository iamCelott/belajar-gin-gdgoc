package routes

import (
	"encoding/json"

	"belajar-gin/models"
	"belajar-gin/utils"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	var data models.Book
	var dataMap map[string]interface{}
	js := ctx.Request.FormValue("data")
	if err := json.Unmarshal([]byte(js), &dataMap); err != nil {
		ctx.Error(err)
		utils.RespondingBadRequest(ctx, &models.BaseResponse{})
		return
	}
	errors := make(map[string]interface{})
	id := ctx.Param("id")
	tx := utils.DB().Begin()
	if result := tx.First(&data, "id = ?", id); result.Error != nil {
		tx.Rollback()
		utils.RespondingNotFound(ctx, &models.BaseResponse{})
		return
	}
	db := *utils.DB()
	if value, ok := dataMap["category_id"].(string); ok {
		var category models.Category
		if err := db.First(&category, "id = ?", value).Error; err != nil {
			errors["category_id"] = "Kategori tidak ditemukan"
		}
	}
	if raw, ok := dataMap["name"]; ok && raw != "" {
		if value, ok := raw.(string); ok {
			data.Name = value
		} else {
			errors["name"] = "Nama harus berupa teks"
		}
	}
	if raw, ok := dataMap["category_id"]; ok && raw != "" {
		if value, ok := raw.(string); ok {
			data.CategoryID = value
		} else {
			errors["category_id"] = "Kategori ID harus berupa teks"
		}
	}
	if raw, ok := dataMap["description"]; ok && raw != "" {
		if value, ok := raw.(string); ok {
			data.Description = &value
		} else {
			errors["description"] = "Deskripsi harus berupa teks"
		}
	}
	if raw, ok := dataMap["author"]; ok && raw != "" {
		if value, ok := raw.(string); ok {
			data.Author = value
		} else {
			errors["author"] = "Author harus berupa teks"
		}
	}

	if len(errors) > 0 {
		utils.RespondingBadRequest(ctx, &models.BaseResponse{Meta: map[string]interface{}{"errors": errors}})
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.Error(err)
		utils.RespondingBadRequest(ctx, &models.BaseResponse{})
		return
	}

	file := form.File["file"]
	if len(file) > 0 {
		if err := utils.SaveFile(file[0], ctx, tx, "buku", data.ID, "file", false); err != nil {
			tx.Rollback()
			utils.RespondingUnprocceable(ctx, &models.BaseResponse{Message: err.Error()})
			return
		}
	}
	if result := tx.Save(&data); result.Error != nil {
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
	utils.Responding(ctx, &models.BaseResponse{Message: "Buku berhasil diperbarui", Data: data})
}
