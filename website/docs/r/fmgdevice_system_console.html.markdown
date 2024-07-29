---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_console"
description: |-
  Configure console.
---

# fmgdevice_system_console
Configure console.

## Example Usage

```hcl
resource "fmgdevice_system_console" "trname" {
  baudrate      = "57600"
  fortiexplorer = "disable"
  login         = "disable"
  mode          = "batch"
  output        = "more"
  device_name   = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `baudrate` - Baudrate. Valid values: `9600`, `19200`, `38400`, `57600`, `115200`.

* `fortiexplorer` - Enable/disable access for FortiExplorer. Valid values: `disable`, `enable`.

* `login` - Enable/disable serial console and FortiExplorer. Valid values: `disable`, `enable`.

* `mode` - Console mode. Valid values: `line`, `batch`.

* `output` - Console output mode. Valid values: `standard`, `more`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Console can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_console.labelname SystemConsole
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

