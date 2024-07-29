---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_snmp_community_hosts"
description: |-
  Configure IPv4 SNMP managers (hosts).
---

# fmgdevice_wirelesscontroller_snmp_community_hosts
Configure IPv4 SNMP managers (hosts).

~> This resource is a sub resource for variable `hosts` of resource `fmgdevice_wirelesscontroller_snmp_community`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_snmp_community_hosts" "trname" {
  fosid       = 10
  ip          = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `community` - Community.

* `fosid` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WirelessController SnmpCommunityHosts can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "community=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_snmp_community_hosts.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

