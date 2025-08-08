# Api de gerenciamento de clientes e lista de favoritos

Api feita em Go utilizando banco de dados postgres

Esse projeto se utiliza dos conceitos do ddd e solid.
Existe uma autenticação simples checando api key no header (X-Api-Key)
Os testes estão presente junto as suas implementações (padrão go), e fazem uso de mocks auto gerados a partir de interfaces. (pacote mockery)
Nos repositórios os testes são de integração com o banco de dados tomando proveito do pacote (testcontainers-go) que sobe um container e o destroy para os testes.

Versão do go 1.24
Pacotes mais importantes utilizados

- gin
- gin-swagger
- dig
- gorm
- gormigrate
- mockery
- testcontainers-go

A estrutura do projeto consiste em:

- src
  - api
    - controllers
    - container de injeção de dependencia
    - swagger docs
    - forms
    - middlewares (autenticação)
    - router
  - domain
    - interfaces
    - modelos
    - servicos
  - infrastrutura
    - configuracao
    - database
      - migrations
      - repositorios
  - internals - exceções - mocks
    main

# Rodando a aplição manualmente

$ go mod tidy
$ go run .\src\main.go

# Url do swagger

http://localhost:8080/docs/index.html

Para rodar com docker é possível via o docker-compose onde ele sobe um container para a aplicação go e seu banco de dados postgres

$ docker-compose up --build

# Instalação do pacote do swagger

$ go install github.com/swaggo/swag/cmd/swag@latest
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/files
$ swag init --dir ./src --output ./src/api/docs

# Pacote para gerar mocks apartir de interfaces

$ go install github.com/vektra/mockery/v2@latest
$ mockery --all --output=./src/internals/mocks

# Pacote que ajuda a debugar localmente

$ go install github.com/go-delve/delve/cmd/dlv@latest
