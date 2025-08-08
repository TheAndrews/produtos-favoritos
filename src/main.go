package main

import (
	"fmt"
	"log"

	di "produtos-favoritos/src/api/container"
	"produtos-favoritos/src/api/router"
	handlers "produtos-favoritos/src/domain/interfaces/controllers"
	"produtos-favoritos/src/infrastructure/config"
	"produtos-favoritos/src/infrastructure/database/migrations"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// @title Ecommerce Aiqfome Api
// @version 1.0
// @description Manage Customers, Whislist
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-Api-Key
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Wire dependency injection
	container := di.BuildContainer()
	// Run migrations
	// Run migrations using the *gorm.DB from the container
	err = container.Invoke(func(db *gorm.DB) {
		migrations.RunMigrations(db)
	})
	if err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
	err = container.Invoke(func(engine *gin.Engine,
		customerHandler handlers.CustomerHandler,
		productHandler handlers.ProductHandler,
		wishlisthandler handlers.WishlistHandler) {
		// Setup Gin router
		router.SetupRouter(engine,
			customerHandler,
			productHandler,
			wishlisthandler)

		// run server
		fmt.Printf("Server running at http://localhost:%s", config.APP_PORT)
		if err := engine.Run(fmt.Sprintf(":%s", config.APP_PORT)); err != nil {
			log.Fatalf("could not start server: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}
