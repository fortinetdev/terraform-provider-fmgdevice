---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_flowtracking_collectors"
description: |-
  Configure collectors for the flow.
---

# fmgdevice_switchcontroller_flowtracking_collectors
Configure collectors for the flow.

~> This resource is a sub resource for variable `collectors` of resource `fmgdevice_switchcontroller_flowtracking`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_flowtracking_collectors" "trname" {
  ip          = "your own value"
  name        = "your own value"
  port        = 10
  transport   = "udp"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `ip` - Collector IP address.
* `name` - Collector name.
* `port` - Collector port number(0-65535, default:0, netflow:2055, ipfix:4739).
* `transport` - Collector L4 transport protocol for exporting packets. Valid values: `udp`, `tcp`, `sctp`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController FlowTrackingCollectors can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_flowtracking_collectors.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

