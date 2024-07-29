---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_hotspot20_h2qposuprovider_servicedescription"
description: |-
  OSU service name.
---

# fmgdevice_wirelesscontroller_hotspot20_h2qposuprovider_servicedescription
OSU service name.

~> This resource is a sub resource for variable `service_description` of resource `fmgdevice_wirelesscontroller_hotspot20_h2qposuprovider`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_hotspot20_h2qposuprovider_servicedescription" "trname" {
  lang                = "your own value"
  service_description = "your own value"
  service_id          = 10
  device_name         = var.device_name # not required if setting is at provider
  device_vdom         = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `h2qp_osu_provider` - H2Qp Osu Provider.

* `lang` - Language code.
* `service_description` - Service description.
* `service_id` - OSU service ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{service_id}}.

## Import

WirelessController Hotspot20H2QpOsuProviderServiceDescription can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "h2qp_osu_provider=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_hotspot20_h2qposuprovider_servicedescription.labelname {{service_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

