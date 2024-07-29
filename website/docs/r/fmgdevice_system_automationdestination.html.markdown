---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_automationdestination"
description: |-
  Automation destinations.
---

# fmgdevice_system_automationdestination
Automation destinations.

## Example Usage

```hcl
resource "fmgdevice_system_automationdestination" "trname" {
  destination = ["your own value"]
  ha_group_id = 10
  name        = "your own value"
  type        = "fortigate"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `destination` - Destinations.
* `ha_group_id` - Cluster group ID set for this destination (default = 0).
* `name` - Name.
* `type` - Destination type. Valid values: `fortigate`, `ha-cluster`, `fortiproxy`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationDestination can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_automationdestination.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

