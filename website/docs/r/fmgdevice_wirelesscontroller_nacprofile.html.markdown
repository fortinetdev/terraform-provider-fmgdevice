---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_nacprofile"
description: |-
  Configure WiFi network access control (NAC) profiles.
---

# fmgdevice_wirelesscontroller_nacprofile
Configure WiFi network access control (NAC) profiles.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_nacprofile" "trname" {
  comment         = "your own value"
  name            = "your own value"
  onboarding_vlan = ["your own value"]
  device_name     = var.device_name # not required if setting is at provider
  device_vdom     = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `name` - Name.
* `onboarding_vlan` - VLAN interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController NacProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_nacprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

