# Contributing

## Build and test

```bash
make build   # go build ./...
make lint    # golangci-lint, configured by .golangci.yml
make test    # unit tests; no device required
```

`make test` supplies `ROS_VERSION` because the version gates read it before the
`TF_ACC` check and fail outright when it is empty.

The acceptance suite needs a device and is not run in CI; see [FORK.md](FORK.md).

## Generated files, and what CI enforces

Three artefacts are generated and checked for staleness on every push. If you
change a schema, run:

```bash
make generate
```

which regenerates, in order:

| File | From | Generator |
|---|---|---|
| `routeros/mikrotik_resource_drift.go` | `mikrotik_resource_drift.yaml` | `tools/drift` |
| `COVERAGE.md` | the resource map | `tools/coverage` |
| `docs/` | the schemas | `tfplugindocs` |

CI fails if any of them differ from what the sources produce. Editing a
generated file by hand does not survive.

## Running against a local build

Point OpenTofu at your working tree with a dev override in `~/.terraformrc`:

```hcl
provider_installation {
  dev_overrides {
    "ripclap/routeros" = "/path/to/your/clone"
  }
  direct {}
}
```

Then build the provider in that directory:

```bash
go build -o terraform-provider-routeros .
```

With a dev override in place, `init` is not required and the version constraint
is ignored.

## Adding a resource

`tools/boilerplate` generates the resource, its test and its example from a
RouterOS menu. Register the result in `routeros/provider.go`, then run
`make generate`.

Schema conventions worth following, because they are what makes a whole-device
import produce an empty plan:

- Attribute names mirror the RouterOS property, `-` becoming `_`. Where RouterOS
  spells a property unusually, reproduce it exactly rather than correcting it;
  the name is the wire format.
- Optional attributes the device always reports need
  `DiffSuppressFunc: AlwaysPresentNotUserProvided`, or every plan shows a diff
  for a value the user never set.
- Durations use `TimeEqual`, so `1m30s` and `90s` compare equal.
- Do not model one device attribute twice. Two schema fields writing the same
  byte produce a diff that cannot be resolved.

## Fixing RouterOS property drift

When RouterOS renames a property between versions:

1. Add the resource, the old name and the new name to
   `routeros/mikrotik_resource_drift.yaml`.
2. Run `make generate`.
3. Include both in the pull request.

## Pull requests

Keep the change focused, describe what device behaviour it depends on, and note
the RouterOS version you observed it against. Lint and the generated-file checks
must pass; see the checklist in the pull request template.
