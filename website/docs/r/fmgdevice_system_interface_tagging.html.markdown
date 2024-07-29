---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_interface_tagging"
description: |-
  Config object tagging.
---

# fmgdevice_system_interface_tagging
Config object tagging.

~> This resource is a sub resource for variable `tagging` of resource `fmgdevice_system_interface`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_interface_tagging" "trname" {
  category    = ["your own value"]
  name        = "your own value"
  tags        = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `interface` - Interface.

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System InterfaceTagging can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "interface=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_interface_tagging.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

