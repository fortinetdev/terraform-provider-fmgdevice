---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_dhcp_server_excluderange"
description: |-
  Exclude one or more ranges of IP addresses from being assigned to clients.
---

# fmgdevice_system_dhcp_server_excluderange
Exclude one or more ranges of IP addresses from being assigned to clients.

~> This resource is a sub resource for variable `exclude_range` of resource `fmgdevice_system_dhcp_server`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `server` - Server.

* `end_ip` - End of IP range.
* `fosid` - ID.
* `lease_time` - Lease time in seconds, 0 means default lease time.
* `start_ip` - Start of IP range.
* `uci_match` - Enable/disable user class identifier (UCI) matching. When enabled only DHCP requests with a matching UCI are served with this range. Valid values: `disable`, `enable`.

* `uci_string` - One or more UCI strings in quotes separated by spaces.
* `vci_match` - Enable/disable vendor class identifier (VCI) matching. When enabled only DHCP requests with a matching VCI are served with this range. Valid values: `disable`, `enable`.

* `vci_string` - One or more VCI strings in quotes separated by spaces.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System DhcpServerExcludeRange can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "server=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_dhcp_server_excluderange.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

