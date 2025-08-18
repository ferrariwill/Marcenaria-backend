package models

type RegraModelo struct {
	ID            uint   `json:"id" gorm:"primaryKey"`
	Nome          string `json:"nome"`
	LarguraExpr   string `json:"largura_expr"`
	AlturaExpr    string `json:"altura_expr"`
	Quantidade    int    `json:"quantidade"`
	TipoUso       string `json:"tipo_uso"`
	MaterialKey   string `json:"material_key"` // Ex: "branco", "freijo", "cinza"
	ModeloMovelID uint   `json:"modelo_movel_id"`
}
