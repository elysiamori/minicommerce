package config

import (
	"minicommerce/internal/api/models"
	
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func Config() (*gorm.DB, error) {

	dsn := "postgresql://postgres:********@roundhouse.proxy.rlwy.net:43089/railway"

	dbs, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	dbs.AutoMigrate(&models.Products{})

	return dbs, nil
}
