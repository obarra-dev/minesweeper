BINARY=minesweeper

test_local:
	go test -cover ./...

test:
	go test -json > report.json -cover -coverprofile=coverage.out -race ./...

format:
	gofmt -s -w .

check_format:
	gofmt -d .

vet:
	go vet ./...

build:
	go build -o ${BINARY} ./*.go
