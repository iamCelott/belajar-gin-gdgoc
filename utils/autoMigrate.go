package utils

import "belajar-gin/models"

func AutoMigrate() {
	db := DB()
	err := db.Migrator().AutoMigrate(
		&models.File{},
		&models.Category{},
		&models.Book{},
	)
	if err != nil {
		panic(err)
	}
}
