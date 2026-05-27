package models

type Produto struct {
	IdProduto         *int     `json:"idproduto"`
	Descricao         *string  `json:"descricao"`
	Marca             *string  `json:"marca"`
	PrecoDeCusto      *float64 `json:"preco_de_custo"`
	PrecoDeVenda      *float64 `json:"preco_de_venda"`
	CodigoReferencia  *string  `json:"codigo_referencia"`
	CodigoBarras      *string  `json:"codigo_barras"`
}