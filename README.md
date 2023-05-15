# Desafio PISMO

Esta é uma aplicação que faz parte do processo admissional da **PISMO**, nela iremos realizar a criação da conta do cliente e realizar transações na sua conta
Nela estaremos aplicando a arquitetura hexagonal proposta pelo Alistair Cockburn.
https://alistair.cockburn.us/hexagonal-architecture

# Como rodar o projeto?
**Antes de tudo, garanta que na sua máquina tenha o docker instalado.**

Segue as variaveis ambiente para rodar o projeto:
```
PG_HOST=
PG_PORT=
PG_USER=
PG_PASSWORD=
PG_DATABASE=

```

Para rodar o projeto basta executar o seguinte comando:
```bash
make run-dev
```
Caso deseje derrubar os conteineres levantandos pelo comando acima, basta executar o seguinte comando:
```bash
make stop-dev
```

# Como rodar os testes?
Para rodar os testes basta executar o seguinte comando:
```bash
make test
```

# Desenho da arquitetura

![image info](./assets/architecture.png)

# Tecnologias utilizadas
 - Banco de dados postgres
 - Golang
 - Docker

## Dependencias do Golang
- github.com/asaskevich/govalidator v0.0.0-20230301143203-a9d515a09cc2
- github.com/gin-gonic/gin v1.9.0
- github.com/golang/mock v1.6.0
- github.com/lib/pq v1.10.9
- github.com/mattn/go-sqlite3 v1.14.16
- github.com/stretchr/testify v1.8.2

# Distribuição das pastas
```
.
├── Dockerfile
├── Makefile
├── README.md
├── assets
│   ├── architecture.drawio
│   └── architecture.png
├── docker-compose.yml
├── go.mod
├── go.sum
├── migrations
│   ├── 20230515180138_create_account_table.sql
│   ├── 20230515180516_create_transaction_table.sql
│   └── 20230515180558_create_transaction_trigger.sql
└── src
    ├── adapters
    │   ├── database
    │   │   ├── account
    │   │   │   ├── account.go
    │   │   │   └── account_test.go
    │   │   └── transaction
    │   │       ├── transaction.go
    │   │       └── transaction_test.go
    │   └── http
    │       ├── account
    │       │   ├── api.go
    │       │   └── api_test.go
    │       └── transaction
    │           ├── api.go
    │           └── api_test.go
    ├── application
    │   ├── dtos
    │   │   ├── account.go
    │   │   └── transaction.go
    │   ├── mocks
    │   │   ├── models
    │   │   │   ├── account.go
    │   │   │   └── transaction.go
    │   │   ├── ports
    │   │   │   ├── account.go
    │   │   │   └── transaction.go
    │   │   └── services
    │   │       ├── account.go
    │   │       └── transaction.go
    │   ├── models
    │   │   ├── account.go
    │   │   └── transaction.go
    │   ├── ports
    │   │   ├── account.go
    │   │   └── transaction.go
    │   └── services
    │       ├── account.go
    │       ├── account_test.go
    │       ├── transaction.go
    │       └── transaction_test.go
    └── cmd
        └── api
            └── main.go
```

# Ideias para melhoria da aplicação
- Criar documentação da API usando OpenAPI Redoc
- Usar o NGINX como load balancer
- Implementar o pattern CQRS para separar as responsabilidades de leitura e escrita usando redis para cache
- Criar um serviço que realize de maneira mais efetiva o pagamento parcelado
- Criar um serviço de notificação para o cliente
- Criar o pattern saga para garantir a consistência das transações
  1. Criar serviço de compensação para caso ocorra falha na transação
- Melhorar as estruturas de dados que comportam as informações de amount da transação
  1. Retirar o ponto flutuante e utilizar inteiros não negativos (uint64) para maior precisão
  2. Utilizar o pattern decimal para maior precisão
- Realizar deploy em nuvem publica usando IaaC (Serverless Framework ou AWS SAM)
