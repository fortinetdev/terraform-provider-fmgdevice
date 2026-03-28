---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_dynamicbypass"
description: |-
  Device WebProxyDynamicBypass
---

# fmgdevice_webproxy_dynamicbypass
Device WebProxyDynamicBypass

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `errors` - Errors. Valid values: `connect-error`, `receive-error`, `non-http`, `400`, `401`, `403`, `405`, `406`, `500`, `502`, `503`, `504`.

* `server_max` - Server-Max.
* `status` - Status. Valid values: `disable`, `enable`.

* `timeout` - Timeout.
* `total_max` - Total-Max.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WebProxy DynamicBypass can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_dynamicbypass.labelname WebProxyDynamicBypass
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

