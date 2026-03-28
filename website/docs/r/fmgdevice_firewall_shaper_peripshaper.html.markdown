---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_shaper_peripshaper"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure per-IP traffic shaper.
---

# fmgdevice_firewall_shaper_peripshaper
<i>This object will be purged after policy copy and install.</i> Configure per-IP traffic shaper.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `bandwidth_unit` - Unit of measurement for maximum bandwidth for this shaper (Kbps, Mbps or Gbps). Valid values: `kbps`, `mbps`, `gbps`.

* `diffserv_forward` - Enable/disable changing the Forward (original) DiffServ setting applied to traffic accepted by this shaper. Valid values: `disable`, `enable`.

* `diffserv_reverse` - Enable/disable changing the Reverse (reply) DiffServ setting applied to traffic accepted by this shaper. Valid values: `disable`, `enable`.

* `diffservcode_forward` - Forward (original) DiffServ setting to be applied to traffic accepted by this shaper.
* `diffservcode_rev` - Reverse (reply) DiffServ setting to be applied to traffic accepted by this shaper.
* `max_bandwidth` - Upper bandwidth limit enforced by this shaper (0 - 100000000). 0 means no limit. Units depend on the bandwidth-unit setting.
* `max_concurrent_session` - Maximum number of concurrent sessions allowed by this shaper (0 - 2097000). 0 means no limit.
* `max_concurrent_tcp_session` - Maximum number of concurrent TCP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
* `max_concurrent_udp_session` - Maximum number of concurrent UDP sessions allowed by this shaper (0 - 2097000). 0 means no limit.
* `name` - Traffic shaper name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ShaperPerIpShaper can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_shaper_peripshaper.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

