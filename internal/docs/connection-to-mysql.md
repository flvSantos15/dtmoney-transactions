Subindo o banco via docker

```bash
docker compose up
```

Executando o banco no terminal:

```bash
docker compose exec [nome-banco] bash
```

Conectando ao MySQL:

```bash
mysql -u root -p
```

Consultar bases:

```bash
SHOW DATABASES;
```

Selecionar um banco de dados:

```bash
USE [nome_do_banco];
```

Listar tabelas do banco selecionado:

```bash
SHOW TABLES;
```

Ver estrutura de uma tabela:

```bash
DESCRIBE [nome_da_tabela];
```

Criar uma tabela:

```bash
CREATE TABLE transactions (
  id BINARY(16) PRIMARY KEY,
  amount DECIMAL(10,2) NOT NULL,
  amount_type ENUM('withdraw', 'deposit') NOT NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
