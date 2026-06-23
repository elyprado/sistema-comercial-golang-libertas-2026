package models

type Fornecedor struct {
	Idfornecedor *int    `json:"idfornecedor"`
	Nome         *string `json:"nome"`
	Telefone     *string `json:"telefone"`
	Cnpj         *string `json:"cnpj"`
	Email        *string `json:"email"`
	Logradouro   *string `json:"logradouro"`
	Numero       *string `json:"numero"`
	Bairro       *string `json:"bairro"`
	Cidade       *string `json:"cidade"`
	Uf           *string `json:"uf"`
}
