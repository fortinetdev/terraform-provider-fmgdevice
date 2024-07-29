---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_vdom"
description: |-
  Configure virtual domain.
---

# fmgdevice_system_vdom
Configure virtual domain.

## Example Usage

```hcl
resource "fmgdevice_system_vdom" "trname" {
  flag        = 10
  name        = "your own value"
  short_name  = "your own value"
  vcluster_id = 10
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `flag` - Flag.
* `name` - VDOM name.
* `short_name` - VDOM short name.
* `vcluster_id` - Virtual cluster ID (0 - 4294967295).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Vdom can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_vdom.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

