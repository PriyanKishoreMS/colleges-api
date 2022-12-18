package config

import (
	"fmt"
	"os"

	"github.com/PriyanKishoreMS/colleges-list-api/entities"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

func Connect() error {
	var err error

	envErr := godotenv.Load(".env")
	if envErr != nil {
		fmt.Println("Error loading .env file", envErr)
		os.Exit(1)
	}

	DATABASE_URI := os.Getenv("DATABASE_URI")

	Db, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		return fmt.Errorf("error opening database connection: %v", err)
	}

	fmt.Println("Database connection opened")

	db, err := Db.DB()

	if err != nil {
		return fmt.Errorf("error getting *sql.DB from gorm.DB: %v", err)
	}

	// Ping database to check the connection
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %v", err)
	}

	Db.AutoMigrate(&entities.College{})
	fmt.Println("Database migrated")

	return nil
}
