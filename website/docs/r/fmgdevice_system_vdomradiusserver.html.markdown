---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_vdomradiusserver"
description: |-
  Configure a RADIUS server to use as a RADIUS Single Sign On (RSSO) server for this VDOM.
---

# fmgdevice_system_vdomradiusserver
Configure a RADIUS server to use as a RADIUS Single Sign On (RSSO) server for this VDOM.

## Example Usage

```hcl
resource "fmgdevice_system_vdomradiusserver" "trname" {
  name               = "your own value"
  radius_server_vdom = ["your own value"]
  status             = "disable"
  device_name        = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `name` - Name of the VDOM that you are adding the RADIUS server to.
* `radius_server_vdom` - Use this option to select another VDOM containing a VDOM RSSO RADIUS server to use for the current VDOM.
* `status` - Enable/disable the RSSO RADIUS server for this VDOM. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System VdomRadiusServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_vdomradiusserver.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

