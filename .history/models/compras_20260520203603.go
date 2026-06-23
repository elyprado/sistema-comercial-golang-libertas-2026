package models

import "time"

type Compras struct {
	IdCompra       *int       `json:"idcompra"`
	Data           *time.Time `json:"data"`
	IdFornecedor   *int       `json:"idfornecedor"`
	DataVencimento *time.Time `json:"datavencimento"`
}
