---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_automationtrigger_fields"
description: |-
  Customized trigger field settings.
---

# fmgdevice_system_automationtrigger_fields
Customized trigger field settings.

~> This resource is a sub resource for variable `fields` of resource `fmgdevice_system_automationtrigger`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_automationtrigger_fields" "trname" {
  fosid       = 10
  name        = "your own value"
  value       = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `automation_trigger` - Automation Trigger.

* `fosid` - Entry ID.
* `name` - Name.
* `value` - Value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System AutomationTriggerFields can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "automation_trigger=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_automationtrigger_fields.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

