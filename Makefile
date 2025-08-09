test:
	go test -v -race -count 1 ./...


generate-supported-chord-qualities:
	mkdir -p ./build
	SCALES_TEST_COLLECT_CHORD_QUALITIES=true go test -v -race -count 1 -run TestParseChord ./...
	cat build/output.txt | sort | uniq > build/uniq.txt
	rm -f build/output.txt
	go run ./internal/tool/main.go build/uniq.txt pkg/guitar_chord_shapes.go pkg/ukulele_chord_shapes.go pkg/bass_guitar_chord_shapes.go pkg/mandolin_chord_shapes.go chord_shapes.json
	rm -f build/uniq.txt
	gofumpt -w .

fmt:
	gofumpt -w .
