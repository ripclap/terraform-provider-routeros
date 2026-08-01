> **This is a fork.** See [FORK.md](FORK.md) for what differs from upstream and [NOTICE](NOTICE) for attribution.

# Terraform Provider RouterOS

[![CI](https://github.com/ripclap/terraform-provider-routeros/actions/workflows/ci.yml/badge.svg?branch=full-device-coverage)](https://github.com/ripclap/terraform-provider-routeros/actions/workflows/ci.yml)
[![OpenTofu Registry](https://img.shields.io/badge/dynamic/json?url=https%3A%2F%2Fregistry.opentofu.org%2Fv1%2Fproviders%2Fripclap%2Frouteros%2Fversions&query=%24.versions%5B0%5D.version&label=opentofu&color=blue)](https://search.opentofu.org/provider/ripclap/routeros/latest)
[![Latest release](https://img.shields.io/github/v/release/ripclap/terraform-provider-routeros?label=release)](https://github.com/ripclap/terraform-provider-routeros/releases)


**Note**: Version 2.0.2 is withdrawn and will not install. Use 2.0.3 or later.

**Note**: In upstream release 1.43, the resource schemas have been changed:
* `routeros_routing_bgp_connection`
* `routeros_ipv6_neighbor_discovery`
* `routeros_interface_wireguard_peer`

For the first two to work correctly, you must remove the resource state (`terraform state rm <name>`) and import it again (`terraform import [options] <name> <id>`).

## Purpose

This provider allows you to configure Mikrotik routers using [old API](https://help.mikrotik.com/docs/display/ROS/API) or [REST API](https://help.mikrotik.com/docs/display/ROS/REST+API), using or not using TLS.
Compatibility testing is only performed within ROS version 7.x.

From version 1.0.0, the provider has been rewritten by [vaerh](https://github.com/vaerh), and their [fork](https://github.com/vaerh/terraform-provider-routeros) has now been merged. This version drastically improves adding new endpoints to the provider, enabling significantly easier development. [vaerh](https://github.com/vaerh) has been added as a maintainer to this project.

_We are not affiliated in any way with Mikrotik or the development of RouterOS_
## Using the provider

To get started with the provider, you first need to enable the REST API on your router. [You can follow the Mikrotik documentation on this](https://help.mikrotik.com/docs/display/ROS/REST+API), but the gist is to create an SSL cert (in `/system/certificates`) and enable the `web-ssl` service (in `/ip/services`) which uses that certificate. After that, include the following in your Terraform manifests:

```terraform
terraform {
  required_providers {
    routeros = {
      source = "ripclap/routeros"
    }
  }
}

provider "routeros" {
  hosturl  = "(http|https|api|apis)://my.router.local[:port]"
  username = "my_username"
  password = "my_super_secret_password"
}

```

This fork is published to the OpenTofu Registry; the source address above resolves
under OpenTofu. It is not published to the HashiCorp Terraform Registry.

For more in-depth documentation about each of the resources and datasources, please read the [documentation on the OpenTofu Registry](https://search.opentofu.org/provider/ripclap/routeros/latest)

### Versions tested

- go 1.25 and ROS 7.23.1, 7.23.2 (stable)

## Changelog

For a detailed changelog, please see the [changelog.md](CHANGELOG.md).

## Coverage

[COVERAGE.md](COVERAGE.md) lists the resources this fork adds over upstream.

## Releasing

Tagging builds a draft release; publishing it is a separate, manual step. See
[RELEASING.md](RELEASING.md) — particularly the reason a published version is
never rebuilt.

## Contributing
This version of the module greatly simplifies the process of adding new resources.
You are welcome!

### Testing

You can build the provider locally to test fixes by following these intructions:
- Build and copy the provider where Terraform reads it
```
go build *.go && \
mkdir -p ~/.terraform.d/plugins/terraform.local/local/routeros/1.0.0/$(uname -s | tr '[:upper:]' '[:lower:]')_$(uname -m) && \
mv main ~/.terraform.d/plugins/terraform.local/local/routeros/1.0.0/$(uname -s | tr '[:upper:]' '[:lower:]')_$(uname -m)/terraform-provider-routeros_v1.0.0
```
- Change provider from 
```hcl
required_providers {
  routeros = {
    source  = "ripclap/routeros"
    version = "2.0.3"
  }
}
```

to
```hcl
required_providers {
  routeros = {
    source  = "terraform.local/local/routeros"
    version = "1.0.0"
  }
}
```
- Clean your providers, init and apply
- Alternatively, you can edit/create ~/.terraformrc add add a provider installation block like:
```hcl
provider_installation {
  dev_overrides {
     "ripclap/routeros" = "/path/to/your/git/clone"
  }

  direct {
  }
}
```

and then build the provider using
```
go build -o terraform-provider-routeros *.go
```
in order for Terraform to find it.

### Fixing RouterOS property drift

Sometimes RouterOS might introduce a breaking change on a property. You can easilfy contribute to the provider by following these intructions:

- Edit `routeros/mikrotik_resource_drift.yaml`. Add the resource used as well as the old property name and the new one
- Perform the generator. It should edit file `routeros/mikrotik_resource_drift.go`.
```bash
cd routeros/
go run ../tools/drift/main.go
```
- Submit your changes!

[Here](https://github.com/terraform-routeros/terraform-provider-routeros/pull/758/files) is a example of pull request.
