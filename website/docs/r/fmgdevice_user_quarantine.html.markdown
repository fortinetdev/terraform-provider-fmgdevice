---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_quarantine"
description: |-
  Configure quarantine support.
---

# fmgdevice_user_quarantine
Configure quarantine support.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `targets`: `fmgdevice_user_quarantine_targets`



## Example Usage

```hcl
resource "fmgdevice_user_quarantine" "trname" {
  firewall_groups = ["your own value"]
  quarantine      = "enable"
  targets {
    description = "your own value"
    entry       = "your own value"
    macs {
      description = "your own value"
      drop        = "enable"
      entry_id    = 10
      mac         = "your own value"
      parent      = "your own value"
    }

  }

  traffic_policy = ["your own value"]
  device_name    = var.device_name # not required if setting is at provider
  device_vdom    = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `firewall_groups` - Firewall address group which includes all quarantine MAC address.
* `quarantine` - Enable/disable quarantine. Valid values: `disable`, `enable`.

* `targets` - Targets. The structure of `targets` block is documented below.
* `traffic_policy` - Traffic policy for quarantined MACs.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `targets` block supports:

* `description` - Description for the quarantine entry.
* `entry` - Quarantine entry name.
* `macs` - Macs. The structure of `macs` block is documented below.

The `macs` block supports:

* `description` - Description for the quarantine MAC.
* `drop` - Enable/disable dropping of quarantined device traffic. Valid values: `disable`, `enable`.

* `entry_id` - FSW entry id for the quarantine MAC.
* `mac` - Quarantine MAC.
* `parent` - Parent.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

User Quarantine can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_quarantine.labelname UserQuarantine
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

