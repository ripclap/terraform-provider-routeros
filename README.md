# Terraform Provider RouterOS

[![CI](https://github.com/ripclap/terraform-provider-routeros/actions/workflows/ci.yml/badge.svg?branch=full-device-coverage)](https://github.com/ripclap/terraform-provider-routeros/actions/workflows/ci.yml)
[![CodeQL](https://github.com/ripclap/terraform-provider-routeros/actions/workflows/codeql.yml/badge.svg?branch=full-device-coverage)](https://github.com/ripclap/terraform-provider-routeros/actions/workflows/codeql.yml)
[![OpenTofu Registry](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fregistry.opentofu.org%2Fv1%2Fproviders%2Fripclap%2Frouteros%2Fversions&query=%24.versions%5B0%5D.version&label=opentofu&color=blue)](https://search.opentofu.org/provider/ripclap/routeros/latest)
[![Latest release](https://img.shields.io/github/v/release/ripclap/terraform-provider-routeros?label=release)](https://github.com/ripclap/terraform-provider-routeros/releases)

Manage MikroTik RouterOS 7 configuration with OpenTofu, over the REST API or the
binary API, with or without TLS.

> **This is a fork** of
> [terraform-routeros/terraform-provider-routeros](https://github.com/terraform-routeros/terraform-provider-routeros),
> licensed under MPL-2.0 and not affiliated with or endorsed by the upstream
> project. See [FORK.md](FORK.md) for what differs and [NOTICE](NOTICE) for
> attribution.

It registers **415 resources** covering 388 RouterOS menus, against 254
upstream, so a whole device can be held in state and drift detected across all
of it rather than a subset. [COVERAGE.md](COVERAGE.md) lists what it adds.

Every list menu also has a **data source** of the same name — 308 in total —
returning the menu's entries under `entries`, narrowed by an optional `filter`:

```terraform
data "routeros_interface_bridge_port" "ports" {
  filter = { bridge = "bridge" }
}

output "member_interfaces" {
  value = data.routeros_interface_bridge_port.ports.entries[*].interface
}
```

## Using the provider

Enable the REST API on the router first: create a certificate under
`/certificate` and enable the `www-ssl` service under `/ip/service` using it.
[MikroTik's documentation](https://help.mikrotik.com/docs/display/ROS/REST+API)
covers this.

```terraform
terraform {
  required_providers {
    routeros = {
      source  = "ripclap/routeros"
      version = "~> 2.0"
    }
  }
}

provider "routeros" {
  hosturl  = "https://my.router.local"
  username = "my_username"
  password = "my_super_secret_password"
}
```

`hosturl` accepts `https://` and `http://` for the REST API, and `apis://`
(8729) or `api://` (8728) for the binary API. The binary API is faster and is
the one that completes a read against a very large routing table; see
[FORK.md](FORK.md#transports).

Published to the [OpenTofu Registry](https://search.opentofu.org/provider/ripclap/routeros/latest),
where the resource and data source reference lives. It is not published to the
HashiCorp Terraform Registry.

### Compatibility

Tested against RouterOS 7.23.x, built with Go 1.25. Compatibility is only
claimed within RouterOS 7.

### Version notes

- **2.0.2 is withdrawn and will not install.** Use 2.0.3 or later.
- Upstream 1.43 changed the schemas of `routeros_routing_bgp_connection`,
  `routeros_ipv6_neighbor_discovery` and `routeros_interface_wireguard_peer`.
  For the first two, remove the resource from state and import it again.

## Verifying a release

Each release ships a GPG-signed checksum file, a CycloneDX SBOM per archive, and
a build provenance attestation:

```bash
gh attestation verify terraform-provider-routeros_2.0.3_linux_amd64.zip \
  --repo ripclap/terraform-provider-routeros
```

## Documentation

| | |
|---|---|
| [FORK.md](FORK.md) | What differs from upstream, and transport behaviour |
| [COVERAGE.md](COVERAGE.md) | Resources this fork adds |
| [CHANGELOG.md](CHANGELOG.md) | Release history |
| [CONTRIBUTING.md](CONTRIBUTING.md) | Building, testing and adding resources |
| [RELEASING.md](RELEASING.md) | How a release is cut, and why a published version is never rebuilt |

_Not affiliated with MikroTik. MikroTik and RouterOS are trademarks of
Mikrotikls SIA._
