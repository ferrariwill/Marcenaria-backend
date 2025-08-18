package services

import (
	"errors"

	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"github.com/ferrariwill/marcenaria-backend/internal/repositories"
)

type FileteService struct {
	Repo *repositories.FileteRepository
}

func Filete(repo *repositories.FileteRepository) *FileteService {
	return &FileteService{Repo: repo}
}

func (s *FileteService) ListarTodas() ([]models.Filete, error) {
	return s.Repo.ListarTodas()
}

func (s *FileteService) BuscarPorID(id uint) (*models.Filete, error) {
	f, err := s.Repo.BuscarPorID(id)
	if err != nil {
		return nil, errors.New("Filete não encontrada")
	}
	return f, nil
}

func (s *FileteService) Criar(f *models.Filete) error {
	err := validacoesFilete(f)
	if err != nil {
		return err
	}
	return s.Repo.Criar(f)
}

func (s *FileteService) Atualizar(id uint, input *models.Filete) (*models.Filete, error) {
	f, err := s.Repo.BuscarPorID(id)
	if err != nil {
		return nil, errors.New("Filete não encontrada")
	}

	err = validacoesFilete(input)
	if err != nil {
		return nil, err
	}

	f.ID = id
	f.Cor = input.Cor
	f.Largura = input.Largura
	f.PrecoMetro = input.PrecoMetro

	err = s.Repo.Atualizar(f)
	return f, err
}

func (s *FileteService) Deletar(id uint) error {
	return s.Repo.Deletar(id)
}

func validacoesFilete(f *models.Filete) error {

	if f.Cor == "" {
		return errors.New("nome inválido")
	}

	if f.Largura == 0 {
		return errors.New("largura inválida")
	}

	if f.PrecoMetro == 0 {
		return errors.New("preço por metro inválido")
	}
	return nil
}
