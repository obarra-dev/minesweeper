BINARY=minesweeper

test_local:
	go test -cover ./...

test:
	go test -json > report.json -cover -coverprofile=coverage.out -race ./...

build:
	go build -o ${BINARY} ./*.go
