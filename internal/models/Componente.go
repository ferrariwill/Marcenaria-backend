package models

type Componente struct {
	Id            uint        `json:"id" gorm:"primaryKey"`
	Nome          string      `json:"nome" gorm:"not null"`
	Largura       float64     `json:"largura"` // em mm
	Altura        float64     `json:"altura"`  // em mm
	Profundidade  float64     `json:"profundidade"`
	Espessura     float64     `json:"espessura"` // em mm
	Quantidade    int         `json:"quantidade" gorm:"default:1"`
	TipoMaterial  string      `json:"tipo_material"` // Ex: "MDF"
	TipoUso       string      `json:"tipo_uso"`      // Ex: "gaveta", "estrutura"
	PlacaMDFID    uint        `json:"placa_mdf_id"`  // FK para a placa usada
	PlacaMDF      PlacaMDF    `json:"placa_mdf" gorm:"foreignKey:PlacaMDFID"`
	ItemMontadoID uint        `json:"item_montado_id"`
	ItemMontado   ItemMontado `json:"item_montado" gorm:"foreignKey:ItemMontadoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
