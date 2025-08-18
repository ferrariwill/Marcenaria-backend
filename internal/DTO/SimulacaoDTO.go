package dto

type SimulacaoDTO struct {
	ModeloID     uint            `json:"modelo_id"`
	Largura      float64         `json:"largura"`
	Altura       float64         `json:"altura"`
	Profundidade float64         `json:"profundidade"`
	Espessura    float64         `json:"espessura"`
	Materiais    map[string]uint `json:"materiais"` // MaterialKey → PlacaMDFID
}
