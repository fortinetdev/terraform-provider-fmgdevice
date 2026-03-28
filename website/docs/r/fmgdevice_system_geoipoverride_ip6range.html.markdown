---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_geoipoverride_ip6range"
description: |-
  <i>This object will be purged after policy copy and install.</i> Table of IPv6 ranges assigned to country.
---

# fmgdevice_system_geoipoverride_ip6range
<i>This object will be purged after policy copy and install.</i> Table of IPv6 ranges assigned to country.

~> This resource is a sub resource for variable `ip6_range` of resource `fmgdevice_system_geoipoverride`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `geoip_override` - Geoip Override.

* `end_ip` - Ending IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `fosid` - ID of individual entry in the IPv6 range table.
* `start_ip` - Starting IP address, inclusive, of the address range (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System GeoipOverrideIp6Range can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "geoip_override=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_geoipoverride_ip6range.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

