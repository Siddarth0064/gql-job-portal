package database

import (
	"context"
	"fmt"
	custommodel "project-gql/models"
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ======================= OPENING DATABASE CONNECTION ==========================================
func Open() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ========================== CONNECTION TO THE DATABASE =========================================
func Connection() (*gorm.DB, error) {
	log.Info().Msg("main : Started : Initializing db support")
	db, err := Open()
	if err != nil {
		return nil, fmt.Errorf("connecting to db %w", err)
	}
	pg, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance %w ", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = pg.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("database is not connected: %w ", err)
	}

	err = db.Migrator().AutoMigrate(&custommodel.User{})
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %w ", err)
	}
	err = db.Migrator().AutoMigrate(&custommodel.Company{})
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %w ", err)
	}
	err = db.Migrator().AutoMigrate(&custommodel.Job{})
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %w ", err)
	}
	return db, nil
}
