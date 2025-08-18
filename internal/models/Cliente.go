package models

import "time"

type Cliente struct {
	Id          uint      `json:"id" gorm:"primaryKey"`
	Nome        string    `json:"nome" gorm:"not null"`
	Email       string    `json:"email" gorm:"unique"`
	Telefone    string    `json:"telefone" gorm:"not null"`
	Ativo       bool      `json:"ativo" gorm:"not null"`
	CEP         string    `json:"cep" gorm:"not null"`
	Endereco    string    `json:"endereco" gorm:"not null"`
	Numero      string    `json:"numero" gorm:"not null"`
	Complemento string    `json:"complemento"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	Projetos    []Projeto `json:"projetos" gorm:"foreignKey:ClienteId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
