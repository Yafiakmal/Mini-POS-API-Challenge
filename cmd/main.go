package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/db"
	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/handler"
	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/routes"
	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/service"

	// "github.com/yafiakmal/Mini-POS-API-Challenge/internal/db"

	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/repository"
)

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "production"
	}

	// Hanya load .env kalau di development
	if env == "development" {
		if err := godotenv.Load(); err != nil {
			log.Println("Warning: .env file not found, using system environment")
		} else {
			log.Println(".env file loaded")
		}
	}
	log.Printf("connecting to database")

	DB := db.InitDB()

	pdRepo := repository.NewProductRepository(DB)
	txRepo := repository.NewTransactionRepository(DB)
	pdService := service.NewProductService(pdRepo)
	txService := service.NewTransactionService(txRepo)
	pdHandler := handler.NewProductHandler(pdService)
	txHandler := handler.NewTransactionHandler(txService)

	r := gin.Default()

	// Setup routes
	routes.SetupProductRoutes(r, pdHandler)
	routes.SetupTransactionRoutes(r, txHandler)

	r.Run(":8080")

}
