package dto

type OrcamentoDTO struct {
	ModeloID     uint
	Largura      float64
	Altura       float64
	Profundidade float64
	Espessura    float64
	Materiais    map[string]uint
	Ferragens    []uint // IDs das ferragens usadas
	Filetes      []uint // IDs dos filetes usados
	CustoExtras  float64
	MargemLucro  float64
	Descricao    string
	ClienteID    uint
}
