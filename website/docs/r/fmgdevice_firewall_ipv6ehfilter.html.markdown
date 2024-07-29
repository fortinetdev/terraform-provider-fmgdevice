---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ipv6ehfilter"
description: |-
  Configure IPv6 extension header filter.
---

# fmgdevice_firewall_ipv6ehfilter
Configure IPv6 extension header filter.

## Example Usage

```hcl
resource "fmgdevice_firewall_ipv6ehfilter" "trname" {
  auth        = "disable"
  dest_opt    = "enable"
  fragment    = "enable"
  hdopt_type  = 10
  hop_opt     = "disable"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `auth` - Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `disable`, `enable`.

* `dest_opt` - Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `disable`, `enable`.

* `fragment` - Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `disable`, `enable`.

* `hdopt_type` - Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255).
* `hop_opt` - Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `disable`, `enable`.

* `no_next` - Enable/disable blocking packets with the No Next header (default = disable). Valid values: `disable`, `enable`.

* `routing` - Enable/disable blocking packets with Routing headers (default = enable). Valid values: `disable`, `enable`.

* `routing_type` - Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall Ipv6EhFilter can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ipv6ehfilter.labelname FirewallIpv6EhFilter
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

