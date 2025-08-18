package utils

import (
	"math"

	"github.com/Knetic/govaluate"
	"github.com/ferrariwill/marcenaria-backend/internal/models"
)

func CalcularPlacasNecessarias(projeto models.Projeto, placa models.PlacaMDF) int {
	areaPlaca := CalcularAreaPlaca(placa)
	var areaTotal float64

	for _, item := range projeto.Itens {
		for _, c := range item.Componentes {
			if c.TipoMaterial == placa.TipoMaterial {
				area := (c.Largura * c.Altura) / 1000000
				areaTotal += area * float64(c.Quantidade)
			}
		}
	}

	return int(math.Ceil(areaTotal / areaPlaca))

}

func CalcularAreaPlaca(placa models.PlacaMDF) float64 {
	return (placa.Largura * placa.Altura) / 1000000 // mm² → m²
}

func AvaliarExpressao(expr string, variaveis map[string]interface{}) float64 {
	e, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return 0
	}
	resultado, err := e.Evaluate(variaveis)
	if err != nil {
		return 0
	}
	return resultado.(float64)
}
