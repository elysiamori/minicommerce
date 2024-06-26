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
	
	form, err := c.MultipartForm()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to parse form data",
		})
	}

	productName := form.Value["product_name"][0]
	typeProduct := form.Value["type_product"][0]
	description := form.Value["desc"][0]
	sold, _ := strconv.Atoi(form.Value["sold"][0])
	location := form.Value["location"][0]
	price, _ := strconv.Atoi(form.Value["price"][0])
	stock, _ := strconv.Atoi(form.Value["stock"][0])

	data := request.ProductAddRequest{
		ProductName: productName,
		TypeProduct: typeProduct,
		Desc:        description,
		Sold:        sold,
		Location:    location,
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

	var imageData []byte
	file, err := c.FormFile("img_product")
	if err == nil {
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

	product := models.Products{
		ProductName:  productName,
		TypeProduct:  typeProduct,
		Description:  description,
		Sold:         sold,
		Location:     location,
		Price:        price,
		Stock:        stock,
		ImageProduct: imageData, 
	}

	response, err := h.ProductService.AddProduct(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

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
		Sold:        product.Sold,
		Location:    product.Location,
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

	productName := form.Value["product_name"][0]
	typeProduct := form.Value["type_product"][0]
	description := form.Value["desc"][0]
	sold, _ := strconv.Atoi(form.Value["sold"][0])
	location := form.Value["location"][0]
	price, _ := strconv.Atoi(form.Value["price"][0])
	stock, _ := strconv.Atoi(form.Value["stock"][0])

	data := request.ProductAddRequest{
		ProductName: productName,
		TypeProduct: typeProduct,
		Desc:        description,
		Sold:        sold,
		Location:    location,
		Price:       price,
		Stock:       stock,
	}

	// Validation
	validate := validator.New()
	errValidate := validate.Struct(data)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   errValidate.Error(),
		})
	}

	existingProduct, err := h.ProductService.GetProductByID(idInt)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Product not found",
		})
	}

	var imageData []byte
	file, err := c.FormFile("img_product")
	if err == nil {
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
	} else {
		imageData = existingProduct.ImageProduct
	}

	newProduct := models.Products{
		ProductName:  productName,
		TypeProduct:  typeProduct,
		Description:  description,
		Sold:         sold,
		Location:     location,
		Price:        price,
		Stock:        stock,
		ImageProduct: imageData, 
	}

	updatedProduct, err := h.ProductService.UpdatedProduct(uint(idInt), &newProduct)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to update product",
		})
	}
	
	response := &response.ProductUpdatedResponse{
		ID:          updatedProduct.ID,
		ProductName: updatedProduct.ProductName,
		ImgProduct:  updatedProduct.ImgProduct,
		TypeProduct: updatedProduct.TypeProduct,
		Desc:        updatedProduct.Desc,
		Sold:        updatedProduct.Sold,
		Location:    updatedProduct.Location,
		Price:       updatedProduct.Price,
		Stock:       updatedProduct.Stock,
		UpdatedAt:   updatedProduct.UpdatedAt,
	}

	return c.Status(fiber.StatusOK).JSON(response)
}


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

	productName := c.Query("product_name")
	typeProduct := c.Query("type_product")

	if productName == "" && typeProduct == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "At least one search parameter (product_name or type_product) must be provided",
		})
	}

	products, err := h.ProductService.SearchProduct(productName, typeProduct)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
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
		Sold:        product.Sold,
		Location:    product.Location,
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
