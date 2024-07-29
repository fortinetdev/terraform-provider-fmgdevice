---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_dhcp6_server_prefixrange"
description: |-
  DHCP prefix configuration.
---

# fmgdevice_system_dhcp6_server_prefixrange
DHCP prefix configuration.

~> This resource is a sub resource for variable `prefix_range` of resource `fmgdevice_system_dhcp6_server`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_dhcp6_server_prefixrange" "trname" {
  end_prefix    = "your own value"
  fosid         = 10
  prefix_length = 10
  start_prefix  = "your own value"
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `server` - Server.

* `end_prefix` - End of prefix range.
* `fosid` - ID.
* `prefix_length` - Prefix length.
* `start_prefix` - Start of prefix range.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System Dhcp6ServerPrefixRange can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "server=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_dhcp6_server_prefixrange.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

