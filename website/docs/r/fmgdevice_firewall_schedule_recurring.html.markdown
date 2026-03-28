---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_schedule_recurring"
description: |-
  <i>This object will be purged after policy copy and install.</i> Recurring schedule configuration.
---

# fmgdevice_firewall_schedule_recurring
<i>This object will be purged after policy copy and install.</i> Recurring schedule configuration.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `color` - Color of icon on the GUI.
* `day` - One or more days of the week on which the schedule is valid. Separate the names of the days with a space. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`, `none`.

* `end` - Time of day to end the schedule, format hh:mm.
* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `global_object` - Global-Object.
* `name` - Recurring schedule name.
* `start` - Time of day to start the schedule, format hh:mm.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ScheduleRecurring can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_schedule_recurring.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

