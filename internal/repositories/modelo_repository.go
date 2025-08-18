package repositories

import (
	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"gorm.io/gorm"
)

type ModeloRepository struct {
	DB *gorm.DB
}

func Modelo(db *gorm.DB) *ModeloRepository {
	return &ModeloRepository{DB: db}
}

func (r *ModeloRepository) ListarTodos() ([]models.ModeloMovel, error) {
	var modelos []models.ModeloMovel
	err := r.DB.Find(&modelos).Error
	return modelos, err
}

func (r *ModeloRepository) BuscarPorID(id uint) (*models.ModeloMovel, error) {
	var modelo models.ModeloMovel
	err := r.DB.First(&modelo, id).Error
	return &modelo, err
}

func (r *ModeloRepository) Criar(modelo *models.ModeloMovel) error {
	return r.DB.Create(modelo).Error
}

func (r *ModeloRepository) Atualizar(modelo *models.ModeloMovel) error {
	return r.DB.Save(modelo).Error
}

func (r *ModeloRepository) Deletar(id uint) error {
	return r.DB.Delete(&models.ModeloMovel{}, id).Error
}
