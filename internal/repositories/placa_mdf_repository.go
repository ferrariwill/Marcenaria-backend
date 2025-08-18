package repositories

import (
	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"gorm.io/gorm"
)

type PlacaMDFRepository struct {
	DB *gorm.DB
}

func PlacaMDF(db *gorm.DB) *PlacaMDFRepository {
	return &PlacaMDFRepository{DB: db}
}

func (r *PlacaMDFRepository) FindAll() ([]models.PlacaMDF, error) {
	var placas []models.PlacaMDF
	err := r.DB.Find(&placas).Error
	return placas, err
}

func (r *PlacaMDFRepository) FindByID(id uint) (*models.PlacaMDF, error) {
	var placa models.PlacaMDF
	err := r.DB.First(&placa, id).Error
	return &placa, err
}

func (r *PlacaMDFRepository) Create(placaMDF *models.PlacaMDF) error {
	return r.DB.Create(placaMDF).Error
}

func (r *PlacaMDFRepository) Update(placa *models.PlacaMDF) error {
	return r.DB.Save(placa).Error
}

func (r *PlacaMDFRepository) Delete(id uint) error {
	return r.DB.Delete(&models.PlacaMDF{}, id).Error
}
