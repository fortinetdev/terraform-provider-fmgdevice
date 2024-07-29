---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extensioncontroller_fortigate"
description: |-
  FortiGate controller configuration.
---

# fmgdevice_extensioncontroller_fortigate
FortiGate controller configuration.

## Example Usage

```hcl
resource "fmgdevice_extensioncontroller_fortigate" "trname" {
  authorized  = "disable"
  description = "your own value"
  device_id   = 10
  hostname    = "your own value"
  fosid       = "your own value"
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `authorized` - Enable/disable FortiGate administration. Valid values: `disable`, `enable`, `discovered`.

* `description` - Description.
* `device_id` - Device ID.
* `hostname` - FortiGate hostname.
* `fosid` - FortiGate serial number.
* `name` - FortiGate entry name.
* `profile` - FortiGate profile configuration.
* `vdom` - Vdom.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtensionController Fortigate can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extensioncontroller_fortigate.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

