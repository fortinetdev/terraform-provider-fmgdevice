---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_automationstitch"
description: |-
  Automation stitches.
---

# fmgdevice_system_automationstitch
Automation stitches.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `actions`: `fmgdevice_system_automationstitch_actions`



## Example Usage

```hcl
resource "fmgdevice_system_automationstitch" "trname" {
  action = ["your own value"]
  actions {
    action   = ["your own value"]
    delay    = 10
    fosid    = 10
    required = "enable"
  }

  description = "your own value"
  destination = ["your own value"]
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `action` - Action.
* `actions` - Actions. The structure of `actions` block is documented below.
* `condition` - Automation conditions.
* `condition_logic` - Apply AND/OR logic to the specified automation conditions. Valid values: `or`, `and`.

* `description` - Description.
* `destination` - Serial number/HA group-name of destination devices.
* `name` - Name.
* `status` - Enable/disable this stitch. Valid values: `disable`, `enable`.

* `trigger` - Trigger name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `actions` block supports:

* `action` - Action name.
* `delay` - Delay before execution (in seconds).
* `id` - Entry ID.
* `required` - Required in action chain. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationStitch can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_automationstitch.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

