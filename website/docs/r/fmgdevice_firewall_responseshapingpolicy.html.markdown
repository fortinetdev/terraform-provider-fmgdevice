---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_responseshapingpolicy"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_firewall_responseshapingpolicy
<i>This object will be purged after policy copy and install.</i>

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `class_id` - Class-Id.
* `class_id_reverse` - Class-Id-Reverse.
* `comment` - Comment.
* `dstaddr` - Dstaddr.
* `dstaddr6` - Dstaddr6.
* `fosid` - Id.
* `ip_version` - Ip-Version. Valid values: `6`, `4`.

* `name` - Name.
* `per_ip_shaper` - Per-Ip-Shaper.
* `schedule` - Schedule.
* `srcaddr` - Srcaddr.
* `status` - Status. Valid values: `disable`, `enable`.

* `traffic_shaper` - Traffic-Shaper.
* `traffic_shaper_reverse` - Traffic-Shaper-Reverse.
* `uuid` - Uuid.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall ResponseShapingPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_responseshapingpolicy.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

