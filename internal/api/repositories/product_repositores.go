package repositories

import (
	"minicommerce/internal/api/models"
	"strings"

	"gorm.io/gorm"
)

type ProductsRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductsRepositoryImpl {
	return &ProductsRepositoryImpl{DB: db}
}

// Menambahkan data produk
func (r *ProductsRepositoryImpl) AddProduct(product *models.Products) (*models.Products, error) {
	err := r.DB.Create(&product).Error
	if err != nil {
		return product, err
	}

	return product, nil
}

// Mendapatkan semua data produk
func (r *ProductsRepositoryImpl) GetAllProduct() ([]models.Products, error) {
	products := []models.Products{}

	err := r.DB.Find(&products).Error
	if err != nil {
		return []models.Products{}, err
	}

	return products, nil
}

// Menampilkan 1 product id yang dipilih
func (r *ProductsRepositoryImpl) GetProductByID(id int) (*models.Products, error) {
	product := models.Products{}

	err := r.DB.Where("id = ?", id).Take(&product).Error
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *ProductsRepositoryImpl) SearchProduct(productName, typeProduct string) ([]models.Products, error) {
	products := []models.Products{}
	query := r.DB

	if productName != "" {
		query = query.Where("LOWER(product_name) LIKE ?", "%"+strings.ToLower(productName)+"%")
	}

	if typeProduct != "" {
		query = query.Or("LOWER(type_product) LIKE ?", "%"+strings.ToLower(typeProduct)+"%")
	}

	err := query.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

// Memperbarui product
func (r *ProductsRepositoryImpl) UpdatedProduct(product *models.Products, newProduct *models.Products) (*models.Products, error) {
	err := r.DB.Model(&product).Updates(newProduct).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

// func (r *ProductsRepositoryImpl) UpdatedProduct(product *models.Products, newProduct *models.Products) (*models.Products, error) {
// 	err := r.DB.Model(&product).Updates(newProduct).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return product, nil
// }

// Menghapus product
func (r *ProductsRepositoryImpl) DeletedProduct(id uint) error {
	product := models.Products{}

	err := r.DB.Where("id = ?", id).Delete(&product).Error
	if err != nil {
		return err
	}

	return nil
}
