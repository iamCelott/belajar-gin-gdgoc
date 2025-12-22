package utils

import "belajar-gin/models"

func AutoMigrate() {
	db := DB()
	err := db.Migrator().AutoMigrate(
		&models.Book{},
	)
	if err != nil {
		panic("Gagal migrate")
	}
}
