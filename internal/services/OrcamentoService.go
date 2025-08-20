package services

import (
	"fmt"
	"time"

	dto "github.com/ferrariwill/marcenaria-backend/internal/DTO"
	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"github.com/ferrariwill/marcenaria-backend/internal/repositories"
)

type OrcamentoService struct {
	PlacaRepo    repositories.PlacaMDFRepository
	FerragemRepo repositories.FerragemRepository
	FileteRepo   repositories.FileteRepository
	ModeloRepo   repositories.ModeloRepository
}

func Orcamento(placaRepo repositories.PlacaMDFRepository,
	ferragemRepo repositories.FerragemRepository,
	fileteRepo repositories.FileteRepository,
	modeloRepo repositories.ModeloRepository) *OrcamentoService {
	return &OrcamentoService{
		PlacaRepo:    placaRepo,
		FerragemRepo: ferragemRepo,
		FileteRepo:   fileteRepo,
		ModeloRepo:   modeloRepo,
	}
}

func (s *OrcamentoService) Calcular(dto dto.OrcamentoDTO) (*models.Orcamento, error) {
	// Montar ItemMontado
	item := models.ItemMontado{
		Largura:      dto.Largura,
		Altura:       dto.Altura,
		Profundidade: dto.Profundidade,
		Espessura:    dto.Espessura,
	}

	//Buscar modelo do móvel
	modelo, err := s.ModeloRepo.BuscarPorID(dto.ModeloID)
	if err != nil {
		return nil, fmt.Errorf("modelo não encontrado: %w", err)
	}

	//Montar mapa de placas MDF
	mapaPlaca := make(map[string]models.PlacaMDF)
	for nome, id := range dto.Materiais {
		placa, err := s.PlacaRepo.FindByID(id)
		if err != nil {
			return nil, fmt.Errorf("placa não encontrada: %w", err)
		}
		mapaPlaca[nome] = *placa
	}

	//Gerar componentes
	componentes := GerarComponentes(item, *modelo, mapaPlaca)

	//Calcular custo mdf
	var custoMDF float64
	for _, comp := range componentes {
		placa := mapaPlaca[comp.TipoMaterial]
		area := (comp.Largura / 1000) * (comp.Altura / 1000)
		precoM2 := placa.PrecoUnitario / ((placa.Largura / 1000) * (placa.Altura / 1000))
		custoMDF += area * precoM2

	}

	//Calcular Ferragens
	var custoFerragens float64
	for _, id := range dto.Ferragens {
		ferragem, err := s.FerragemRepo.BuscarPorID(id)
		if err != nil {
			return nil, fmt.Errorf("ferragem não encontrada: %w", err)
		}
		custoFerragens += ferragem.PrecoUnitario
	}

	//Calcular Filetes
	var custoFiletes float64
	for _, id := range dto.Filetes {
		filete, err := s.FileteRepo.BuscarPorID(id)
		if err != nil {
			return nil, fmt.Errorf("filete não encontrado: %w", err)
		}
		custoFiletes += filete.PrecoMetro
	}

	// Somar e aplicar margem de lucro
	custoTotal := custoMDF + custoFerragens + custoFiletes + dto.CustoExtras
	precoFinal := custoTotal * (1 + dto.MargemLucro/100)

	//Montar Orcamento
	orcamento := models.Orcamento{
		Descricao:      dto.Descricao,
		Componentes:    componentes,
		CustoMateriais: custoMDF,
		CustoFerragens: custoFerragens,
		CustoFiletes:   custoFiletes,
		CustosExtras:   dto.CustoExtras,
		MargemLucro:    dto.MargemLucro,
		PrecoFinal:     precoFinal,
		DataCriacao:    time.Now(),
		ClienteID:      dto.ClienteID,
	}

	return &orcamento, nil

}
