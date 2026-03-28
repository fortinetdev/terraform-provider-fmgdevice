---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_shaper_trafficshaper"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure shared traffic shaper.
---

# fmgdevice_firewall_shaper_trafficshaper
<i>This object will be purged after policy copy and install.</i> Configure shared traffic shaper.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `bandwidth_unit` - Unit of measurement for guaranteed and maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.

* `cos` - VLAN CoS mark.
* `cos_marking` - Enable/disable VLAN CoS marking. Valid values: `disable`, `enable`.

* `cos_marking_method` - Select VLAN CoS marking method. Valid values: `multi-stage`, `static`.

* `diffserv` - Enable/disable changing the DiffServ setting applied to traffic accepted by this shaper. Valid values: `disable`, `enable`.

* `diffservcode` - DiffServ setting to be applied to traffic accepted by this shaper.
* `dscp_marking_method` - Select DSCP marking method. Valid values: `multi-stage`, `static`.

* `exceed_bandwidth` - Exceed bandwidth used for DSCP/VLAN CoS multi-stage marking. Units depend on the bandwidth-unit setting.
* `exceed_class_id` - Class ID for traffic in guaranteed-bandwidth and maximum-bandwidth.
* `exceed_cos` - VLAN CoS mark for traffic in [guaranteed-bandwidth, exceed-bandwidth].
* `exceed_dscp` - DSCP mark for traffic in guaranteed-bandwidth and exceed-bandwidth.
* `guaranteed_bandwidth` - Amount of bandwidth guaranteed for this shaper (0 - 100000000). Units depend on the bandwidth-unit setting.
* `maximum_bandwidth` - Upper bandwidth limit enforced by this shaper (0 - 100000000). 0 means no limit. Units depend on the bandwidth-unit setting.
* `maximum_cos` - VLAN CoS mark for traffic in [exceed-bandwidth, maximum-bandwidth].
* `maximum_dscp` - DSCP mark for traffic in exceed-bandwidth and maximum-bandwidth.
* `name` - Traffic shaper name.
* `overhead` - Per-packet size overhead used in rate computations.
* `per_policy` - Enable/disable applying a separate shaper for each policy. For example, if enabled the guaranteed bandwidth is applied separately for each policy. Valid values: `disable`, `enable`.

* `priority` - Higher priority traffic is more likely to be forwarded without delays and without compromising the guaranteed bandwidth. Valid values: `high`, `medium`, `low`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ShaperTrafficShaper can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_shaper_trafficshaper.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

