---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_fssopolling"
description: |-
  Configure Fortinet Single Sign On (FSSO) server.
---

# fmgdevice_system_fssopolling
Configure Fortinet Single Sign On (FSSO) server.

## Example Usage

```hcl
resource "fmgdevice_system_fssopolling" "trname" {
  _gui_meta      = "your own value"
  auth_password  = ["your own value"]
  authentication = "enable"
  listening_port = 10
  status         = "disable"
  device_name    = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `_gui_meta` - _Gui_Meta.
* `auth_password` - Password to connect to FSSO Agent.
* `authentication` - Enable/disable FSSO Agent Authentication. Valid values: `disable`, `enable`.

* `listening_port` - Listening port to accept clients (1 - 65535).
* `status` - Enable/disable FSSO Polling Mode. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FssoPolling can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_fssopolling.labelname SystemFssoPolling
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

