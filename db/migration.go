package db

import (
	"github.com/dpurbosakti/fiber-gorm/app/models"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) {
	err := db.AutoMigrate(models.User{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to migrate")
	}

	log.Info().Msg("database migrated successfully")
}
