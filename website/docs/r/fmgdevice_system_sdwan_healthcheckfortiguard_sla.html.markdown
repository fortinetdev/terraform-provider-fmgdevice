---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sdwan_healthcheckfortiguard_sla"
description: |-
  Service level agreement (SLA).
---

# fmgdevice_system_sdwan_healthcheckfortiguard_sla
Service level agreement (SLA).

~> This resource is a sub resource for variable `sla` of resource `fmgdevice_system_sdwan_healthcheckfortiguard`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `health_check_fortiguard` - Health Check Fortiguard.

* `fosid` - SLA ID.
* `jitter_threshold` - Jitter for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `latency_threshold` - Latency for SLA to make decision in milliseconds. (0 - 10000000, default = 5).
* `link_cost_factor` - Criteria on which to base link selection. Valid values: `latency`, `jitter`, `packet-loss`, `mos`, `remote`.

* `mos_threshold` - Minimum Mean Opinion Score for SLA to be marked as pass. (1.0 - 5.0, default = 3.6).
* `packetloss_threshold` - Packet loss for SLA to make decision in percentage. (0 - 100, default = 0).
* `priority_in_sla` - Value to be distributed into routing table when in-sla (0 - 65535, default = 0).
* `priority_out_sla` - Value to be distributed into routing table when out-sla (0 - 65535, default = 0).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SdwanHealthCheckFortiguardSla can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "health_check_fortiguard=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sdwan_healthcheckfortiguard_sla.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

