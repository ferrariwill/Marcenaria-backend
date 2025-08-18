package models

type ItemMontado struct {
	Id           uint    `json:"id" gorm:"primaryKey"`
	Nome         string  `json:"nome" gorm:"not null"`
	Largura      float64 `json:"largura" gorm:"not null"`
	Altura       float64 `json:"altura" gorm:"not null"`
	Profundidade float64 `json:"profundidade" gorm:"not null"`
	Espessura    float64 `json:"espessura" gorm:"not null"`

	ModeloId    uint         `json:"modelo_id" gorm:"not null"`
	ProjetoId   uint         `json:"projeto_id" gorm:"not null"`
	Modelo      ModeloMovel  `json:"modelo" gorm:"foreignKey:ModeloId"`
	Projeto     Projeto      `json:"projeto" gorm:"foreignKey:ProjetoId"`
	Componentes []Componente `json:"Componente"`
}
