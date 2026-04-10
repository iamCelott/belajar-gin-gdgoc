package routes

import (
	"encoding/json"

	"belajar-gin/models"
	"belajar-gin/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Create(ctx *gin.Context) {
	var data models.Book
	js := ctx.Request.FormValue("data")
	var dataMap map[string]interface{}
	// Bind the JSON body to the Client struct
	if err := json.Unmarshal([]byte(js), &dataMap); err != nil {
		ctx.Error(err)
		utils.RespondingBadRequest(ctx, &models.BaseResponse{})
		return
	}
	errors := make(map[string]interface{})
	db := utils.DB()
	if value, ok := dataMap["category_id"].(string); ok {
		var category models.Category
		if err := db.First(&category, "id = ?", value).Error; err != nil {
			errors["category_id"] = "Kategori tidak ditemukan"
		}
	}
	id := uuid.New()
	data.ID = id.String()
	if raw, ok := dataMap["name"]; ok && raw != "" {
		if value, ok := raw.(string); ok {
			data.Name = value
		} else {
			errors["name"] = "Nama harus berupa teks"
		}
	} else {
		errors["name"] = "Nama tidak boleh kosong"
	}
	if raw, ok := dataMap["category_id"]; ok && raw != "" {
		if value, ok := raw.(string); ok {
			data.CategoryID = value
		} else {
			errors["category_id"] = "Kategori ID harus berupa teks"
		}
	} else {
		errors["category_id"] = "Kategori ID tidak boleh kosong"
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
	} else {
		errors["author"] = "Author tidak boleh kosong"
	}

	if len(errors) > 0 {
		utils.RespondingBadRequest(ctx, &models.BaseResponse{Meta: map[string]interface{}{"errors": errors}})
		return
	}

	tx := utils.DB().Begin()
	if result := tx.Create(&data); result.Error != nil {
		tx.Rollback()
		ctx.Error(result.Error)
		utils.RespondingInternalError(ctx, &models.BaseResponse{})
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

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		ctx.Error(err)
		utils.RespondingInternalError(ctx, &models.BaseResponse{})
		return
	}
	utils.Responding(ctx, &models.BaseResponse{Message: "Buku berhasil ditambahkan", Data: data})
}
