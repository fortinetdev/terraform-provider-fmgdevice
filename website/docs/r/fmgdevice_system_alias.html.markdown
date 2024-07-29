---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_alias"
description: |-
  Configure alias command.
---

# fmgdevice_system_alias
Configure alias command.

## Example Usage

```hcl
resource "fmgdevice_system_alias" "trname" {
  command     = "your own value"
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `command` - Command list to execute.
* `name` - Alias command name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Alias can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_alias.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

