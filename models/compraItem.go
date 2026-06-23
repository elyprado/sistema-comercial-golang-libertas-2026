package models

// Importação do pacote decimal para lidar com valores monetários com precisão
import "github.com/shopspring/decimal"

type CompraItem struct {
	IdCompraItem  *int             `json:"idcompraitem"`
	IdCompra      *int             `json:"idcompra"`
	IdProduto     *int             `json:"idproduto"`
	Quantidade    *int             `json:"quantidade"`
	CustoUnitario *decimal.Decimal `json:"custo_unitario"`
	CustoTotal    *decimal.Decimal `json:"custo_total"`
}
