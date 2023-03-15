package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/alikhanturusbekov/redirects/internal/app/model"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	err = db.AutoMigrate(&model.Redirect{})
	if err != nil {
		log.Fatal("Failed to migrate to database. \n", err)
	}

	DB = Dbinstance{
		Db: db,
	}
}
