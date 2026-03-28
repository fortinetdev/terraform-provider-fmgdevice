---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_systemdhcpserver"
description: |-
  Configure DHCP servers.
---

# fmgdevice_switchcontroller_managedswitch_systemdhcpserver
Configure DHCP servers.

~> This resource is a sub resource for variable `system_dhcp_server` of resource `fmgdevice_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `ip_range`: `fmgdevice_switchcontroller_managedswitch_systemdhcpserver_iprange`
>- `options`: `fmgdevice_switchcontroller_managedswitch_systemdhcpserver_options`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `default_gateway` - Default gateway IP address assigned by the DHCP server.
* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `dns_server3` - DNS server 3.
* `dns_service` - Options for assigning DNS servers to DHCP clients. Valid values: `default`, `specify`, `local`.

* `fosid` - ID.
* `interface` - DHCP server can assign IP configurations to clients connected to this interface.
* `ip_range` - Ip-Range. The structure of `ip_range` block is documented below.
* `lease_time` - Lease time in seconds, 0 means unlimited.
* `netmask` - Netmask assigned by the DHCP server.
* `ntp_server1` - NTP server 1.
* `ntp_server2` - NTP server 2.
* `ntp_server3` - NTP server 3.
* `ntp_service` - Options for assigning Network Time Protocol (NTP) servers to DHCP clients. Valid values: `default`, `specify`, `local`.

* `options` - Options. The structure of `options` block is documented below.
* `status` - Enable/disable this DHCP configuration. Valid values: `disable`, `enable`.

* `switch_id` - Switch ID.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ip_range` block supports:

* `end_ip` - End of IP range.
* `id` - ID.
* `start_ip` - Start of IP range.

The `options` block supports:

* `code` - DHCP option code.
* `id` - ID.
* `ip` - DHCP option IPs.
* `type` - DHCP option type. Valid values: `hex`, `string`, `ip`, `fqdn`.

* `value` - DHCP option value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SwitchController ManagedSwitchSystemDhcpServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_systemdhcpserver.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

