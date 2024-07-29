---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_centralmanagement_serverlist"
description: |-
  Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers.
---

# fmgdevice_system_centralmanagement_serverlist
Additional severs that the FortiGate can use for updates (for AV, IPS, updates) and ratings (for web filter and antispam ratings) servers.

~> This resource is a sub resource for variable `server_list` of resource `fmgdevice_system_centralmanagement`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_centralmanagement_serverlist" "trname" {
  addr_type       = "ipv6"
  fqdn            = "your own value"
  fosid           = 10
  server_address  = "your own value"
  server_address6 = "your own value"
  device_name     = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `addr_type` - Indicate whether the FortiGate communicates with the override server using an IPv4 address, an IPv6 address or a FQDN. Valid values: `fqdn`, `ipv4`, `ipv6`.

* `fqdn` - FQDN address of override server.
* `fosid` - ID.
* `server_address` - IPv4 address of override server.
* `server_address6` - IPv6 address of override server.
* `server_type` - FortiGuard service type. Valid values: `update`, `rating`, `iot-query`, `iot-collect`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System CentralManagementServerList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_centralmanagement_serverlist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

