-- Schema para o módulo de Vendas + VendaItens

CREATE TABLE IF NOT EXISTS venda (
  idvenda INT AUTO_INCREMENT PRIMARY KEY,
  data DATETIME NOT NULL,
  idcliente INT NOT NULL,
  idvendedor INT NOT NULL,
  data_vencimento DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS vendaitem (
  idvendaitem INT AUTO_INCREMENT PRIMARY KEY,
  idvenda INT NOT NULL,
  idproduto INT NOT NULL,
  quantidade INT NOT NULL,
  valor_unitario DECIMAL(10,2) NOT NULL,
  valor_total DECIMAL(10,2) NOT NULL,
  CONSTRAINT fk_vendaitem_venda FOREIGN KEY (idvenda) REFERENCES venda(idvenda)
);
