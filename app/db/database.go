package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/s-bose/project-mgmt-go/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db *gorm.DB
}

func ConnectDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to initialize database")
	}

	log.Println("Connected to database")
	return database
}

// add newly defined models here
func Migrate(d *Database) {
	if d.Db.Debug().AutoMigrate(&models.User{}) != nil {
		panic("Failed to migrate ORM models")
	}
	log.Println("Successfully migrated models")
}

func InitDatabase() *Database {
	godotenv.Load()
	gormDb := ConnectDatabase()
	db := &Database{Db: gormDb}
	Migrate(db)
	return db
}
