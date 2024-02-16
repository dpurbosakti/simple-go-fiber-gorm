package db

import (
	"github.com/dpurbosakti/fiber-gorm/models/user"
	"github.com/rs/zerolog/log"
)

func RunMigration() {
	err := DB.AutoMigrate(user.User{})
	if err != nil {
		log.Fatal().Err(err).Msg("failed to migrate")
	}

	log.Info().Msg("database migrated successfully")
}
