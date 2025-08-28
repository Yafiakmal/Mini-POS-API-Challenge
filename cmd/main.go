package main

import (
	"log"

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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
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
