package routes

import (
	"github.com/ferrariwill/marcenaria-backend/internal/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	controllers.InitPlacaController(db)
	controllers.InitFerragemController(db)
	controllers.InitFileteController(db)
	controllers.InitModeloController(db)

	placas := router.Group("/placas")
	{
		placas.GET("/", controllers.GetPlacas)
		placas.GET("/:id", controllers.GetPlacaById)
		placas.POST("/", controllers.CriarPlaca)
		placas.PUT("/:id", controllers.AlterarPlaca)
		placas.DELETE("/:id", controllers.DeletarPlaca)
	}

	ferragens := router.Group("/ferragens")
	{
		ferragens.GET("/", controllers.GetFerragens)
		ferragens.GET("/:id", controllers.GetFerragemById)
		ferragens.POST("/", controllers.CriarFerragem)
		ferragens.PUT("/:id", controllers.AlterarFerragem)
		ferragens.DELETE("/:id", controllers.DeletarFerragem)
	}

	filetes := router.Group("/filetes")
	{
		filetes.GET("/", controllers.GetFiletes)
		filetes.GET("/:id", controllers.GetFileteById)
		filetes.POST("/", controllers.CriarFilete)
		filetes.PUT("/:id", controllers.AlterarFilete)
		filetes.DELETE("/:id", controllers.DeletarFilete)
	}

	modelos := router.Group("/modelos")
	{
		modelos.GET("/", controllers.GetModelos)
		modelos.GET("/:id", controllers.GetModeloById)
		modelos.POST("/", controllers.CriarModelo)
		modelos.PUT("/:id", controllers.AlterarModelo)
		modelos.DELETE("/:id", controllers.DeletarModelo)
	}
}
