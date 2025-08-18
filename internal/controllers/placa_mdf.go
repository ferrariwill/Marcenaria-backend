package controllers

import (
	"net/http"
	"strconv"

	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"github.com/ferrariwill/marcenaria-backend/internal/repositories"
	"github.com/ferrariwill/marcenaria-backend/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var placaService *services.PlacaMDFService

func InitPlacaController(db *gorm.DB) {
	repo := repositories.PlacaMDF(db)
	placaService = services.PlacaMDF(repo)
}

func GetPlacas(c *gin.Context) {

	placas, err := placaService.ListarTodas()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, placas)

}

func GetPlacaById(c *gin.Context) {

	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(idInt)
	placa, err := placaService.BuscaPorId(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, placa)

}

func CriarPlaca(c *gin.Context) {
	var placa models.PlacaMDF
	erro := c.ShouldBindJSON(&placa)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}
	erro = placaService.Criar(&placa)

	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erro.Error()})
		return
	}

	c.JSON(http.StatusCreated, placa)
}

func DeletarPlaca(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(idInt)
	erro := placaService.Deletar(id)

	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erro.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Placa deletada com sucesso"})
}

func AlterarPlaca(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(idInt)
	var placa models.PlacaMDF
	erro := c.ShouldBindJSON(&placa)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	placaAtualizada, erro := placaService.Atualizar(id, &placa)

	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erro.Error()})
		return
	}

	c.JSON(http.StatusOK, placaAtualizada)
}
