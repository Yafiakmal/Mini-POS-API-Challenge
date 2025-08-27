package db

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/model"
)

var (
	db   *gorm.DB
	once sync.Once
)

func InitDB() *gorm.DB {
	once.Do(func() {
		dbConfig := GetDBConfig()
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			dbConfig.Host, dbConfig.User, dbConfig.Password,
			dbConfig.DBName, dbConfig.Port, dbConfig.SSLMode, dbConfig.TimeZone,
		)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("failed to connect database: %v", err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("failed to get sql.DB: %v", err)
		}

		poolConfig := GetPoolConfig()
		sqlDB.SetMaxIdleConns(poolConfig.MaxIdleConns)
		sqlDB.SetMaxOpenConns(poolConfig.MaxOpenConns)
		sqlDB.SetConnMaxLifetime(poolConfig.ConnMaxLifetime)

		// AutoMigrate untuk semua model
		if err := model.AutoMigrate(db); err != nil {
			log.Fatal(err)
		}
		log.Println("Database migrated successfully")
	})

	return db
}
