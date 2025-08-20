package models

import (
	"time"
)

type Orcamento struct {
	Id             uint         `json:"id" gorm:"primaryKey"`
	Descricao      string       `json:"descricao"`
	Componentes    []Componente `json:"componentes" gorm:"foreignKey:OrcamentoID"`
	CustoMateriais float64      `json:"custo_materiais"`
	CustoFerragens float64      `json:"custo_ferragens"`
	CustoFiletes   float64      `json:"custo_filetes"`
	CustosExtras   float64      `json:"custos_extras"`
	MargemLucro    float64      `json:"margem_lucro"`
	PrecoFinal     float64      `json:"preco_final"`
	DataCriacao    time.Time    `json:"data_criacao" gorm:"autoCreateTime""`
}
