---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webcache_prefetch"
description: |-
  Device WebcachePrefetch
---

# fmgdevice_webcache_prefetch
Device WebcachePrefetch

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `crawl_depth` - Crawl-Depth.
* `ignore_robots` - Ignore-Robots. Valid values: `disable`, `enable`.

* `interval` - Interval.
* `name` - Name.
* `password` - Password.
* `repeat` - Repeat.
* `start_delay` - Start-Delay.
* `url` - Url.
* `user` - User.
* `user_agent` - User-Agent.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Webcache Prefetch can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webcache_prefetch.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

