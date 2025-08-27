package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/db"
	"github.com/yafiakmal/Mini-POS-API-Challenge/internal/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Printf("connecting to database")

	DB := db.InitDB()
	model.AutoMigrate(DB)

	DB.Create(&model.Product{
		Name:  "Papan Tulis",
		Price: 35000,
		Stock: 30,
	})

	product := model.Product{}
	DB.First(&product)
	log.Printf("Product: %v", product)

}
