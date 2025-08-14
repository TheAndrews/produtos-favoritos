# Api to manage customers and wishlists

Api made with Go connecting to a postgres db

This project utilizes concept of DDD and solid principles<br>
It utilizes uber-go-dig as a injection container, from repositories to services and controllers all are implementing depending on interfaces which are fulfilled by the container.
There's a simple authentication that checks for an api key (header x-api-key) <br>
Tests are present together with their implementation following Go test pattern, they make use of mocks generated from interfaces (mockery package) <br>
At the repository level tests are Integration tests, they hit the database using a docker postgres pod (testcontainers-go package). <br>

Go version 1.24 <br>
Packages utilized <br>

- gin
- gin-swagger
- dig
- gorm
- gormigrate
- mockery
- testcontainers-go

Project Structure

- src
  - api
    - controllers
    - di container
    - swagger docs
    - forms
    - middlewares (auth)
    - router
  - domain
    - interfaces
    - models
    - services
  - infrastructure
    - config
    - database
      - migrations
      - repositories
  - internals
    - exceptions
    - mocks <br>

# Running manually

go mod tidy <br>
go run .\src\main.go

# Running in docker compose

api & db

docker-compose up --build

# Swagger Url

http://localhost:8080/docs/index.html

# Swagger installation

go install github.com/swaggo/swag/cmd/swag@latest
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
swag init --dir ./src --output ./src/api/docs

# Mockery Installation

go install github.com/vektra/mockery/v2@latest
mockery --all --output=./src/internals/mocks

# To help debug locally

go install github.com/go-delve/delve/cmd/dlv@latest

quickly showcase video:
https://vimeo.com/1108470413
