---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webcache_reversecacheserver"
description: |-
  Device WebcacheReverseCacheServer
---

# fmgdevice_webcache_reversecacheserver
Device WebcacheReverseCacheServer

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `ip` - Ip.
* `name` - Name.
* `port` - Port.
* `prefetch_file` - Prefetch-File.
* `priority` - Priority.
* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Webcache ReverseCacheServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webcache_reversecacheserver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

