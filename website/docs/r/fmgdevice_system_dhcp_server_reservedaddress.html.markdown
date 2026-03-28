---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_dhcp_server_reservedaddress"
description: |-
  Options for the DHCP server to assign IP settings to specific MAC addresses.
---

# fmgdevice_system_dhcp_server_reservedaddress
Options for the DHCP server to assign IP settings to specific MAC addresses.

~> This resource is a sub resource for variable `reserved_address` of resource `fmgdevice_system_dhcp_server`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `server` - Server.

* `action` - Options for the DHCP server to configure the client with the reserved MAC address. Valid values: `assign`, `block`, `reserved`.

* `circuit_id` - Option 82 circuit-ID of the client that will get the reserved IP address.
* `circuit_id_type` - DHCP option type. Valid values: `hex`, `string`.

* `description` - Description.
* `fosid` - ID.
* `ip` - IP address to be reserved for the MAC address.
* `mac` - MAC address of the client that will get the reserved IP address.
* `remote_id` - Option 82 remote-ID of the client that will get the reserved IP address.
* `remote_id_type` - DHCP option type. Valid values: `hex`, `string`.

* `type` - DHCP reserved-address type. Valid values: `mac`, `option82`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System DhcpServerReservedAddress can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "server=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_dhcp_server_reservedaddress.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

