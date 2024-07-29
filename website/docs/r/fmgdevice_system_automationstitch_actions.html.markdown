---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_automationstitch_actions"
description: |-
  Configure stitch actions.
---

# fmgdevice_system_automationstitch_actions
Configure stitch actions.

~> This resource is a sub resource for variable `actions` of resource `fmgdevice_system_automationstitch`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_automationstitch_actions" "trname" {
  action      = ["your own value"]
  delay       = 10
  fosid       = 10
  required    = "enable"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `automation_stitch` - Automation Stitch.

* `action` - Action name.
* `delay` - Delay before execution (in seconds).
* `fosid` - Entry ID.
* `required` - Required in action chain. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System AutomationStitchActions can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "automation_stitch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_automationstitch_actions.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

