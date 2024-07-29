---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_automationaction_httpheaders"
description: |-
  Request headers.
---

# fmgdevice_system_automationaction_httpheaders
Request headers.

~> This resource is a sub resource for variable `http_headers` of resource `fmgdevice_system_automationaction`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_automationaction_httpheaders" "trname" {
  fosid       = 10
  key         = "your own value"
  value       = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `automation_action` - Automation Action.

* `fosid` - Entry ID.
* `key` - Request header key.
* `value` - Request header value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System AutomationActionHttpHeaders can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "automation_action=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_automationaction_httpheaders.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

