package services

import (
	utils "github.com/ferrariwill/marcenaria-backend/Utils"
	"github.com/ferrariwill/marcenaria-backend/internal/models"
	"github.com/ferrariwill/marcenaria-backend/internal/repositories"
)

type ComponenteService struct {
	Repo *repositories.ComponenteRepository
}

func Componente(repo *repositories.ComponenteRepository) *ComponenteService {
	return &ComponenteService{Repo: repo}
}

func GerarComponentes(item models.ItemMontado, modelo models.ModeloMovel, placas map[string]models.PlacaMDF) []models.Componente {
	var componentes []models.Componente

	vars := map[string]interface{}{
		"largura":      item.Largura,
		"altura":       item.Altura,
		"profundidade": item.Profundidade,
		"espessura":    item.Espessura,
	}

	for _, regra := range modelo.Regras {
		largura := utils.AvaliarExpressao(regra.LarguraExpr, vars)
		altura := utils.AvaliarExpressao(regra.AlturaExpr, vars)

		placa, ok := placas[regra.MaterialKey]
		if !ok {
			continue // ou usar uma placa padrão
		}

		componentes = append(componentes, models.Componente{
			Nome:          regra.Nome,
			Altura:        altura,
			Largura:       largura,
			Quantidade:    regra.Quantidade,
			TipoUso:       regra.TipoUso,
			Profundidade:  item.Profundidade,
			Espessura:     item.Espessura,
			ItemMontadoID: item.Id,
			PlacaMDFID:    placa.ID,
		})
	}

	return componentes
}
