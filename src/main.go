package main

import (
	"fmt"
	"log"

	di "produtos-favoritos/src/api/container"
	"produtos-favoritos/src/api/router"
	handlers "produtos-favoritos/src/domain/interfaces/controllers"
	"produtos-favoritos/src/infrastructure/database/migrations"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Ecommerce Aiqfome Api
// @version 1.0
// @description Manage Customers, Whislist
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-api-key

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Run migrations
	migrations.RunMigrations()

	// Wire dependency injection
	container := di.BuildContainer()

	err = container.Invoke(func(engine *gin.Engine,
		customerHandler handlers.CustomerHandler,
		productHandler handlers.ProductHandler,
		wishlisthandler handlers.Wishlisthandler) {
		// Setup Gin router
		router.SetupRouter(engine,
			customerHandler,
			productHandler,
			wishlisthandler)

		// run server
		fmt.Println("Server running at http://localhost:8080")
		if err := engine.Run(":8080"); err != nil {
			log.Fatalf("could not start server: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
