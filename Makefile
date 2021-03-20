BINARY=minesweeper

test:
	go test -v -cover ./...

report:
	go test -json > report.json -cover -coverprofile=coverage.out -race ./...

format:
	gofmt -s -w .

check_format:
	gofmt -d .

go_lint:
	golint ./...

vet:
	go vet ./...

# Run all code checks.
lint:  check_format go_lint vet

build:
	go build -o ${BINARY} ./*.go

