---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_quarantine_targets"
description: |-
  Quarantine entry to hold multiple MACs.
---

# fmgdevice_user_quarantine_targets
Quarantine entry to hold multiple MACs.

~> This resource is a sub resource for variable `targets` of resource `fmgdevice_user_quarantine`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `macs`: `fmgdevice_user_quarantine_targets_macs`



## Example Usage

```hcl
resource "fmgdevice_user_quarantine_targets" "trname" {
  description = "your own value"
  entry       = "your own value"
  macs {
    description = "your own value"
    drop        = "disable"
    entry_id    = 10
    mac         = "your own value"
    parent      = "your own value"
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `description` - Description for the quarantine entry.
* `entry` - Quarantine entry name.
* `macs` - Macs. The structure of `macs` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `macs` block supports:

* `description` - Description for the quarantine MAC.
* `drop` - Enable/disable dropping of quarantined device traffic. Valid values: `disable`, `enable`.

* `entry_id` - FSW entry id for the quarantine MAC.
* `mac` - Quarantine MAC.
* `parent` - Parent.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{entry}}.

## Import

User QuarantineTargets can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_quarantine_targets.labelname {{entry}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

