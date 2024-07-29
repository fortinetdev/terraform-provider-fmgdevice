---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_ipsurlfiltersetting"
description: |-
  Configure IPS URL filter settings.
---

# fmgdevice_webfilter_ipsurlfiltersetting
Configure IPS URL filter settings.

## Example Usage

```hcl
resource "fmgdevice_webfilter_ipsurlfiltersetting" "trname" {
  device      = "your own value"
  distance    = 10
  gateway     = "your own value"
  geo_filter  = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `device` - Interface for this route.
* `distance` - Administrative distance (1 - 255) for this route.
* `gateway` - Gateway IP address for this route.
* `geo_filter` - Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Webfilter IpsUrlfilterSetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_ipsurlfiltersetting.labelname WebfilterIpsUrlfilterSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

