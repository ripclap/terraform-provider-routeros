VERSION := $(shell git describe --tags --abbrev=0)

EXT :=
ifeq ($(OS),Windows_NT)
	EXT := .exe
endif

.PHONY: build test lint docs generate tfformat debug snapshot

build:
	go build ./...

# Mirrors the unit-test step in .github/workflows/ci.yml. The acceptance suite
# needs a device; see FORK.md.
test:
	ROS_VERSION=7.23.1 go test ./routeros/ -count=1 -skip 'TestClientTransport_SendRequest'

lint:
	golangci-lint run ./...

# Regenerates everything CI checks for staleness.
generate:
	cd routeros && go run ../tools/drift/main.go
	go run tools/coverage/main.go
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name routeros

docs:
	go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name routeros

tfformat:
	terraform fmt -recursive examples/

debug:
	go build -gcflags="all=-N -l" -o terraform-provider-routeros_$(VERSION)$(EXT) .

# Release builds run in CI only; this produces the same artifacts locally
# without publishing or signing them.
snapshot:
	goreleaser release --snapshot --clean --skip=publish,sign,sbom
