package utils

import (
	"errors"
	"fmt"
	"mime/multipart"

	"belajar-gin/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
func SaveFile(file *multipart.FileHeader, ctx *gin.Context, tx *gorm.DB, ownerType string, ownerId string, filetype string, isMultiple bool) error {
	if !isMultiple {
		var existedFiles []models.File
		if result := DB().Find(&existedFiles, "owner_id = ? AND owner_type = ? AND type = ?", ownerId, ownerType, filetype); result.Error == nil {
			for _, v := range existedFiles {
				if r := DB().Delete(v); r.Error == nil {
					err := DeleteFile(v.Path)
					if err != nil {
						ctx.Error(err)
					}
				}
			}
		} else {
			ctx.Error(result.Error)
		}

	}

	// contentType := file.Header.Get("Content-Type")
	allowedExtension := []string{"png", "jpeg", "jpg"}
	ext := GetExtension(file.Filename)
	if !Contains(allowedExtension, ext) {
		return errors.New("tipe file tidak di perbolehkan")
	}
	// var fileType string
	// if strings.HasPrefix(contentType, "image/") {
	// 	fileType = "Image"
	// } else if strings.HasPrefix(contentType, "video/") {
	// 	fileType = "Video"
	// } else if contentType == "application/pdf" {
	// 	fileType = "Pdf"
	// } else {
	// 	fileType = "Other"
	// }
	filePath := fmt.Sprintf("./data/%v/%v/%v/%v", ownerType, ownerId, filetype, file.Filename)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		return errors.New("gagal menyimpan file")
	}
	var fileData models.File
	id := uuid.New()
	fileData.ID = id.String()
	fileData.Name = file.Filename
	fileData.Extension = ext
	fileData.Path = filePath
	fileData.Type = filetype
	fileData.OwnerType = ownerType
	fileData.OwnerID = ownerId
	if result := tx.Create(&fileData); result.Error != nil {
		return errors.New("gagal menyimpan file")
	}
	return nil
}
