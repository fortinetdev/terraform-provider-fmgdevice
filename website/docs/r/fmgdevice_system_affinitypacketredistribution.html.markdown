---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_affinitypacketredistribution"
description: |-
  Configure packet redistribution.
---

# fmgdevice_system_affinitypacketredistribution
Configure packet redistribution.

## Example Usage

```hcl
resource "fmgdevice_system_affinitypacketredistribution" "trname" {
  affinity_cpumask = "your own value"
  fosid            = 10
  interface        = ["port2"]
  round_robin      = "disable"
  rxqid            = 10
  device_name      = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `affinity_cpumask` - Affinity setting for VM throughput (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
* `fosid` - ID of the packet redistribution setting.
* `interface` - Physical interface name on which to perform packet redistribution.
* `round_robin` - Enable/disable round-robin redistribution to multiple CPUs. Valid values: `disable`, `enable`.

* `rxqid` - ID of the receive queue (when the interface has multiple queues) on which to perform packet redistribution (255 = all queues).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System AffinityPacketRedistribution can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_affinitypacketredistribution.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

