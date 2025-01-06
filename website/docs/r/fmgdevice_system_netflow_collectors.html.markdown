---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_netflow_collectors"
description: |-
  Netflow collectors.
---

# fmgdevice_system_netflow_collectors
Netflow collectors.

~> This resource is a sub resource for variable `collectors` of resource `fmgdevice_system_netflow`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_netflow_collectors" "trname" {
  collector_ip            = "your own value"
  collector_port          = 10
  fosid                   = 10
  interface               = ["port2"]
  interface_select_method = "specify"
  device_name             = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `collector_ip` - Collector IP.
* `collector_port` - NetFlow collector port number.
* `fosid` - ID.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `source_ip` - Source IP address for communication with the NetFlow agent.
* `source_ip_interface` - Name of the interface used to determine the source IP for exporting packets.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System NetflowCollectors can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_netflow_collectors.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

