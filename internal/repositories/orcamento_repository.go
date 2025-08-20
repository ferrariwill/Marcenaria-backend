package repositories

import (
	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"gorm.io/gorm"
)

type OrcamentoRepository struct {
	DB *gorm.DB
}

func Orcamento(db *gorm.DB) *OrcamentoRepository {
	return &OrcamentoRepository{DB: db}
}

func (o *OrcamentoRepository) ListarTodos() ([]models.Orcamento, error) {
	var orcamentos []models.Orcamento
	err := o.DB.Find(&orcamentos).Error

	if err != nil {
		return nil, err
	}
	return orcamentos, nil
}

func (o *OrcamentoRepository) BuscarPorId(id uint) (*models.Orcamento, error) {
	var orcamento models.Orcamento
	err := o.DB.First(&orcamento, id).Error

	if err != nil {
		return nil, err
	}
	return &orcamento, nil
}

func (o *OrcamentoRepository) BuscarPorClienteId(id uint) ([]models.Orcamento, error) {
	var orcamentos []models.Orcamento
	err := o.DB.Where("cliente_id = ?", id).Find(&orcamentos).Error

	if err != nil {
		return nil, err
	}
	return orcamentos, nil
}

func (o *OrcamentoRepository) Criar(orcamento *models.Orcamento) error {
	err := o.DB.Create(orcamento).Error

	if err != nil {
		return err
	}
	return nil
}

func (o *OrcamentoRepository) Atualizar(orcamento *models.Orcamento) error {
	err := o.DB.Save(orcamento).Error

	if err != nil {
		return err
	}
	return nil
}

func (o *OrcamentoRepository) Deletar(id uint) error {
	err := o.DB.Delete(&models.Orcamento{}, id).Error

	if err != nil {
		return err
	}
	return nil
}
