---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ips_sensor_entries_exemptip"
description: |-
  <i>This object will be purged after policy copy and install.</i> Traffic from selected source or destination IP addresses is exempt from this signature.
---

# fmgdevice_ips_sensor_entries_exemptip
<i>This object will be purged after policy copy and install.</i> Traffic from selected source or destination IP addresses is exempt from this signature.

~> This resource is a sub resource for variable `exempt_ip` of resource `fmgdevice_ips_sensor_entries`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `sensor` - Sensor.
* `entries` - Entries.

* `dst_ip` - Destination IP address and netmask (applies to packet matching the signature).
* `fosid` - Exempt IP ID.
* `src_ip` - Source IP address and netmask (applies to packet matching the signature).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Ips SensorEntriesExemptIp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "sensor=YOUR_VALUE", "entries=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ips_sensor_entries_exemptip.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

