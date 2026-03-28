---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_vap_ip6prefixlist"
description: |-
  Device WirelessControllerVapIp6PrefixList
---

# fmgdevice_wirelesscontroller_vap_ip6prefixlist
Device WirelessControllerVapIp6PrefixList

~> This resource is a sub resource for variable `ip6_prefix_list` of resource `fmgdevice_wirelesscontroller_vap`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `vap` - Vap.

* `autonomous_flag` - Autonomous-Flag. Valid values: `disable`, `enable`.

* `dnssl` - Dnssl.
* `onlink_flag` - Onlink-Flag. Valid values: `disable`, `enable`.

* `preferred_life_time` - Preferred-Life-Time.
* `prefix` - Prefix.
* `rdnss` - Rdnss.
* `valid_life_time` - Valid-Life-Time.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController VapIp6PrefixList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "vap=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_vap_ip6prefixlist.labelname WirelessControllerVapIp6PrefixList
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

