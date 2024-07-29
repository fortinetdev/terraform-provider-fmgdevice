---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_quarantine_targets_macs"
description: |-
  Quarantine MACs.
---

# fmgdevice_user_quarantine_targets_macs
Quarantine MACs.

~> This resource is a sub resource for variable `macs` of resource `fmgdevice_user_quarantine_targets`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_user_quarantine_targets_macs" "trname" {
  description = "your own value"
  drop        = "enable"
  entry_id    = 10
  mac         = "your own value"
  parent      = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `targets` - Targets.

* `description` - Description for the quarantine MAC.
* `drop` - Enable/disable dropping of quarantined device traffic. Valid values: `disable`, `enable`.

* `entry_id` - FSW entry id for the quarantine MAC.
* `mac` - Quarantine MAC.
* `parent` - Parent.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mac}}.

## Import

User QuarantineTargetsMacs can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "targets=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_quarantine_targets_macs.labelname {{mac}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

