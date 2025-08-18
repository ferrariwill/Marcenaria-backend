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

var ferragemService *services.FerragemService

func InitFerragemController(db *gorm.DB) {
	repo := repositories.Ferragem(db)
	ferragemService = services.Ferragem(repo)
}

func GetFerragens(c *gin.Context) {
	ferragens, err := ferragemService.ListarTodas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ferragens)
}

func GetFerragemById(c *gin.Context) {

	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(idInt)
	ferragem, err := ferragemService.BuscarPorID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, ferragem)

}

func CriarFerragem(c *gin.Context) {
	var ferragem models.Ferragem

	erro := c.ShouldBindJSON(&ferragem)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}
	erro = ferragemService.Criar(&ferragem)

	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erro.Error()})
		return
	}

	c.JSON(http.StatusCreated, ferragem)
}

func DeletarFerragem(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(idInt)
	erro := ferragemService.Deletar(id)

	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erro.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ferragem deletada com sucesso"})
}

func AlterarFerragem(c *gin.Context) {
	idInt, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uint(idInt)
	var ferragem models.Ferragem
	erro := c.ShouldBindJSON(&ferragem)

	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
		return
	}

	ferragemAtualizada, erro := ferragemService.Atualizar(id, &ferragem)

	if erro != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": erro.Error()})
		return
	}

	c.JSON(http.StatusOK, ferragemAtualizada)
}
