package main

import (
	"log"
	"minicommerce/config"
	"minicommerce/internal/api/app"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db, err := config.Config()
	if err != nil {
		panic(err)
	}

	app1 := fiber.New()

	app.SetupRoutes(app1, db)

	// Channel untuk sinyal shutdown
	done := make(chan bool, 1)

	// Menjalankan server dalam goroutine
	go func() {
		if err := app1.Listen(":3000"); err != nil {
			log.Fatalf("Failed to start server: %v", err)
		}
		done <- true
	}()

	// Menunggu sinyal shutdown
	<-done

}
