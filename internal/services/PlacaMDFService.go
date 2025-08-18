package services

import (
	"errors"

	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"github.com/ferrariwill/marcenaria-backend/internal/repositories"
)

type PlacaMDFService struct {
	Repo *repositories.PlacaMDFRepository
}

func PlacaMDF(repo *repositories.PlacaMDFRepository) *PlacaMDFService {
	return &PlacaMDFService{Repo: repo}
}

func (s *PlacaMDFService) ListarTodas() ([]models.PlacaMDF, error) {
	return s.Repo.FindAll()
}

func (s *PlacaMDFService) BuscaPorId(id uint) (*models.PlacaMDF, error) {
	placa, err := s.Repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return placa, nil
}

func (s *PlacaMDFService) Criar(placa *models.PlacaMDF) error {

	erro := validacoesDePlacaMDF(placa)

	if erro != nil {
		return erro
	}

	return s.Repo.Create(placa)
}

func (s *PlacaMDFService) Atualizar(id uint, input *models.PlacaMDF) (*models.PlacaMDF, error) {

	erro := validacoesDePlacaMDF(input)

	if erro != nil {
		return nil, erro
	}

	placa, err := s.Repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	placa.ID = id
	placa.Cor = input.Cor
	placa.Espessura = input.Espessura
	placa.Altura = input.Altura
	placa.Largura = input.Largura
	placa.PrecoUnitario = input.PrecoUnitario

	err = s.Repo.Update(placa)
	return placa, err

}

func (s *PlacaMDFService) Deletar(id uint) error {
	placa, err := s.Repo.FindByID(id)
	if err != nil {
		return err
	}

	return s.Repo.Delete(placa.ID)
}

func validacoesDePlacaMDF(placa *models.PlacaMDF) error {
	if placa.Cor == "" {
		return errors.New("A cor não pode ser vazia!")
	}

	if placa.Espessura == 0 {
		return errors.New("A espessura não pode ser vazia!")
	}

	if placa.Altura == 0 {
		return errors.New("A altura não pode ser vazia!")
	}

	if placa.Largura == 0 {
		return errors.New("A largura não pode ser vazia!")
	}

	if placa.PrecoUnitario == 0 {
		return errors.New("O preço não pode ser vazio!")
	}

	return nil
}
