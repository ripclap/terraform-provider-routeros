# About this fork

This is a fork of [terraform-provider-routeros](https://github.com/terraform-routeros/terraform-provider-routeros),
licensed under MPL-2.0. See [NOTICE](NOTICE) for attribution and [LICENSE.md](LICENSE.md)
for the license text.

## Why it exists

The goal was to represent the **complete** running configuration of a RouterOS 7.23
device as OpenTofu state, so that configuration drift is detectable. Upstream covers
the resources most people need; it does not cover the whole configuration tree, and
several schemas reject or silently discard values that RouterOS itself accepts.

Measured against a live RouterOS 7.23.2 device:

| | upstream v1.99.1 | this fork |
|---|---|---|
| Registered resources | 255 | **411** |
| RouterOS menus represented | 207 / 547 | **363 / 547** |
| Schema round-trip warnings on a full import | 130 | **93** |

## What changed

Three categories, all detailed in [NOTICE](NOTICE):

1. **New resources** for menus upstream does not model.
2. **Missing schema fields** that RouterOS 7.23 returns but upstream drops on read —
   the value does not round-trip, so drift in it is invisible.
3. **Validation relaxations** where upstream was stricter than the device.

## Testing

Verified against physical MikroTik hardware running RouterOS **7.23.1**,
over both the REST and the binary API transport.

| | count |
|---|---|
| Tests discovered | 352 |
| Passing | **335** |
| Skipped - menu or hardware absent | 17 |
| Failing | **0** |
| Regressions introduced by this fork | **0** |

The zero-regression claim is measured, not asserted: every non-passing test was re-run
against a pristine `v1.99.1` worktree on the same device, and any test that failed on
both is upstream's, not this fork's.

The full-suite run finished with 324 passing, 16 skipped and 12 not passing. All twelve
were then reproduced and resolved:

* **Six** were device-side timeouts under sustained load, or cascades from one. Bridge
  and veth churn on a switch provokes `action timed out (13)`; the tests pass against a
  rested device. Running the suite with a cleanup and a settle step between tests is what
  makes the result reproducible.
* **Three** were an upstream defect. `GetDriftMap` called `log.Fatal` when the RouterOS
  version was unknown, which aborted the entire test binary rather than the one test.
  Confirmed present in upstream `v1.99.1`. An unknown version now applies no
  version-gated renames, so the unit tests no longer need `ROS_VERSION` to be set.
* **Two** were leftover device state. `interface_ethernet` renames a physical port, and
  Terraform cannot undo a rename on destroy, so every later test naming that port failed;
  `tool_sniffer` left a filter set. Both tests now clean up after themselves.
* **One**, `interface_ethernet`, forces a link speed the port must support and then
  asserts the port is running. It now skips when the port has no link.

### Running the acceptance suite

**OpenTofu cannot run these tests.** It rejects the SDK test framework's provider
address outright (`Invalid provider namespace "-"`). Use Terraform:

```bash
export ROS_HOSTURL="https://<device>" ROS_USERNAME=<user> ROS_PASSWORD=<pass>
export ROS_INSECURE=true ROS_VERSION="7.23.1"
export TF_ACC=1 TF_ACC_TERRAFORM_PATH="$(command -v terraform)"
go test ./routeros/ -count=1 -v -timeout 180m
```

Run tests one at a time with a per-test timeout, and clean up created objects between
them. A single hanging test otherwise stalls the suite, and one test's residue shows up
as an unrelated test's failure.

The device also needs a couple of fixtures, or ~10 tests fail for reasons unrelated to
the provider: an interface literally named `bridge`, plus `ether3` and `ether4` (veth
is acceptable on x86).

## Known limitations

- **~5 menus remain unmodelled**: `/routing/isis/interface`,
  `/routing/pimsm/igmp-interface-template`, `/interface/wifi/radio/settings`,
  `/interface/wifi/steering/neighbor-group`, `/ipv6/dhcp-relay/routes`.
- **93 schema warnings remain**, dominated by `managed` (33),
  `client_allowed_address` (31), `vrf` (13) and `l3_hw_offloading` (12).
- **9 resources are deliberately untested**, because exercising them risks the device:
  `disk`, `disk_btrfs_filesystem`, `disk_btrfs_subvolume`, `disk_btrfs_transfer`,
  `partitions`, `system_package_update`, both `system_package_local_update_*`, and
  `system_watchdog`.
- **`routeros_ip_service` import** requires the service *name* as the ID, not `.id`.

## Transport matters more than you would expect

The provider speaks both REST (`https://`) and the binary API (`api://` port 8728,
`apis://` port 8729). They are **not** interchangeable:

- Reading a single static route out of a full BGP table (1.07M routes) **never
  returns within 60s over REST**, because each `GET /rest/ip/route?.id=` scans the
  table. The same read completes in ~72s over `apis://`.
- The binary API is roughly 18% faster on ordinary reads.

If a resource times out on import, try `apis://` before assuming it is unsupported.

## Importing a whole device

Concurrency is safe *except* against unbounded tables. Importing 560 resources at
`-parallelism=10` completes in 39s with zero errors and a transient CPU peak; the
same run including 13 static routes from a full BGP table produced 224 cascading
timeouts. Exclude routes from bulk imports, or read them over `apis://`.

