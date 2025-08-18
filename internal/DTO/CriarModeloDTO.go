package dto

type CriarModeloDTO struct {
	Nome      string     `json:"nome"`
	Descricao string     `json:"descricao"`
	Regras    []RegraDTO `json:"regras"`
}
