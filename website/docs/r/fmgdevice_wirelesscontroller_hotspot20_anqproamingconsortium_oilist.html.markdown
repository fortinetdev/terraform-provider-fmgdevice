---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_anqproamingconsortium_oilist"
description: |-
  Organization identifier list.
---

# fmgdevice_wirelesscontroller_hotspot20_anqproamingconsortium_oilist
Organization identifier list.

~> This resource is a sub resource for variable `oi_list` of resource `fmgdevice_wirelesscontroller_hotspot20_anqproamingconsortium`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_anqproamingconsortium_oilist" "trname" {
  comment     = "your own value"
  index       = 10
  oi          = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `anqp_roaming_consortium` - Anqp Roaming Consortium.

* `comment` - Comment.
* `index` - OI index.
* `oi` - Organization identifier.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{index}}.

## Import

WirelessController Hotspot20AnqpRoamingConsortiumOiList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "anqp_roaming_consortium=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_anqproamingconsortium_oilist.labelname {{index}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

