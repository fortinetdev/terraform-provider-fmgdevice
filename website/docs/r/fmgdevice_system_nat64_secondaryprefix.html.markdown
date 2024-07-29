---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_nat64_secondaryprefix"
description: |-
  Device SystemNat64SecondaryPrefix
---

# fmgdevice_system_nat64_secondaryprefix
Device SystemNat64SecondaryPrefix

~> This resource is a sub resource for variable `secondary_prefix` of resource `fmgdevice_system_nat64`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_nat64_secondaryprefix" "trname" {
  name         = "your own value"
  nat64_prefix = "your own value"
  device_name  = var.device_name # not required if setting is at provider
  device_vdom  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `name` - Name.
* `nat64_prefix` - Nat64-Prefix.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Nat64SecondaryPrefix can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_nat64_secondaryprefix.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

