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

var FileteService *services.FileteService

func InitFileteController(db *gorm.DB) {
	repo := repositories.Filete(db)
	FileteService = services.Filete(repo)
}

func GetFiletes(c *gin.Context) {
	filetes, err := FileteService.ListarTodas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, filetes)
}

func GetFileteById(c *gin.Context) {

	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(idInt)
	Filete, err := FileteService.BuscarPorID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, Filete)

}

func CriarFilete(c *gin.Context) {
	var Filete models.Filete

	erro := c.ShouldBindJSON(&Filete)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}
	erro = FileteService.Criar(&Filete)

	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erro.Error()})
		return
	}

	c.JSON(http.StatusCreated, Filete)
}

func DeletarFilete(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(idInt)
	erro := FileteService.Deletar(id)

	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erro.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Fielte deletado com sucesso"})
}

func AlterarFilete(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(idInt)
	var Filete models.Filete
	erro := c.ShouldBindJSON(&Filete)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	FileteAtualizada, erro := FileteService.Atualizar(id, &Filete)

	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erro.Error()})
		return
	}

	c.JSON(http.StatusOK, FileteAtualizada)
}
