package config

import (
	"os"

	"github.com/juliocesar014/go-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	// Check if database file exists

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Infof("database file not found, creating...")
		// Create dabase file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Create Databases and Connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("db initialization error: %v", err)
		return nil, err
	}

	// Migrate database
	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("sqlite error migrate: %v", err)
		return nil, err
	}

	// Return the DB
	return db, nil
}
