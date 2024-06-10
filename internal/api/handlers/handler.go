package handlers

import (
	"io"
	"minicommerce/internal/api/dto/request"
	"minicommerce/internal/api/dto/response"
	"minicommerce/internal/api/helpers"
	"minicommerce/internal/api/models"
	"minicommerce/internal/api/services"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ProductHandlerImpl struct {
	ProductService services.ProductServiceImpl
}

func NewProductHandlers(productService services.ProductServiceImpl) *ProductHandlerImpl {
	return &ProductHandlerImpl{ProductService: productService}
}

// CRUD API Product

// Add product form
func (h *ProductHandlerImpl) AddProduct(c *fiber.Ctx) error {
	// Parsing form data
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse form data",
		})
	}

	// Mendapatkan field lain dari form
	productName := form.Value["product_name"][0]
	typeProduct := form.Value["type_product"][0]
	description := form.Value["desc"][0]
	price, _ := strconv.Atoi(form.Value["price"][0])
	stock, _ := strconv.Atoi(form.Value["stock"][0])

	data := request.ProductAddRequest{
		ProductName: productName,
		TypeProduct: typeProduct,
		Desc:        description,
		Price:       price,
		Stock:       stock,
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(data)

	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// Mendapatkan file gambar
	file, err := c.FormFile("img_product")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to get image file",
		})
	}

	// Membaca file gambar sebagai byte array
	imageFile, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to open image file",
		})
	}
	defer imageFile.Close()

	imageData, err := io.ReadAll(imageFile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to read image file",
		})
	}

	product := models.Products{
		ProductName:  productName,
		TypeProduct:  typeProduct,
		Description:  description,
		Price:        price,
		Stock:        stock,
		ImageProduct: imageData, // Menyimpan gambar sebagai byte array
	}

	response, err := h.ProductService.AddProduct(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

// Add product json
// func (h *ProductHandlerImpl) AddProductJSON(c *fiber.Ctx) error {
// 	requests := new(request.ProductAddRequest)

// 	if err := c.BodyParser(requests); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	// Validation
// 	validate := validator.New()
// 	errValidate := validate.Struct(requests)
// 	if errValidate != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"message": "failed",
// 			"error":   errValidate.Error(),
// 		})
// 	}

// 	product := models.Products{
// 		ProductName:  requests.ProductName,
// 		ImageProduct: requests.ImgProduct,
// 		TypeProduct:  requests.TypeProduct,
// 		Description:  requests.Desc,
// 		Price:        requests.Price,
// 		Stock:        requests.Stock,
// 	}

// 	responses, err := h.ProductService.AddProduct(&product)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(responses)
// }

// Get product id
func (h *ProductHandlerImpl) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	product, err := h.ProductService.GetProductByID(idInt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	responses := response.ProductAllResponse{
		ID:          product.ID,
		ProductName: product.ProductName,
		ImgProduct:  product.ImageProduct,
		TypeProduct: product.TypeProduct,
		Desc:        product.Description,
		Price:       helpers.FormatMoney(product.Price),
		Stock:       product.Stock,
		CreatedAt:   product.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   product.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return c.Status(http.StatusOK).JSON(responses)
}

func (h *ProductHandlerImpl) UpdatedProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid product ID",
		})
	}

	// Parsing form data
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse form data",
		})
	}

	// Mendapatkan field lain dari form
	productName := form.Value["product_name"][0]
	typeProduct := form.Value["type_product"][0]
	description := form.Value["desc"][0]
	price, _ := strconv.Atoi(form.Value["price"][0])
	stock, _ := strconv.Atoi(form.Value["stock"][0])

	data := request.ProductAddRequest{
		ProductName: productName,
		TypeProduct: typeProduct,
		Desc:        description,
		Price:       price,
		Stock:       stock,
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(data)

	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// Mendapatkan file gambar, jika ada
	var imageData []byte
	file, err := c.FormFile("img_product")
	if err == nil {
		// Membaca file gambar sebagai byte array
		imageFile, err := file.Open()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to open image file",
			})
		}
		defer imageFile.Close()

		imageData, err = io.ReadAll(imageFile)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to read image file",
			})
		}
	}

	// Membuat objek produk baru
	newProduct := models.Products{
		ProductName: productName,
		TypeProduct: typeProduct,
		Description: description,
		Price:       price,
		Stock:       stock,
	}

	if len(imageData) > 0 {
		newProduct.ImageProduct = imageData // Menyimpan gambar sebagai byte array jika ada
	}

	// Create Validate

	// Memperbarui produk
	updatedProduct, err := h.ProductService.UpdatedProduct(uint(idInt), &newProduct)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to update product",
		})
	}

	// Menyusun respons
	response := &response.ProductUpdatedResponse{
		ID:          updatedProduct.ID,
		ProductName: updatedProduct.ProductName,
		ImgProduct:  updatedProduct.ImgProduct,
		TypeProduct: updatedProduct.TypeProduct,
		Desc:        updatedProduct.Desc,
		Price:       updatedProduct.Price,
		Stock:       updatedProduct.Stock,
		UpdatedAt:   updatedProduct.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}

// Updated product
// func (h *ProductHandlerImpl) UpdatedProduct(c *fiber.Ctx) error {
// 	id := c.Params("id")
// 	idInt, err := strconv.Atoi(id)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	requests := new(request.ProductUpdatedRequest)
// 	err = c.BodyParser(requests)

// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// validate := validator.New()
// 	errValidate := validate.Struct(requests)
// 	if errValidate != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"message": "failed",
// 			"error":   errValidate.Error(),
// 		})
// 	}

// 	_, err = h.ProductService.GetProductByID(idInt)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	productUpdated, err := h.ProductService.UpdatedProduct(uint(idInt), requests)
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"message": err.Error(),
// 		})
// 	}

// 	return c.Status(http.StatusOK).JSON(productUpdated)
// }

// Deleted product
func (h *ProductHandlerImpl) DeletedProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	_, err = h.ProductService.GetProductByID(idInt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	err = h.ProductService.DeletedProduct(uint(idInt))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "product unavailable",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "product successfully deleted",
	})
}

// Search product
func (h *ProductHandlerImpl) SearchProduct(c *fiber.Ctx) error {

	// parameter untuk mencari product dalam bentuk json
	var searchParams struct {
		ProductName string `json:"product_name"`
		TypeProduct string `json:"type_product"`
	}

	err := c.BodyParser(&searchParams)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if searchParams.ProductName == "" && searchParams.TypeProduct == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "At least one search parameter (product_name or type_product) must be provided",
		})
	}

	products, err := h.ProductService.SearchProduct(searchParams.ProductName, searchParams.TypeProduct)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "product unavailable",
		})
	}

	return c.Status(http.StatusOK).JSON(products)
}

// Get all product
func (h *ProductHandlerImpl) GetAllProduct(c *fiber.Ctx) error {

	product, err := h.ProductService.GetAllProduct()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(product)
}

// Detailed product
func (h *ProductHandlerImpl) DetailedProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid syntax variable",
		})
	}

	product, err := h.ProductService.GetProductByID(idInt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	responses := response.ProductDetailProduct{
		ProductName: product.ProductName,
		ImgProduct:  product.ImageProduct,
		TypeProduct: product.TypeProduct,
		Price:       helpers.FormatMoney(product.Price),
		Stock:       product.Stock,
		Desc:        product.Description,
	}

	return c.Status(http.StatusOK).JSON(responses)
}

// List products
func (h *ProductHandlerImpl) ListProduct(c *fiber.Ctx) error {

	product, err := h.ProductService.ListProduct()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"Product": product,
	})
}
