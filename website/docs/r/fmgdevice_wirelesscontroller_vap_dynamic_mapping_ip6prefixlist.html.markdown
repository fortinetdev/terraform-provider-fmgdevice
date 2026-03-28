---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_vap_dynamic_mapping_ip6prefixlist"
description: |-
  Device WirelessControllerVapDynamicMappingIp6PrefixList
---

# fmgdevice_wirelesscontroller_vap_dynamic_mapping_ip6prefixlist
Device WirelessControllerVapDynamicMappingIp6PrefixList

~> This resource is a sub resource for variable `ip6_prefix_list` of resource `fmgdevice_wirelesscontroller_vap_dynamic_mapping`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `vap` - Vap.
* `dynamic_mapping_name` - Dynamic Mapping Name.
* `dynamic_mapping_vdom` - Dynamic Mapping Vdom.

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

WirelessController VapDynamicMappingIp6PrefixList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "vap=YOUR_VALUE", "dynamic_mapping_name=YOUR_VALUE", "dynamic_mapping_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_vap_dynamic_mapping_ip6prefixlist.labelname WirelessControllerVapDynamicMappingIp6PrefixList
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

