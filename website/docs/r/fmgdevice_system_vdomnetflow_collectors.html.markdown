---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_vdomnetflow_collectors"
description: |-
  Netflow collectors.
---

# fmgdevice_system_vdomnetflow_collectors
Netflow collectors.

~> This resource is a sub resource for variable `collectors` of resource `fmgdevice_system_vdomnetflow`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_vdomnetflow_collectors" "trname" {
  collector_ip            = "your own value"
  collector_port          = 10
  fosid                   = 10
  interface               = ["port2"]
  interface_select_method = "sdwan"
  device_name             = var.device_name # not required if setting is at provider
  device_vdom             = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `collector_ip` - Collector IP.
* `collector_port` - NetFlow collector port number.
* `fosid` - ID.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `source_ip` - Source IP address for communication with the NetFlow agent.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System VdomNetflowCollectors can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_vdomnetflow_collectors.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

