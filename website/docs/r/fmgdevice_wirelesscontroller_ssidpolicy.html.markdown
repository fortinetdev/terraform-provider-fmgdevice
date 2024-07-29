---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_ssidpolicy"
description: |-
  Configure WiFi SSID policies.
---

# fmgdevice_wirelesscontroller_ssidpolicy
Configure WiFi SSID policies.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_ssidpolicy" "trname" {
  description = "your own value"
  name        = "your own value"
  vlan        = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `description` - Description.
* `name` - Name.
* `vlan` - VLAN interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController SsidPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_ssidpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

