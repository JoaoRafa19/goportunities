package config

import (
	"os"

	"github.com/JoaoRafa19/goplaning/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
) 	 

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	//check database 
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.InfoF("Database file not found")
		// Create database file
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err 
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.ErrF("sqlite opening error: %v", err)
		return nil, err
	}
	logger.InfoF("SQLite opened: %v", db.Config)
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.ErrF("SQLite auto migration error: %v", err)
		return nil, err
	}
	return db, nil

}
