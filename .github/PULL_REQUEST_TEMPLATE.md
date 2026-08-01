## What this changes

<!-- The behaviour that differs, and why. -->

## Device behaviour it depends on

<!-- The RouterOS menu and property involved, and the version you observed it
     against. If the device rejects or renames something, say what it returned. -->

RouterOS version:

## Checklist

- [ ] `make lint` passes
- [ ] `make test` passes
- [ ] `make generate` run, and any resulting changes committed
      (`mikrotik_resource_drift.go`, `COVERAGE.md`, `docs/`)
- [ ] `CHANGELOG.md` updated under `## [Unreleased]`
- [ ] No attribute renamed that mirrors a RouterOS property name
- [ ] No device-specific values (addresses, serial numbers, hostnames,
      captured counters) in schema samples, descriptions or comments
