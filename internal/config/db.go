package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Carrega o .env
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DB_HOST, DB_USER, DB_PASSWORD, DB_NAME, DB_PORT)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)

	}

	err = db.AutoMigrate(
		&models.Cliente{},
		&models.Componente{},
		&models.Ferragem{},
		&models.Filete{},
		&models.ItemMontado{},
		&models.ModeloMovel{},
		&models.Orcamento{},
		&models.PlacaMDF{},
		&models.Projeto{},
		&models.RegraModelo{},
	)

	if err != nil {
		log.Fatal("Erro ao migrar modelos:", err)
	}

	DB = db

}
