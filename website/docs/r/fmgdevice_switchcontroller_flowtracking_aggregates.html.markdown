---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_flowtracking_aggregates"
description: |-
  Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow.
---

# fmgdevice_switchcontroller_flowtracking_aggregates
Configure aggregates in which all traffic sessions matching the IP Address will be grouped into the same flow.

~> This resource is a sub resource for variable `aggregates` of resource `fmgdevice_switchcontroller_flowtracking`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_flowtracking_aggregates" "trname" {
  fosid       = 10
  ip          = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `fosid` - Aggregate id.
* `ip` - IP address to group all matching traffic sessions to a flow.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SwitchController FlowTrackingAggregates can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_flowtracking_aggregates.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

