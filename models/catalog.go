package models

type Client struct {
	Idcliente int    `json:"idcliente"`
	Nome      string `json:"nome"`
}

type Vendor struct {
	Idvendedor int    `json:"idvendedor"`
	Nome       string `json:"nome"`
}

type Product struct {
	Idproduto  int     `json:"idproduto"`
	Descricao  string  `json:"descricao"`
	PrecoVenda float64 `json:"preco_de_venda"`
}
