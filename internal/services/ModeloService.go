package services

import (
	"errors"

	dto "github.com/ferrariwill/marcenaria-backend/internal/DTO"
	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"github.com/ferrariwill/marcenaria-backend/internal/repositories"
)

type ModeloService struct {
	Repo *repositories.ModeloRepository
}

func Modelos(repo *repositories.ModeloRepository) *ModeloService {
	return &ModeloService{Repo: repo}
}

func (s *ModeloService) ListarTodos() ([]models.ModeloMovel, error) {
	return s.Repo.ListarTodos()
}

func (s *ModeloService) BuscarPorId(id uint) (*models.ModeloMovel, error) {
	m, err := s.Repo.BuscarPorID(id)

	if err != nil {
		return nil, errors.New("ferragem não encontrada")
	}

	return m, nil
}

func (s *ModeloService) Criar(dto *dto.CriarModeloDTO) error {
	m := &models.ModeloMovel{
		Nome:      dto.Nome,
		Descricao: dto.Descricao,
	}

	for _, r := range dto.Regras {
		m.Regras = append(m.Regras, models.RegraModelo{
			Nome:        r.Nome,
			LarguraExpr: r.LarguraExpr,
			AlturaExpr:  r.AlturaExpr,
			Quantidade:  r.Quantidade,
			TipoUso:     r.TipoUso,
		})
	}

	err := validacoesModelos(m)

	if err != nil {
		return err
	}

	return s.Repo.Criar(m)
}

func (s *ModeloService) Atualizar(id uint, dto *dto.CriarModeloDTO) (*models.ModeloMovel, error) {
	m, err := s.Repo.BuscarPorID(id)

	if err != nil {
		return nil, errors.New("Modelo não encontrada")
	}

	m.Nome = dto.Nome
	m.Descricao = dto.Descricao

	for _, r := range dto.Regras {
		m.Regras = append(m.Regras, models.RegraModelo{
			Nome:        r.Nome,
			LarguraExpr: r.LarguraExpr,
			AlturaExpr:  r.AlturaExpr,
			Quantidade:  r.Quantidade,
			TipoUso:     r.TipoUso,
		})
	}

	err = validacoesModelos(m)

	if err != nil {
		return nil, err
	}

	err = s.Repo.Atualizar(m)

	return m, err
}

func (s *ModeloService) Deletar(id uint) error {
	m, err := s.Repo.BuscarPorID(id)

	if err != nil {
		return errors.New("Modelo não encontrado")
	}
	return s.Repo.Deletar(m.ID)
}

func validacoesModelos(m *models.ModeloMovel) error {
	if m == nil {
		return errors.New("modelo não pode ser nulo")
	}

	if m.Nome == "" {
		return errors.New("nome não pode ser vazio")
	}

	if m.Regras == nil {
		return errors.New("regras não podem ser nulas")
	}

	if len(m.Regras) == 0 {
		return errors.New("regras não podem ser vazias")
	}

	return nil
}
