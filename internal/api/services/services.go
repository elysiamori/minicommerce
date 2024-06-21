package services

import (
	"minicommerce/internal/api/dto/response"
	"minicommerce/internal/api/helpers"
	"minicommerce/internal/api/models"
	"minicommerce/internal/api/repositories"
)

type ProductServiceImpl struct {
	ProductRepo repositories.ProductsRepositoryImpl
}

func NewProductServices(productRepo repositories.ProductsRepositoryImpl) *ProductServiceImpl {
	return &ProductServiceImpl{ProductRepo: productRepo}
}

// Add product services
func (s *ProductServiceImpl) AddProduct(product *models.Products) (*response.ProductAddResponse, error) {
	product, err := s.ProductRepo.AddProduct(product)
	if err != nil {
		return nil, err
	}

	response := &response.ProductAddResponse{
		ID:          product.ID,
		ProductName: product.ProductName,
		ImgProduct:  product.ImageProduct,
		TypeProduct: product.TypeProduct,
		Desc:        product.Description,
		Sold:        product.Sold,
		Location:    product.Location,
		Price:       helpers.FormatMoney(product.Price),
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

// Get products id services
func (s *ProductServiceImpl) GetProductByID(id int) (*models.Products, error) {
	product, err := s.ProductRepo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// List products services
func (s *ProductServiceImpl) ListProduct() ([]response.ProductListResponse, error) {
	products, err := s.ProductRepo.GetAllProduct()
	if err != nil {
		return nil, err
	}

	responses := []response.ProductListResponse{}

	for _, product := range products {
		responses = append(responses, response.ProductListResponse{
			ID:          int(product.ID),
			ProductName: product.ProductName,
			ImgProduct:  product.ImageProduct,
			TypeProduct: product.TypeProduct,
			Sold:        product.Sold,
			Location:    product.Location,
			Price:       helpers.FormatMoney(product.Price),
			Stock:       product.Stock,
		})
	}

	return responses, nil
}

// Search product services
func (s *ProductServiceImpl) SearchProduct(productName, typeProduct string) ([]response.ProductSearchResponse, error) {
	products, err := s.ProductRepo.SearchProduct(productName, typeProduct)
	if err != nil {
		return nil, err
	}

	responses := []response.ProductSearchResponse{}

	for _, product := range products {
		responses = append(responses, response.ProductSearchResponse{
			ProductName: product.ProductName,
			ImgProduct:  product.ImageProduct,
			TypeProduct: product.TypeProduct,
			Desc:        product.Description,
			Sold:        product.Sold,
			Location:    product.Location,
			Price:       helpers.FormatMoney(product.Price),
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return responses, nil
}

// All products as an editor services
func (s *ProductServiceImpl) GetAllProduct() ([]response.ProductAllResponse, error) {
	products, err := s.ProductRepo.GetAllProduct()
	if err != nil {
		return nil, err
	}

	responses := []response.ProductAllResponse{}

	for _, product := range products {
		responses = append(responses, response.ProductAllResponse{
			ID:          product.ID,
			ProductName: product.ProductName,
			ImgProduct:  product.ImageProduct,
			TypeProduct: product.TypeProduct,
			Desc:        product.Description,
			Sold:        product.Sold,
			Location:    product.Location,
			Price:       helpers.FormatMoney(product.Price),
			Stock:       product.Stock,
			CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   product.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return responses, nil
}

// Update product services
func (s *ProductServiceImpl) UpdatedProduct(productID uint, newProduct *models.Products) (*response.ProductUpdatedResponse, error) {

	oldProduct, err := s.ProductRepo.GetProductByID(int(productID))
	if err != nil {
		return nil, err
	}

	
	updatedProduct, err := s.ProductRepo.UpdatedProduct(oldProduct, newProduct)
	if err != nil {
		return nil, err
	}
	
	
	response := &response.ProductUpdatedResponse{
		ID:          updatedProduct.ID,
		ProductName: updatedProduct.ProductName,
		ImgProduct:  updatedProduct.ImageProduct,
		TypeProduct: updatedProduct.TypeProduct,
		Desc:        updatedProduct.Description,
		Sold:        updatedProduct.Sold,
		Location:    updatedProduct.Location,
		Price:       helpers.FormatMoney(updatedProduct.Price),
		Stock:       updatedProduct.Stock,
		UpdatedAt:   updatedProduct.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

// Deleted prodcuct services
func (s *ProductServiceImpl) DeletedProduct(productID uint) error {
	product_ID, err := s.ProductRepo.GetProductByID(int(productID))
	if err != nil {
		return err
	}

	if product_ID.ID == 0 {
		return err
	}

	err = s.ProductRepo.DeletedProduct(productID)
	if err != nil {
		return err
	}

	return nil
}
