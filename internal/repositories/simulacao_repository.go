package repositories

import (
	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"gorm.io/gorm"
)

type SimulacaoRepository struct {
	DB *gorm.DB
}

func Simulacao(db *gorm.DB) *SimulacaoRepository {
	return &SimulacaoRepository{DB: db}
}

func (r *SimulacaoRepository) BuscarModeloComRegras(modeloId uint) (*models.ModeloMovel, error) {
	var modelo models.ModeloMovel
	err := r.DB.Preload("Regras").First(&modelo, modeloId).Error
	if err != nil {
		return nil, err
	}
	return &modelo, nil
}

func (r *SimulacaoRepository) BuscarPlacasPorIds(ids []uint) (map[uint]models.PlacaMDF, error) {
	var placas []models.PlacaMDF

	err := r.DB.Where("id IN ?", ids).Find(&placas).Error
	if err != nil {
		return nil, err
	}

	placasMap := make(map[uint]models.PlacaMDF)
	for _, placa := range placas {
		placasMap[placa.ID] = placa
	}

	return placasMap, nil
}
