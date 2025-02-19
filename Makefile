test:
	go test -v -race -count 1 ./...

fmt:
	gofumpt -w .
