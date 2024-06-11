package config

import (
	"fmt"
	"log"
	"minicommerce/internal/api/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConfigDB struct {
	host     string
	port     string
	user     string
	password string
	dbname   string
}

func Config() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Gagal memuat file .env: %v", err)
	}

	config := ConfigDB{
		host:     os.Getenv("PGHOST"),
		port:     os.Getenv("PGPORT"),
		user:     os.Getenv("PGUSER"),
		password: os.Getenv("PGPASSWORD"),
		dbname:   os.Getenv("PGDATABASE"),
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.host, config.port, config.user, config.password, config.dbname)

	dbs, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	dbs.AutoMigrate(&models.Products{})

	return dbs, nil
}
