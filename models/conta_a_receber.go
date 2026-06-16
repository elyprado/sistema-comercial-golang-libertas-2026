package models

type ContaAReceber struct {
	Idcontaareceber *int     `json:"idcontaareceber"`
	Idcliente       *int     `json:"idcliente"`
	NomeCliente     *string  `json:"nome_cliente,omitempty"`
	Data            *string  `json:"data"`
	DataVencimento  *string  `json:"data_vencimento"`
	Valor           *float64 `json:"valor"`
	DataPagamento   *string  `json:"data_pagamento"`
}
