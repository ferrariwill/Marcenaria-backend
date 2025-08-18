package repositories

import (
	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"gorm.io/gorm"
)

type FerragemRepository struct {
	DB *gorm.DB
}

func Ferragem(db *gorm.DB) *FerragemRepository {
	return &FerragemRepository{DB: db}
}

func (r *FerragemRepository) ListarTodas() ([]models.Ferragem, error) {
	var ferragens []models.Ferragem
	err := r.DB.Find(&ferragens).Error
	return ferragens, err
}

func (r *FerragemRepository) BuscarPorID(id uint) (*models.Ferragem, error) {
	var ferragem models.Ferragem
	err := r.DB.First(&ferragem, id).Error
	return &ferragem, err
}

func (r *FerragemRepository) Criar(f *models.Ferragem) error {
	return r.DB.Create(f).Error
}

func (r *FerragemRepository) Atualizar(f *models.Ferragem) error {
	return r.DB.Save(f).Error
}

func (r *FerragemRepository) Deletar(id uint) error {
	return r.DB.Delete(&models.Ferragem{}, id).Error
}
