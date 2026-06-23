package models

import "time"

type ContasPagar struct {
	IDContasPagar  int       `json:"id_contas_pagar"`
	IDFornecedor   int       `json:"id_forncedor"`
	Data           time.Time `json:"data"`
	DataVencimento time.Time `json:"data_vencimento"`
	Valor          float32   `json:"valor"`
	DataPagamento  time.Time `json:"data_pagamento"`
}

type ContasPagarRequest struct {
	IDFornecedor   int       `json:"id_forncedor"`
	Data           time.Time `json:"data"`
	DataVencimento time.Time `json:"data_vencimento"`
	Valor          float32   `json:"valor"`
}

type ContasPagarResponse struct {
	IDContasPagar  int        `json:"id_contas_pagar"`
	IDFornecedor   int        `json:"id_forncedor"`
	Data           *time.Time `json:"data"`
	DataVencimento *time.Time `json:"data_vencimento"`
	Valor          float32    `json:"valor"`
	DataPagamento  *time.Time `json:"data_pagamento"`
}
