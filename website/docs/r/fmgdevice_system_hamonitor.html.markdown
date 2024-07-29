---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_hamonitor"
description: |-
  Configure HA monitor.
---

# fmgdevice_system_hamonitor
Configure HA monitor.

## Example Usage

```hcl
resource "fmgdevice_system_hamonitor" "trname" {
  monitor_vlan           = "enable"
  vlan_hb_interval       = 10
  vlan_hb_lost_threshold = 10
  device_name            = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `monitor_vlan` - Enable/disable monitor VLAN interfaces. Valid values: `disable`, `enable`.

* `vlan_hb_interval` - Configure heartbeat interval (seconds).
* `vlan_hb_lost_threshold` - VLAN lost heartbeat threshold (1 - 60).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System HaMonitor can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_hamonitor.labelname SystemHaMonitor
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

