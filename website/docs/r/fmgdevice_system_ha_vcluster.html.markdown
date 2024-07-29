---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ha_vcluster"
description: |-
  Virtual cluster table.
---

# fmgdevice_system_ha_vcluster
Virtual cluster table.

~> This resource is a sub resource for variable `vcluster` of resource `fmgdevice_system_ha`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_ha_vcluster" "trname" {
  monitor                       = ["your own value"]
  override                      = "enable"
  override_wait_time            = 10
  pingserver_failover_threshold = 10
  pingserver_flip_timeout       = 10
  vcluster_id                   = 10
  device_name                   = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `monitor` - Interfaces to check for port monitoring (or link failure).
* `override` - Enable and increase the priority of the unit that should always be primary (master). Valid values: `disable`, `enable`.

* `override_wait_time` - Delay negotiating if override is enabled (0 - 3600 sec). Reduces how often the cluster negotiates.
* `pingserver_failover_threshold` - Remote IP monitoring failover threshold (0 - 50).
* `pingserver_flip_timeout` - Time to wait in minutes before renegotiating after a remote IP monitoring failover.
* `pingserver_monitor_interface` - Interfaces to check for remote IP monitoring.
* `pingserver_secondary_force_reset` - Enable to force the cluster to negotiate after a remote IP monitoring failover. Valid values: `disable`, `enable`.

* `pingserver_slave_force_reset` - Pingserver-Slave-Force-Reset. Valid values: `disable`, `enable`.

* `priority` - Increase the priority to select the primary unit (0 - 255).
* `vcluster_id` - ID.
* `vdom` - Virtual domain(s) in the virtual cluster.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{vcluster_id}}.

## Import

System HaVcluster can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ha_vcluster.labelname {{vcluster_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

