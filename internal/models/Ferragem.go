package models

import "gorm.io/gorm"

type Ferragem struct {
	gorm.Model
	ID            uint    `gorm:"primaryKey" json:"id"`
	Nome          string  `gorm:"not null" json:"nome"`           // Ex: Corrediça 35mm
	Tipo          string  `gorm:"not null" json:"tipo"`           // Ex: Corrediça, Puxador
	PrecoUnitario float64 `gorm:"not null" json:"preco_unitario"` // R$ por unidade
}
