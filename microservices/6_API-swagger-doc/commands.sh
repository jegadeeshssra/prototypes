choco install make
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
swagger generate spec -o ./swagger.yaml --scan-models
make swagger