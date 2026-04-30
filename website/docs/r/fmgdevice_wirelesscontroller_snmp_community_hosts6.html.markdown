---
subcategory: "Wireless Controller"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_snmp_community_hosts6"
description: |-
  Configure IPv6 SNMP managers (hosts).
---

# fmgdevice_wirelesscontroller_snmp_community_hosts6
Configure IPv6 SNMP managers (hosts).

~> This resource is a sub resource for variable `hosts6` of resource `fmgdevice_wirelesscontroller_snmp_community`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `community` - Community.

* `fosid` - Host6 entry ID.
* `ipv6` - IPv6 address of the SNMP manager (host).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WirelessController SnmpCommunityHosts6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "community=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_snmp_community_hosts6.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

