package models

type Projeto struct {
	Id           uint          `json:"id" gorm:"primary_key"`
	Nome         string        `json:"nome"`
	ClienteID    uint          `json:"cliente_id"`
	Cliente      Cliente       `json:"cliente" gorm:"foreignKey:ClienteID"`
	Itens        []ItemMontado `json:"itens" gorm:"foreignKey:ProjetoID"`
	ValorTotal   float64       `json:"valor_total"`
	PlacasUsadas int           `json:"placas_usadas"`
}
