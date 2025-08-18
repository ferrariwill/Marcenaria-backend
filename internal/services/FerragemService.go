package services

import (
	"errors"

	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"github.com/ferrariwill/marcenaria-backend/internal/repositories"
)

type FerragemService struct {
	Repo *repositories.FerragemRepository
}

func Ferragem(repo *repositories.FerragemRepository) *FerragemService {
	return &FerragemService{Repo: repo}
}

func (s *FerragemService) ListarTodas() ([]models.Ferragem, error) {
	return s.Repo.ListarTodas()
}

func (s *FerragemService) BuscarPorID(id uint) (*models.Ferragem, error) {
	f, err := s.Repo.BuscarPorID(id)
	if err != nil {
		return nil, errors.New("ferragem não encontrada")
	}
	return f, nil
}

func (s *FerragemService) Criar(f *models.Ferragem) error {
	err := validacoesFerragem(f)
	if err != nil {
		return err
	}
	return s.Repo.Criar(f)
}

func (s *FerragemService) Atualizar(id uint, input *models.Ferragem) (*models.Ferragem, error) {
	f, err := s.Repo.BuscarPorID(id)
	if err != nil {
		return nil, errors.New("ferragem não encontrada")
	}

	err = validacoesFerragem(input)
	if err != nil {
		return nil, err
	}

	f.ID = id
	f.Nome = input.Nome
	f.Tipo = input.Tipo
	f.PrecoUnitario = input.PrecoUnitario

	err = s.Repo.Atualizar(f)
	return f, err
}

func (s *FerragemService) Deletar(id uint) error {
	return s.Repo.Deletar(id)
}

func validacoesFerragem(f *models.Ferragem) error {

	if f.PrecoUnitario <= 0 {
		return errors.New("preço unitário inválido")
	}

	if f.Nome == "" {
		return errors.New("nome inválido")
	}

	if f.Tipo == "" {
		return errors.New("tipo inválido")
	}

	return nil
}
