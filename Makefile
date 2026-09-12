DIST := dist
PLATFORMS := \
	linux/amd64 \
	linux/arm64 \
	darwin/amd64 \
	darwin/arm64 \
	windows/amd64 \
	windows/arm64

# cross_compile,<binary-name>,<package-path>
define cross_compile
	@for platform in $(PLATFORMS); do \
		os=$${platform%/*}; arch=$${platform#*/}; \
		out=$(DIST)/$(1)_$${os}_$${arch}; \
		if [ "$${os}" = "windows" ]; then out=$${out}.exe; fi; \
		echo "building $${out}"; \
		CGO_ENABLED=0 GOOS=$${os} GOARCH=$${arch} \
			go build -trimpath -ldflags "-s -w" -o $${out} $(2) || exit 1; \
	done
endef

test:
	go test -v -race -count 1 ./...

# Cross-compile the CLI and the MCP server for every platform in PLATFORMS into ./dist.
build-all: build-all-cli build-all-mcp-server

build-all-cli:
	mkdir -p $(DIST)
	$(call cross_compile,scales,.)

build-all-mcp-server:
	mkdir -p $(DIST)
	$(call cross_compile,scales-mcp-server,./internal/mcpserver)


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

generate-docs:
	go run ./internal/gendocs

mcp-server:
	go run ./internal/mcpserver
