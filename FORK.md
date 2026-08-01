# About this fork

A fork of [terraform-provider-routeros](https://github.com/terraform-routeros/terraform-provider-routeros),
licensed under MPL-2.0. See [NOTICE](NOTICE) for attribution and [LICENSE.md](LICENSE.md)
for the license text.

## Purpose

To represent a complete RouterOS 7.23 configuration as OpenTofu state, so that drift is
detectable across the whole device rather than a subset of it.

| | upstream v1.99.1 | this fork |
|---|---|---|
| Registered resources | 255 | **415** |
| Distinct RouterOS menus covered | 207 | **388** |

Resources outnumber menus because some menus are reachable under more than one resource
name.

## What differs

1. **Additional resources** for menus upstream does not model.
2. **Additional schema fields** that RouterOS 7.23 returns but upstream does not declare.
3. **Relaxed validation** where upstream was stricter than the device.

`routeros_ip_cloud.ddns_enabled` is a boolean rather than a string, and
`routeros_routing_bgp_vpn.instance` is required. See [CHANGELOG.md](CHANGELOG.md).

## Transports

The provider speaks REST (`https://`) and the binary API (`api://` port 8728,
`apis://` port 8729). They are not equivalent in practice:

- A single route read from a very large routing table can exceed the REST timeout,
  because each request scans the table. The same read completes over `apis://`.
- The binary API is measurably faster on ordinary reads.

If a resource times out on import, try `apis://` before assuming it is unsupported.

## Importing a whole device

Bulk import is safe at the default parallelism except against unbounded tables. Exclude
routing tables from bulk imports, or read them over `apis://`.

## Notes

- `routeros_ip_service` imports by service name rather than by `.id`.
- Menus whose entries the router derives, and menus that only report state, are not
  modelled as resources.
- Resources for menus that ship in separate packages (`caps-man`, `iot`, `lte`,
  `user-manager`, `zerotier`, `openflow`, wireless) are present but exercised less.
- `gps`, `tr069-client` and `dude` are not modelled.

## Acceptance tests

The suite requires a device and must be run with Terraform; OpenTofu rejects the SDK test
framework's provider address.

```bash
export ROS_HOSTURL="https://<device>" ROS_USERNAME=<user> ROS_PASSWORD=<pass>
export ROS_INSECURE=true ROS_VERSION="7.23.1"
export TF_ACC=1 TF_ACC_TERRAFORM_PATH="$(command -v terraform)"
go test ./routeros/ -count=1 -v -timeout 180m
```

Run tests individually with a per-test timeout and clean up created objects between them.
The device needs an interface named `bridge` and two spare ethernet or veth interfaces.

Resources that manage disks, packages or the watchdog are not covered by the suite.
