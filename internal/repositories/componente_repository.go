package repositories

import (
	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"gorm.io/gorm"
)

type ComponenteRepository struct {
	DB *gorm.DB
}

func Componente(db *gorm.DB) *ComponenteRepository {
	return &ComponenteRepository{DB: db}
}

func (r *ComponenteRepository) ListarTodas() ([]models.Componente, error) {
	var componentes []models.Componente
	err := r.DB.Find(&componentes).Error
	return componentes, err
}

func (r *ComponenteRepository) BuscarPorID(id uint) (*models.Componente, error) {
	var componentes models.Componente
	err := r.DB.First(&componentes, id).Error
	return &componentes, err
}

func (r *ComponenteRepository) Criar(f *models.Componente) error {
	return r.DB.Create(f).Error
}

func (r *ComponenteRepository) Atualizar(f *models.Componente) error {
	return r.DB.Save(f).Error
}

func (r *ComponenteRepository) Deletar(id uint) error {
	return r.DB.Delete(&models.Componente{}, id).Error
}
