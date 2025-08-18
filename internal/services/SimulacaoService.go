package services

import (
	"fmt"

	dto "github.com/ferrariwill/marcenaria-backend/internal/DTO"
	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"github.com/ferrariwill/marcenaria-backend/internal/repositories"
)

type SimulacaoService struct {
	Repo *repositories.SimulacaoRepository
}

func Simulacao(repo *repositories.SimulacaoRepository) *SimulacaoService {
	return &SimulacaoService{Repo: repo}
}

func (s *SimulacaoService) CreateSimulacao(req *dto.SimulacaoDTO) ([]models.Componente, error) {
	modelo, err := s.Repo.BuscarModeloComRegras(req.ModeloID)
	if err != nil {
		return nil, err
	}

	var ids []uint
	for _, id := range req.Materiais {
		ids = append(ids, id)
	}

	placasMap, err := s.Repo.BuscarPlacasPorIds(ids)

	if err != nil {
		return nil, fmt.Errorf("erro ao buscar placas MDF: %w", err)
	}

	placas := make(map[string]models.PlacaMDF)

	for key, id := range req.Materiais {
		placa, ok := placasMap[id]
		if !ok {
			return nil, fmt.Errorf("placa MDF com ID %d não encontrada", id)
		}
		placas[key] = placa
	}

	item := models.ItemMontado{
		Largura:      req.Largura,
		Altura:       req.Altura,
		Profundidade: req.Profundidade,
		Espessura:    req.Espessura,
	}
	componentes := GerarComponentes(item, *modelo, placas)
	return componentes, nil
}
