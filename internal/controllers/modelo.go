package controllers

import (
	"net/http"
	"strconv"

	dto "github.com/ferrariwill/marcenaria-backend/internal/DTO"
	"github.com/ferrariwill/marcenaria-backend/internal/repositories"
	"github.com/ferrariwill/marcenaria-backend/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var modeloService *services.ModeloService

func InitModeloController(db *gorm.DB) {
	repo := repositories.Modelo(db)
	modeloService = services.Modelos(repo)
}

func GetModelos(c *gin.Context) {
	modelos, err := modeloService.ListarTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, modelos)
}

func GetModeloById(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	id := uint(idInt)

	if id == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id inválido"})
		return
	}

	modelo, err := modeloService.BuscarPorId(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, modelo)
}

func CriarModelo(c *gin.Context) {
	var dto dto.CriarModeloDTO
	err := c.ShouldBindJSON(&dto)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = modeloService.Criar(&dto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Modelo criado com sucesso!"})
}

func AlterarModelo(c *gin.Context) {
	var dto dto.CriarModeloDTO
	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(idInt)
	erro := c.ShouldBindJSON(&dto)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	modeloAtualizado, erro := modeloService.Atualizar(id, &dto)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}
	c.JSON(http.StatusOK, modeloAtualizado)
}

func DeletarModelo(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(idInt)

	erro := modeloService.Deletar(id)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Modelo deletado com sucesso!"})
}
