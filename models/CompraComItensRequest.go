package models

type CompraComItensRequest struct {
	Data           *string      `json:"data"`
	IdFornecedor   *int         `json:"idfornecedor"`
	DataVencimento *string      `json:"data_vencimento"`
	Itens          []CompraItem `json:"itens"`
}
