.PHONY: deps test build

deps:
	go mod tidy -v
	go get -u ./...

build:
	go build -o layer ./main.go

test:
	go build -o clean-layer-lint ./main.go
	./clean-layer-lint ./...
