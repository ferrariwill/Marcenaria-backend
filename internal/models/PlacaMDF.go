package models

import "gorm.io/gorm"

type PlacaMDF struct {
	gorm.Model
	ID            uint    `gorm:"primaryKey" json:"id"`
	Cor           string  `gorm:"not null" json:"cor"`
	Espessura     float64 `gorm:"not null" json:"espessura"`      // mm
	Altura        float64 `gorm:"default:1850" json:"altura"`     // mm
	Largura       float64 `gorm:"default:2750" json:"largura"`    // mm
	TipoMaterial  string  `gorm:"not null" json:"tipo_material"`  // material
	Fornecedor    string  `gorm:"not null" json:"fornecedor"`     // fornecedor
	PrecoUnitario float64 `gorm:"not null" json:"preco_unitario"` // R$ por placa
}
