---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_systemdhcpserver_iprange"
description: |-
  DHCP IP range configuration.
---

# fmgdevice_switchcontroller_managedswitch_systemdhcpserver_iprange
DHCP IP range configuration.

~> This resource is a sub resource for variable `ip_range` of resource `fmgdevice_switchcontroller_managedswitch_systemdhcpserver`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.
* `system_dhcp_server` - System Dhcp Server.

* `end_ip` - End of IP range.
* `fosid` - ID.
* `start_ip` - Start of IP range.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SwitchController ManagedSwitchSystemDhcpServerIpRange can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE", "system_dhcp_server=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_systemdhcpserver_iprange.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

