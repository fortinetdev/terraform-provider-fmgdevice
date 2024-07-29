---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_h2qposuprovider"
description: |-
  Configure online sign up (OSU) provider list.
---

# fmgdevice_wirelesscontroller_hotspot20_h2qposuprovider
Configure online sign up (OSU) provider list.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `friendly_name`: `fmgdevice_wirelesscontroller_hotspot20_h2qposuprovider_friendlyname`
>- `service_description`: `fmgdevice_wirelesscontroller_hotspot20_h2qposuprovider_servicedescription`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_h2qposuprovider" "trname" {
  friendly_name {
    friendly_name = "your own value"
    index         = 10
    lang          = "your own value"
  }

  icon        = ["your own value"]
  name        = "your own value"
  osu_method  = ["oma-dm"]
  osu_nai     = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `friendly_name` - Friendly-Name. The structure of `friendly_name` block is documented below.
* `icon` - OSU provider icon.
* `name` - OSU provider ID.
* `osu_method` - OSU method list. Valid values: `oma-dm`, `soap-xml-spp`, `reserved`.

* `osu_nai` - OSU NAI.
* `server_uri` - Server URI.
* `service_description` - Service-Description. The structure of `service_description` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `friendly_name` block supports:

* `friendly_name` - OSU provider friendly name.
* `index` - OSU provider friendly name index.
* `lang` - Language code.

The `service_description` block supports:

* `lang` - Language code.
* `service_description` - Service description.
* `service_id` - OSU service ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController Hotspot20H2QpOsuProvider can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_h2qposuprovider.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

