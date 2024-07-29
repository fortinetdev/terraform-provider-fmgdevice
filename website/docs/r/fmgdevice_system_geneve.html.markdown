---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_geneve"
description: |-
  Configure GENEVE devices.
---

# fmgdevice_system_geneve
Configure GENEVE devices.

## Example Usage

```hcl
resource "fmgdevice_system_geneve" "trname" {
  dstport     = 10
  interface   = ["port2"]
  ip_version  = "ipv6-unicast"
  name        = "your own value"
  remote_ip   = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `dstport` - GENEVE destination port (1 - 65535, default = 6081).
* `interface` - Outgoing interface for GENEVE encapsulated traffic.
* `ip_version` - IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast. Valid values: `ipv4-unicast`, `ipv6-unicast`.

* `name` - GENEVE device or interface name. Must be an unique interface name.
* `remote_ip` - IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
* `remote_ip6` - IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
* `type` - GENEVE type. Valid values: `ethernet`, `ppp`.

* `vni` - GENEVE network ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Geneve can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_geneve.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

