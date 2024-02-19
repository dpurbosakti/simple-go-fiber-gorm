package db

import (
	zl "github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type Queries struct {
// 	*queries.UserQueries // load queries from User model
// }

// var (
// 	db  *gorm.DB
// 	err error
// )

func OpenDBConnection() *gorm.DB {
	// db logger
	// dbLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold: time.Second,
	// 		LogLevel:      logger.Info,
	// 		Colorful:      true,
	// 	},
	// )

	dsn := "postgresql://root:mokopass@localhost:5432/fiber_gorm?sslmode=disable"

	dialector := postgres.Open(dsn)

	db, err := gorm.Open(dialector)
	zl.Info().Msg("connecting to db...")
	if err != nil {
		zl.Fatal().Err(err).Msg("failed to connect to the database")
	}

	sqlDB, err := db.DB()
	if err != nil {
		zl.Fatal().Err(err).Msg("failed to get *sql.DB")
	}

	if err := sqlDB.Ping(); err != nil {
		zl.Fatal().Err(err).Msg("failed to ping database")
	}

	zl.Info().Msg("connected to the database successfully")
	return db
}
