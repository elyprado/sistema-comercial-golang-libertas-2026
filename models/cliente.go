package models

type Cliente struct {
	Idcliente  *int    `json:"idcliente"`
	Nome       *string `json:"nome"`
	Telefone   *string `json:"telefone"`
	Cpf        *string `json:"cpf"`
	Email      *string `json:"email"`
	Logradouro *string `json:"logradouro"`
	Numero     *string `json:"numero"`
	Bairro     *string `json:"bairro"`
	Cidade     *string `json:"cidade"`
	Uf         *string `json:"uf"`
}
