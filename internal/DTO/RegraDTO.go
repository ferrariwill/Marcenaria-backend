package dto

type RegraDTO struct {
	Nome        string `json:"nome"`
	LarguraExpr string `json:"largura_expr"`
	AlturaExpr  string `json:"altura_expr"`
	Quantidade  int    `json:"quantidade"`
	TipoUso     string `json:"tipo_uso"`
}
