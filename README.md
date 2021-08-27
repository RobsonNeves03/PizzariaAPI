# PizzariaAPI
Teste de Golang + HTML + MySQL

Comandos para criar o banco de dados
CREATE DATABASE pizzaria;
CREATE TABLE pizzas(id int auto_increment not null primary key, nome varchar(30) not null, preco int not null);
INSERT INTO pizzas (nome, preco) VALUES ("mussarella", 20);
INSERT INTO pizzas (nome, preco) VALUES ("calabresa", 30);
INSERT INTO pizzas (nome, preco) VALUES ("frango c/ catupiry", 35);
INSERT INTO pizzas (nome, preco) VALUES ("portuguesa", 30);
INSERT INTO pizzas (nome, preco) VALUES ("4 queijos", 25);
INSERT INTO pizzas (nome, preco) VALUES ("baiana", 36);
INSERT INTO pizzas (nome, preco) VALUES ("bacon", 30);
INSERT INTO pizzas (nome, preco) VALUES ("pepperoni", 38);
