---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_shapingprofile"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure shaping profiles.
---

# fmgdevice_firewall_shapingprofile
<i>This object will be purged after policy copy and install.</i> Configure shaping profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `classes`: `fmgdevice_firewall_shapingprofile_classes`
>- `shaping_entries`: `fmgdevice_firewall_shapingprofile_shapingentries`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `default_class_id` - Default class ID to handle unclassified packets (including all local traffic).
* `npu_offloading` - Enable/disable NPU offloading. Valid values: `disable`, `enable`.

* `profile_name` - Shaping profile name.
* `shaping_entries` - Shaping-Entries. The structure of `shaping_entries` block is documented below.
* `type` - Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.

* `classes` - Classes. The structure of `classes` block is documented below.
* `default_class` - Default-Class.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `shaping_entries` block supports:

* `burst_in_msec` - Number of bytes that can be burst at maximum-bandwidth speed. Formula: burst = maximum-bandwidth*burst-in-msec.
* `cburst_in_msec` - Number of bytes that can be burst as fast as the interface can transmit. Formula: cburst = maximum-bandwidth*cburst-in-msec.
* `class_id` - Class ID.
* `guaranteed_bandwidth_percentage` - Guaranteed bandwidth in percentage.
* `id` - ID number.
* `limit` - Hard limit on the real queue size in packets.
* `max` - Average queue size in packets at which RED drop probability is maximal.
* `maximum_bandwidth_percentage` - Maximum bandwidth in percentage.
* `min` - Average queue size in packets at which RED drop becomes a possibility.
* `priority` - Priority. Valid values: `low`, `medium`, `high`, `critical`, `top`.

* `red_probability` - Maximum probability (in percentage) for RED marking.

The `classes` block supports:

* `class_id` - Class-Id.
* `guaranteed_bandwidth` - Guaranteed-Bandwidth.
* `maximum_bandwidth` - Maximum-Bandwidth.
* `name` - Name.
* `priority` - Priority. Valid values: `top`, `critical`, `high`, `medium`, `low`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{profile_name}}.

## Import

Firewall ShapingProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_shapingprofile.labelname {{profile_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

