package main

import (
	"github.com/ferrariwill/marcenaria-backend/internal/config"
	"github.com/ferrariwill/marcenaria-backend/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.InitDB()

	router := gin.Default()

	routes.SetupRoutes(router, config.DB)

	router.Run(":8080")
}
