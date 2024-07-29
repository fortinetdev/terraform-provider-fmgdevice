---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_setting"
description: |-
  Configure router settings.
---

# fmgdevice_router_setting
Configure router settings.

## Example Usage

```hcl
resource "fmgdevice_router_setting" "trname" {
  hostname    = "your own value"
  show_filter = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `hostname` - Hostname for this virtual domain router.
* `show_filter` - Prefix-list as filter for showing routes.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Setting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_setting.labelname RouterSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

