---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_anqp3gppcellular_mccmnclist"
description: |-
  Mobile Country Code and Mobile Network Code configuration.
---

# fmgdevice_wirelesscontroller_hotspot20_anqp3gppcellular_mccmnclist
Mobile Country Code and Mobile Network Code configuration.

~> This resource is a sub resource for variable `mcc_mnc_list` of resource `fmgdevice_wirelesscontroller_hotspot20_anqp3gppcellular`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_anqp3gppcellular_mccmnclist" "trname" {
  fosid       = 10
  mcc         = "your own value"
  mnc         = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `anqp_3gpp_cellular` - Anqp 3Gpp Cellular.

* `fosid` - ID.
* `mcc` - Mobile country code.
* `mnc` - Mobile network code.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WirelessController Hotspot20Anqp3GppCellularMccMncList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "anqp_3gpp_cellular=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_anqp3gppcellular_mccmnclist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

