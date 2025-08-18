package models

import "gorm.io/gorm"

type Filete struct {
	gorm.Model
	ID         uint    `gorm:"primaryKey" json:"id"`
	Cor        string  `gorm:"not null" json:"cor"`
	Largura    float64 `gorm:"not null" json:"largura"`     // mm
	PrecoMetro float64 `gorm:"not null" json:"preco_metro"` // R$ por metro
}
