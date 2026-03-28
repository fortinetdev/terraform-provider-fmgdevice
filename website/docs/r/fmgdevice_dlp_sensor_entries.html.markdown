---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_sensor_entries"
description: |-
  <i>This object will be purged after policy copy and install.</i> DLP sensor entries.
---

# fmgdevice_dlp_sensor_entries
<i>This object will be purged after policy copy and install.</i> DLP sensor entries.

~> This resource is a sub resource for variable `entries` of resource `fmgdevice_dlp_sensor`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `sensor` - Sensor.

* `fmgcount` - Count of dictionary matches to trigger sensor entry match (Dictionary might not be able to trigger more than once based on its 'repeat' option, 1 - 255, default = 1).
* `dictionary` - Select a DLP dictionary or exact-data-match.
* `fosid` - ID.
* `status` - Enable/disable this entry. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Dlp SensorEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "sensor=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_sensor_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

