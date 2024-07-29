---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_interface_dhcpsnoopingserverlist"
description: |-
  Configure DHCP server access list.
---

# fmgdevice_system_interface_dhcpsnoopingserverlist
Configure DHCP server access list.

~> This resource is a sub resource for variable `dhcp_snooping_server_list` of resource `fmgdevice_system_interface`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_interface_dhcpsnoopingserverlist" "trname" {
  name        = "your own value"
  server_ip   = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `interface` - Interface.

* `name` - DHCP server name.
* `server_ip` - IP address for DHCP server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System InterfaceDhcpSnoopingServerList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_interface_dhcpsnoopingserverlist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

