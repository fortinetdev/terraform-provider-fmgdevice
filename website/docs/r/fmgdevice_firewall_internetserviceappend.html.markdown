---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_internetserviceappend"
description: |-
  Configure additional port mappings for Internet Services.
---

# fmgdevice_firewall_internetserviceappend
Configure additional port mappings for Internet Services.

## Example Usage

```hcl
resource "fmgdevice_firewall_internetserviceappend" "trname" {
  addr_mode   = "ipv4"
  append_port = 10
  match_port  = 10
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `addr_mode` - Address mode (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`, `both`.

* `append_port` - Appending TCP/UDP/SCTP destination port (1 to 65535).
* `match_port` - Matching TCP/UDP/SCTP destination port (0 to 65535, 0 means any port).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall InternetServiceAppend can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_internetserviceappend.labelname FirewallInternetServiceAppend
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

