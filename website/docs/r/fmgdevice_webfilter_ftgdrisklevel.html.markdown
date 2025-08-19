---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_ftgdrisklevel"
description: |-
  Configure FortiGuard Web Filter risk level.
---

# fmgdevice_webfilter_ftgdrisklevel
Configure FortiGuard Web Filter risk level.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `high` - Risk level high.
* `low` - Risk level low.
* `name` - Risk level name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Webfilter FtgdRiskLevel can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_ftgdrisklevel.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

