go get github.com/go-openapi/runtime
# downloads required modules and removes unused ones
go mod tidy

go test -v

choco install make
go install github.com/go-swagger/go-swagger/cmd/swagger@latest
swagger generate spec -o ./swagger.yaml --scan-models
make swagger

