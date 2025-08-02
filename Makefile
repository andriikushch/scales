test:
	go test -v -race -count 1 ./...


generate-supported-chord-qualities:
	mkdir -p ./build
	SCALES_TEST_COLLECT_CHORD_QUALITIES=true go test -v -race -count 1 -run TestParseChord ./...
	cat build/output.txt | sort | uniq > build/uniq.txt
	rm -f build/output.txt
	go run ./internal/tool/main.go build/uniq.txt
	rm -f build/uniq.txt

fmt:
	gofumpt -w .
