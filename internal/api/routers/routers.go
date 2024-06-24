package routers

import (
	"minicommerce/internal/api/handlers"
	"minicommerce/internal/api/services"

	"github.com/gofiber/fiber/v2"
)

func Minicommerce(app *fiber.App, productServices *services.ProductServiceImpl) {

	productHandler := handlers.NewProductHandlers(*productServices)

	product := app.Group("/api")
	{
		// CRUD products
		product.Post("/products", productHandler.AddProduct)
		product.Get("/products", productHandler.GetAllProduct)
		product.Get("/products/:id", productHandler.GetProductByID)
		product.Put("/products/:id", productHandler.UpdatedProduct)
		product.Delete("/products/:id", productHandler.DeletedProduct)

		// Seacrh product
		product.Get("/search-products", productHandler.SearchProduct)

		// List products
		product.Get("/list-products", productHandler.ListProduct)

		// Get details product
		product.Get("/list-products/detail-product/:id", productHandler.DetailedProduct)
	}
}
