---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_ftgdlocalrisk"
description: |-
  Configure FortiGuard Web Filter local risk score.
---

# fmgdevice_webfilter_ftgdlocalrisk
Configure FortiGuard Web Filter local risk score.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `risk_score` - Local risk score.
* `status` - Enable/disable local risk score. Valid values: `disable`, `enable`.

* `url` - URL to rate locally.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{url}}.

## Import

Webfilter FtgdLocalRisk can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_ftgdlocalrisk.labelname {{url}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

