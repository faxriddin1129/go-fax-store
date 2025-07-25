package migrations

import (
	"gofax-store/internal/models"
	"gorm.io/gorm"
)

func MigrateAll(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Files{},
	)
}
