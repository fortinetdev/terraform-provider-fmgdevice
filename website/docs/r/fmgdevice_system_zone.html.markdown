---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_zone"
description: |-
  Configure zones to group two or more interfaces. When a zone is created you can configure policies for the zone instead of individual interfaces in the zone.
---

# fmgdevice_system_zone
Configure zones to group two or more interfaces. When a zone is created you can configure policies for the zone instead of individual interfaces in the zone.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `tagging`: `fmgdevice_system_zone_tagging`



## Example Usage

```hcl
resource "fmgdevice_system_zone" "trname" {
  description = "your own value"
  interface   = ["port2"]
  intrazone   = "allow"
  name        = "your own value"
  tagging {
    category = ["your own value"]
    name     = "your own value"
    tags     = ["your own value"]
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `description` - Description.
* `interface` - Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined.
* `intrazone` - Allow or deny traffic routing between different interfaces in the same zone (default = deny). Valid values: `allow`, `deny`.

* `name` - Zone name.
* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `tagging` block supports:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Zone can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_zone.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

