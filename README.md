# generate swagger docs

$ go install github.com/swaggo/swag/cmd/swag@latest
$ swag init --dir ./src --output ./src/api/docs

# to debug locally

$ go install github.com/go-delve/delve/cmd/dlv@latest

# Add mockery generation

$ go install github.com/vektra/mockery/v2@latest
$ mockery --all --output=./src/tests/mocks
