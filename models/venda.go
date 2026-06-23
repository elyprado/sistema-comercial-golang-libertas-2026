package models

import "time"

type Venda struct {
	Idvenda        int         `json:"idvenda"`
	Data           time.Time   `json:"data"`
	Idcliente      int         `json:"idcliente"`
	Idvendedor     int         `json:"idvendedor"`
	DataVencimento time.Time   `json:"data_vencimento"`
	Total          float64     `json:"total,omitempty"`
	ItemCount      int         `json:"item_count,omitempty"`
	Items          []VendaItem `json:"items,omitempty"`
}

type VendaItem struct {
	Idvendaitem   int     `json:"idvendaitem"`
	Idvenda       int     `json:"idvenda"`
	Idproduto     int     `json:"idproduto"`
	Quantidade    int     `json:"quantidade"`
	ValorUnitario float64 `json:"valor_unitario"`
	ValorTotal    float64 `json:"valor_total"`
}
