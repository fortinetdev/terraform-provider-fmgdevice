---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_fabricvpn_advertisedsubnets"
description: |-
  Local advertised subnets.
---

# fmgdevice_system_fabricvpn_advertisedsubnets
Local advertised subnets.

~> This resource is a sub resource for variable `advertised_subnets` of resource `fmgdevice_system_fabricvpn`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_fabricvpn_advertisedsubnets" "trname" {
  access           = "inbound"
  bgp_network      = ["your own value"]
  firewall_address = ["your own value"]
  fosid            = 10
  policies         = ["your own value"]
  device_name      = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `access` - Access policy direction. Valid values: `inbound`, `bidirectional`.

* `bgp_network` - Underlying BGP network.
* `firewall_address` - Underlying firewall address.
* `fosid` - ID.
* `policies` - Underlying policies.
* `prefix` - Network prefix.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System FabricVpnAdvertisedSubnets can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_fabricvpn_advertisedsubnets.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

