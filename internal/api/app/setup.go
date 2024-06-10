package app

import (
	"minicommerce/internal/api/repositories"
	"minicommerce/internal/api/routers"
	"minicommerce/internal/api/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	productsRepository := repositories.NewProductRepository(db)
	productService := services.NewProductServices(*productsRepository)

	// untuk mengizinkan CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "*",
		AllowHeaders: "*",
	}))

	// untuk menampilkan log
	app.Use(logger.New())

	routers.Minicommerce(app, productService)
}
