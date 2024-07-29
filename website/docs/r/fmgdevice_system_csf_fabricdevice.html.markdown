---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_csf_fabricdevice"
description: |-
  Fabric device configuration.
---

# fmgdevice_system_csf_fabricdevice
Fabric device configuration.

~> This resource is a sub resource for variable `fabric_device` of resource `fmgdevice_system_csf`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_csf_fabricdevice" "trname" {
  access_token = ["your own value"]
  device_ip    = "your own value"
  https_port   = 10
  name         = "your own value"
  device_name  = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `access_token` - Device access token.
* `device_ip` - Device IP.
* `https_port` - HTTPS port for fabric device.
* `name` - Device name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CsfFabricDevice can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_csf_fabricdevice.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

