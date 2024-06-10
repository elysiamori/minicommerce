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
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASS"),
		dbname:   os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.user, config.password, config.host, config.port, config.dbname)

	dbs, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	dbs.AutoMigrate(&models.Products{})

	return dbs, nil
}
