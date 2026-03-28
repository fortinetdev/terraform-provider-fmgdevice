---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_systemdhcpserver_options"
description: |-
  DHCP options.
---

# fmgdevice_switchcontroller_managedswitch_systemdhcpserver_options
DHCP options.

~> This resource is a sub resource for variable `options` of resource `fmgdevice_switchcontroller_managedswitch_systemdhcpserver`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.
* `system_dhcp_server` - System Dhcp Server.

* `code` - DHCP option code.
* `fosid` - ID.
* `ip` - DHCP option IPs.
* `type` - DHCP option type. Valid values: `hex`, `string`, `ip`, `fqdn`.

* `value` - DHCP option value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SwitchController ManagedSwitchSystemDhcpServerOptions can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE", "system_dhcp_server=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_systemdhcpserver_options.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

