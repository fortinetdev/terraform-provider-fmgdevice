---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_autoscript"
description: |-
  Configure auto script.
---

# fmgdevice_system_autoscript
Configure auto script.

## Example Usage

```hcl
resource "fmgdevice_system_autoscript" "trname" {
  interval    = 10
  name        = "your own value"
  output_size = 10
  repeat      = 10
  script      = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `interval` - Repeat interval in seconds.
* `name` - Auto script name.
* `output_size` - Number of megabytes to limit script output to (10 - 1024, default = 10).
* `repeat` - Number of times to repeat this script (0 = infinite).
* `script` - List of FortiOS CLI commands to repeat.
* `start` - Script starting mode. Valid values: `auto`, `manual`.

* `timeout` - Maximum running time for this script in seconds (0 = no timeout).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutoScript can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_autoscript.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

