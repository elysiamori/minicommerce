package models

import (
	"time"
)

type Products struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	ProductName  string    `gorm:"char not null" json:"product_name" valid:"required"`
	ImageProduct []byte    `gorm:"type:bytea" json:"img_product" valid:"required"`
	TypeProduct  string    `gorm:"not null" json:"type_product" valid:"required"`
	Description  string    `gorm:"not null" json:"desc" valid:"required"`
	Price        int       `gorm:"not null" json:"price" valid:"required"`
	Stock        int       `gorm:"not null" json:"stock" valid:"required"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at" time_format:"2006-01-02 15:04:05"`
}
