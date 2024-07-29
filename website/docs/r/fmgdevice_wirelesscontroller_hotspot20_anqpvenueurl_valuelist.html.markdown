---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_anqpvenueurl_valuelist"
description: |-
  URL list.
---

# fmgdevice_wirelesscontroller_hotspot20_anqpvenueurl_valuelist
URL list.

~> This resource is a sub resource for variable `value_list` of resource `fmgdevice_wirelesscontroller_hotspot20_anqpvenueurl`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_anqpvenueurl_valuelist" "trname" {
  index       = 10
  number      = 10
  value       = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `anqp_venue_url` - Anqp Venue Url.

* `index` - URL index.
* `number` - Venue number.
* `value` - Venue URL value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{index}}.

## Import

WirelessController Hotspot20AnqpVenueUrlValueList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "anqp_venue_url=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_anqpvenueurl_valuelist.labelname {{index}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

