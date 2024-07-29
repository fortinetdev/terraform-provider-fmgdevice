---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_internetserviceextension_entry"
description: |-
  Entries added to the Internet Service extension database.
---

# fmgdevice_firewall_internetserviceextension_entry
Entries added to the Internet Service extension database.

~> This resource is a sub resource for variable `entry` of resource `fmgdevice_firewall_internetserviceextension`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `port_range`: `fmgdevice_firewall_internetserviceextension_entry_portrange`



## Example Usage

```hcl
resource "fmgdevice_firewall_internetserviceextension_entry" "trname" {
  addr_mode = "ipv4"
  dst       = ["your own value"]
  dst6      = ["your own value"]
  fosid     = 10
  port_range {
    end_port   = 10
    fosid      = 10
    start_port = 10
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `internet_service_extension` - Internet Service Extension.

* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

* `dst` - Destination address or address group name.
* `dst6` - Destination address6 or address6 group name.
* `fosid` - Entry ID(1-255).
* `port_range` - Port-Range. The structure of `port_range` block is documented below.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `port_range` block supports:

* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (0 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (0 to 65535).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceExtensionEntry can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "internet_service_extension=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_internetserviceextension_entry.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

