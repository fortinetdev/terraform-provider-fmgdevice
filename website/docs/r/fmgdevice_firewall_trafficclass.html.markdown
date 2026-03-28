---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_trafficclass"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure names for shaping classes.
---

# fmgdevice_firewall_trafficclass
<i>This object will be purged after policy copy and install.</i> Configure names for shaping classes.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `class_id` - Class ID to be named.
* `class_name` - Define the name for this class-id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{class_id}}.

## Import

Firewall TrafficClass can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_trafficclass.labelname {{class_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

