---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_dhcp6_server_iprange"
description: |-
  DHCP IP range configuration.
---

# fmgdevice_system_dhcp6_server_iprange
DHCP IP range configuration.

~> This resource is a sub resource for variable `ip_range` of resource `fmgdevice_system_dhcp6_server`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_dhcp6_server_iprange" "trname" {
  end_ip      = "your own value"
  fosid       = 10
  start_ip    = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `server` - Server.

* `end_ip` - End of IP range.
* `fosid` - ID.
* `start_ip` - Start of IP range.
* `vci_match` - Enable/disable vendor class option matching. When enabled only DHCP requests with a matching VC are served with this range. Valid values: `disable`, `enable`.

* `vci_string` - One or more VCI strings in quotes separated by spaces.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System Dhcp6ServerIpRange can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "server=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_dhcp6_server_iprange.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

