---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_affinityinterrupt"
description: |-
  Configure interrupt affinity.
---

# fmgdevice_system_affinityinterrupt
Configure interrupt affinity.

## Example Usage

```hcl
resource "fmgdevice_system_affinityinterrupt" "trname" {
  affinity_cpumask         = "your own value"
  default_affinity_cpumask = "your own value"
  fosid                    = 10
  interrupt                = "your own value"
  device_name              = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `affinity_cpumask` - Affinity setting (64-bit hexadecimal value in the format of 0xxxxxxxxxxxxxxxxx).
* `default_affinity_cpumask` - Default-Affinity-Cpumask.
* `fosid` - ID of the interrupt affinity setting.
* `interrupt` - Interrupt name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System AffinityInterrupt can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_affinityinterrupt.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

