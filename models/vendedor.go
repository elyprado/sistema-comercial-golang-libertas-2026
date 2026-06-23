package models

type Vendedor struct {
    Idvendedor         *int     `json:"idvendedor"`
    Nome               *string  `json:"nome"`
    Telefone           *string  `json:"telefone"`
    Cpf                *string  `json:"cpf"`
    Email              *string  `json:"email"`
    ComissaoPercentual *float64 `json:"comissao_percentual"`
    Salario            *float64 `json:"salario"`
}