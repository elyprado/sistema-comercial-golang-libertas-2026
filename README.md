# Projeto - Gestao de Loja

Projeto desenvolvido pelos alunos do 5o periodo do curso de Sistemas de Informacao da Libertas Faculdades Integradas.

## Informacoes da disciplina

- Professor: Ely Fernando do Prado
- Disciplina: Desenvolvimento de Sistemas Web II
- Valor: 50 pontos
- Data de entrega: 19/06/2026
- Formato de entrega: commit no repositorio do GitHub

## Tecnologias

- Backend: Go (Golang)
- Frontend: HTML, CSS e JavaScript

Desenvolver a aplicacao conforme o exemplo apresentado em aula.

## Requisitos do projeto

### 1. Produto

**Escopo:** permitir o cadastro de produtos com todas as operacoes CRUD (Create, Retrieve, Update, Delete).

**Atributos:**
- idproduto
- descricao
- marca
- preco_de_custo
- preco_de_venda
- codigo_referencia
- codigo_barras

### 2. Cliente

**Escopo:** permitir o cadastro de clientes com todas as operacoes CRUD.

**Atributos:**
- idcliente
- nome
- telefone
- cpf
- e-mail
- logradouro
- numero
- bairro
- cidade
- uf

### 3. Fornecedor

**Escopo:** permitir o cadastro de fornecedores com todas as operacoes CRUD.

**Atributos:**
- idfornecedor
- nome
- telefone
- cnpj
- e-mail
- logradouro
- numero
- bairro
- cidade
- uf

### 4. Vendedor

**Escopo:** permitir o cadastro de vendedores com todas as operacoes CRUD.

**Atributos:**
- idvendedor
- nome
- telefone
- cpf
- e-mail
- percentual_comissao
- salario

### 5. Vendas + VendaItens

**Escopo:** permitir o registro de vendas em uma interface de PDV (Ponto de Venda), na qual o usuario deve selecionar o cliente, o vendedor, a data de vencimento e adicionar em uma lista os produtos e suas respectivas quantidades.

Tambem deve existir uma interface para consulta de vendas, com filtro por data inicial e data final, permitindo detalhar os itens da venda.

Implementar:
- servico de insercao que recebe os dados da venda e o array de itens;
- servico de consulta que retorna a lista de vendas;
- servico `getById` que retorna todos os dados da venda, incluindo seus itens.

No frontend, para os campos `select` de cliente, vendedor e produto, devem ser consumidas as APIs dos respectivos cadastros.

**Atributos:**
- venda: idvenda, data, idcliente, idvendedor, data_vencimento
- vendaitem: idvendaitem, idvenda, idproduto, quantidade, valor_unitario, valor_total

### 6. Compras + CompraItens

**Escopo:** permitir o registro de compras em uma interface na qual o usuario deve selecionar o fornecedor, a data de vencimento e adicionar em uma lista os produtos, as quantidades e o preco de custo de cada item.

Tambem deve existir uma interface para consulta de compras, com filtro por data inicial e data final, permitindo detalhar os itens da compra.

Implementar:
- servico de insercao que recebe os dados da compra e o array de itens;
- servico de consulta que retorna a lista de compras;
- servico `getById` que retorna todos os dados da compra, incluindo seus itens.

No frontend, para os campos `select` de fornecedor e produto, devem ser consumidas as APIs dos respectivos cadastros.

**Atributos:**
- compra: idcompra, data, idfornecedor, data_vencimento
- compraitem: idcompraitem, idcompra, idproduto, quantidade, custo_unitario, custo_total

### 7. Contas a Receber

**Escopo:** criar uma interface para listar os registros de contas a receber, com filtro por data de vencimento. Nessa mesma interface, o usuario deve poder inserir uma nova conta a receber e baixar um item das contas a receber.

Exibir em cor diferente os itens que ja foram pagos. Ao clicar no botao de baixar, o sistema deve salvar o atributo `data_pagamento`.

**Atributos:**
- idcontaareceber
- idcliente
- data
- data_vencimento
- valor
- data_pagamento

### 8. Contas a Pagar

**Escopo:** criar uma interface para listar os registros de contas a pagar, com filtro por data de vencimento. Nessa mesma interface, o usuario deve poder inserir uma nova conta a pagar e baixar um item das contas a pagar.

Exibir em cor diferente os itens que ja foram pagos. Ao clicar no botao de baixar, o sistema deve salvar o atributo `data_pagamento`.

**Atributos:**
- idcontaapagar
- idfornecedor
- data
- data_vencimento
- valor
- data_pagamento

## Grupos

| Grupo | Aluno 1 | Aluno 2 | Aluno 3 | Tema |
|---|---|---|---|---|
| 1 | OTAVIO AUGUSTO DIZARO CANDIDO | PEDRO AUGUSTO VIEIRA DE SOUZA | KAIO FELIPE DUARTE ZUMERLE | Produto |
| 2 | TULIO MENEZES BESSA | JULIA SUDARIO SILVA | ADRYAN RYAN SANTOS | Cliente |
| 3 | GUSTAVO FERRAREZ GONCALVES | RAFAEL VINICIUS DOS SANTOS | LUIZ PAULLO CAETANO | Fornecedor |
| 4 | RENAN MORAES DOS SANTOS | NATHAN ANTONIO VALVASSOURA DA MATA | VINICIUS AUGUSTO SANTANELI DE ANDRADE | Vendedor |
| 5 | LARISSA ONUZIK SIGIANI | MATEUS ALVES LIMA | THALISON DE OLIVEIRA SANTOS | Vendas + VendaItens |
| 6 | GUSTAVO APARECIDO DE OLIVEIRA CUSTODIO | ULISSES SOARES FILHO | KAUA PEREIRA PAIM | Compras + CompraItens |
| 7 | JONATHAN MIGUEL DE ALMEIDA PAIVA | EDUARDO JOSE ATAIR | GUILHERME DO NASCIMENTO SOUZA | Contas a Receber |
| 8 | CAIO DA SILVA MELO | FELIPE DE SOUZA ALVES MATHEUS | GABRIEL CRISTIANO FERRAREZ | Contas a Pagar |
