package repositories

import (
	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"gorm.io/gorm"
)

type FileteRepository struct {
	DB *gorm.DB
}

func Filete(db *gorm.DB) *FileteRepository {
	return &FileteRepository{DB: db}
}

func (r *FileteRepository) ListarTodas() ([]models.Filete, error) {
	var ferragens []models.Filete
	err := r.DB.Find(&ferragens).Error
	return ferragens, err
}

func (r *FileteRepository) BuscarPorID(id uint) (*models.Filete, error) {
	var Filete models.Filete
	err := r.DB.First(&Filete, id).Error
	return &Filete, err
}

func (r *FileteRepository) Criar(f *models.Filete) error {
	return r.DB.Create(f).Error
}

func (r *FileteRepository) Atualizar(f *models.Filete) error {
	return r.DB.Save(f).Error
}

func (r *FileteRepository) Deletar(id uint) error {
	return r.DB.Delete(&models.Filete{}, id).Error
}
