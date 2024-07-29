---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_anqpvenuename_valuelist"
description: |-
  Name list.
---

# fmgdevice_wirelesscontroller_hotspot20_anqpvenuename_valuelist
Name list.

~> This resource is a sub resource for variable `value_list` of resource `fmgdevice_wirelesscontroller_hotspot20_anqpvenuename`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_anqpvenuename_valuelist" "trname" {
  index       = 10
  lang        = "your own value"
  value       = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `anqp_venue_name` - Anqp Venue Name.

* `index` - Value index.
* `lang` - Language code.
* `value` - Venue name value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{index}}.

## Import

WirelessController Hotspot20AnqpVenueNameValueList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "anqp_venue_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_anqpvenuename_valuelist.labelname {{index}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

