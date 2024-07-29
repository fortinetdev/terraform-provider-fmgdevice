---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_h2qposuprovidernai_nailist"
description: |-
  OSU NAI list.
---

# fmgdevice_wirelesscontroller_hotspot20_h2qposuprovidernai_nailist
OSU NAI list.

~> This resource is a sub resource for variable `nai_list` of resource `fmgdevice_wirelesscontroller_hotspot20_h2qposuprovidernai`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_h2qposuprovidernai_nailist" "trname" {
  name        = "your own value"
  osu_nai     = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `h2qp_osu_provider_nai` - H2Qp Osu Provider Nai.

* `name` - OSU NAI ID.
* `osu_nai` - OSU NAI.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Hotspot20H2QpOsuProviderNaiNaiList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "h2qp_osu_provider_nai=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_h2qposuprovidernai_nailist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

