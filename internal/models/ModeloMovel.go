package models

import "time"

type ModeloMovel struct {
	ID        uint          `gorm:"primaryKey"`
	Nome      string        `gorm:"unique;not null"` // Ex: "Gaveta", "Cama"
	Descricao string        `gorm:"type:text"`
	Regras    []RegraModelo `gorm:"foreignKey:ModeloID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
