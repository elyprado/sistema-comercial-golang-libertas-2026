package models

type Compras struct {
	IdCompra       *int    `json:"idcompra"`
	Data           *string `json:"data"`
	IdFornecedor   *int    `json:"idfornecedor"`
	DataVencimento *string `json:"datavencimento"`
}
