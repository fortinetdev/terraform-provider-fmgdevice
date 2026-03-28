---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_shapingprofile_shapingentries"
description: |-
  <i>This object will be purged after policy copy and install.</i> Define shaping entries of this shaping profile.
---

# fmgdevice_firewall_shapingprofile_shapingentries
<i>This object will be purged after policy copy and install.</i> Define shaping entries of this shaping profile.

~> This resource is a sub resource for variable `shaping_entries` of resource `fmgdevice_firewall_shapingprofile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `shaping_profile` - Shaping Profile.

* `burst_in_msec` - Number of bytes that can be burst at maximum-bandwidth speed. Formula: burst = maximum-bandwidth*burst-in-msec.
* `cburst_in_msec` - Number of bytes that can be burst as fast as the interface can transmit. Formula: cburst = maximum-bandwidth*cburst-in-msec.
* `class_id` - Class ID.
* `guaranteed_bandwidth_percentage` - Guaranteed bandwidth in percentage.
* `fosid` - ID number.
* `limit` - Hard limit on the real queue size in packets.
* `max` - Average queue size in packets at which RED drop probability is maximal.
* `maximum_bandwidth_percentage` - Maximum bandwidth in percentage.
* `min` - Average queue size in packets at which RED drop becomes a possibility.
* `priority` - Priority. Valid values: `low`, `medium`, `high`, `critical`, `top`.

* `red_probability` - Maximum probability (in percentage) for RED marking.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall ShapingProfileShapingEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "shaping_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_shapingprofile_shapingentries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

