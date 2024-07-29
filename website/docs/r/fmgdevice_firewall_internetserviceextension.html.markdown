---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_internetserviceextension"
description: |-
  Configure Internet Services Extension.
---

# fmgdevice_firewall_internetserviceextension
Configure Internet Services Extension.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `disable_entry`: `fmgdevice_firewall_internetserviceextension_disableentry`
>- `entry`: `fmgdevice_firewall_internetserviceextension_entry`



## Example Usage

```hcl
resource "fmgdevice_firewall_internetserviceextension" "trname" {
  comment = "your own value"
  disable_entry {
    addr_mode = "ipv6"
    fosid     = 10
    ip_range {
      end_ip   = "your own value"
      fosid    = 10
      start_ip = "your own value"
    }

    ip6_range {
      end_ip6   = "your own value"
      fosid     = 10
      start_ip6 = "your own value"
    }

    port_range {
      end_port   = 10
      fosid      = 10
      start_port = 10
    }

    protocol = 10
  }

  entry {
    addr_mode = "ipv4"
    dst       = ["your own value"]
    dst6      = ["your own value"]
    fosid     = 10
    port_range {
      end_port   = 10
      fosid      = 10
      start_port = 10
    }

    protocol = 10
  }

  fosid       = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `disable_entry` - Disable-Entry. The structure of `disable_entry` block is documented below.
* `entry` - Entry. The structure of `entry` block is documented below.
* `fosid` - Internet Service ID in the Internet Service database.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `disable_entry` block supports:

* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

* `id` - Disable entry ID.
* `ip_range` - Ip-Range. The structure of `ip_range` block is documented below.
* `ip6_range` - Ip6-Range. The structure of `ip6_range` block is documented below.
* `port_range` - Port-Range. The structure of `port_range` block is documented below.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).

The `ip_range` block supports:

* `end_ip` - End IPv4 address.
* `id` - Disable entry range ID.
* `start_ip` - Start IPv4 address.

The `ip6_range` block supports:

* `end_ip6` - End IPv6 address.
* `id` - Disable entry range ID.
* `start_ip6` - Start IPv6 address.

The `port_range` block supports:

* `end_port` - Ending TCP/UDP/SCTP destination port (0 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Starting TCP/UDP/SCTP destination port (0 to 65535).

The `entry` block supports:

* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.

* `dst` - Destination address or address group name.
* `dst6` - Destination address6 or address6 group name.
* `id` - Entry ID(1-255).
* `port_range` - Port-Range. The structure of `port_range` block is documented below.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).

The `port_range` block supports:

* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (0 to 65535).
* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (0 to 65535).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceExtension can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_internetserviceextension.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

