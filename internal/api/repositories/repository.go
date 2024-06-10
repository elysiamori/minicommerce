package repositories

import "minicommerce/internal/api/models"

type ProductRepository interface {
	AddProduct(product *models.Products) (*models.Products, error)
	GetAllProduct() ([]*models.Products, error)
	GetProductByID(id int) (*models.Products, error)
	SearchProduct(productName, typeProduct *models.Products) ([]models.Products, error)
	UpdatedProduct(product *models.Products, newProduct *models.Products) (*models.Products, error)
	DeletedProduct(id uint) error
}
