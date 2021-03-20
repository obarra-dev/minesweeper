BINARY=minesweeper

test:
	go test -cover ./...

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

build:
	go build -o ${BINARY} ./*.go
