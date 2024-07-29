---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_interface_ipv6_dhcp6iapdlist"
description: |-
  DHCPv6 IA-PD list.
---

# fmgdevice_system_interface_ipv6_dhcp6iapdlist
DHCPv6 IA-PD list.

~> This resource is a sub resource for variable `dhcp6_iapd_list` of resource `fmgdevice_system_interface_ipv6`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_interface_ipv6_dhcp6iapdlist" "trname" {
  iaid            = 10
  prefix_hint     = "your own value"
  prefix_hint_plt = 10
  prefix_hint_vlt = 10
  device_name     = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `interface` - Interface.

* `iaid` - Identity association identifier.
* `prefix_hint` - DHCPv6 prefix that will be used as a hint to the upstream DHCPv6 server.
* `prefix_hint_plt` - DHCPv6 prefix hint preferred life time (sec), 0 means unlimited lease time.
* `prefix_hint_vlt` - DHCPv6 prefix hint valid life time (sec).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{iaid}}.

## Import

System InterfaceIpv6Dhcp6IapdList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_interface_ipv6_dhcp6iapdlist.labelname {{iaid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

