package utils

import (
	"math"

	"belajar-gin/models"
	"gorm.io/gorm"
)

func Paginate[T any](db *gorm.DB, model interface{}, page, pageSize int) ([]T, models.Pagination, error) {
	var data []T
	var totalRecords int64
	if err := db.Model(model).Count(&totalRecords).Error; err != nil {
		return nil, models.Pagination{}, err
	}
	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))
	offset := (page - 1) * pageSize
	if pageSize < 1 {
		if err := db.Model(model).Find(&data).Error; err != nil {
			return nil, models.Pagination{}, err
		}
	} else {
		if err := db.Model(model).Limit(pageSize).Offset(offset).Find(&data).Error; err != nil {
			return nil, models.Pagination{}, err
		}
	}

	meta := models.Pagination{
		TotalRecords: totalRecords,
		TotalPages:   totalPages,
		CurrentPage:  page,
		PageSize:     pageSize,
	}
	return data, meta, nil
}

// func Paginate[T any](db *gorm.DB, data map[string]interface{}, param ...interface{}) ([]T, error) {
// 	page := data["index"]
// 	pageSize := data["size"]
// 	isPaginate := data["isPaginate"]
// 	var records []T
// 	if isPaginate == true {
// 		if page == nil {
// 			page = 1
// 		} else {
// 			page = int(page.(float64))
// 		}
// 		if pageSize == nil {
// 			pageSize = 20
// 		} else {
// 			pageSize = int(pageSize.(float64))
// 		}
// 		offset := (page.(int) - 1) * pageSize.(int)
// 		if len(param) > 0 {
// 			err := db.Where(param[0], param[1:]...).Limit(pageSize.(int)).Offset(offset).Find(&records).Error
// 			if err != nil {
// 				return nil, err
// 			}
// 		} else {
// 			err := db.Limit(pageSize.(int)).Offset(offset).Find(&records).Error
// 			if err != nil {
// 				return nil, err
// 			}

// 		}
// 	} else {
// 		if len(param) > 0 {
// 			err := db.Where(param[0], param[1:]...).Find(&records).Error
// 			if err != nil {
// 				return nil, err
// 			}
// 		} else {
// 			err := db.Find(&records).Error
// 			if err != nil {
// 				return nil, err
// 			}
// 		}
// 	}

// 	return records, nil
// }
