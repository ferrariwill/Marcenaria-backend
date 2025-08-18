package models

import "time"

type RegraComponente struct {
	ID          uint   `gorm:"primaryKey"`
	ModeloID    uint   `gorm:"not null"` // FK para ModeloMovel
	Nome        string `gorm:"not null"` // Ex: "Lateral esquerda"
	LarguraExpr string `gorm:"not null"` // Ex: "profundidade"
	AlturaExpr  string `gorm:"not null"` // Ex: "altura - 2*espessura"
	Quantidade  int    `gorm:"not null"`
	TipoUso     string `gorm:"not null"` // Ex: "estrutura", "gaveta"
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
