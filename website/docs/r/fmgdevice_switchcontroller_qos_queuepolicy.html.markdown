---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_qos_queuepolicy"
description: |-
  Configure FortiSwitch QoS egress queue policy.
---

# fmgdevice_switchcontroller_qos_queuepolicy
Configure FortiSwitch QoS egress queue policy.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `cos_queue`: `fmgdevice_switchcontroller_qos_queuepolicy_cosqueue`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_qos_queuepolicy" "trname" {
  cos_queue {
    description      = "your own value"
    drop_policy      = "weighted-random-early-detection"
    ecn              = "enable"
    max_rate         = 10
    max_rate_percent = 10
    min_rate         = 10
    min_rate_percent = 10
    name             = "your own value"
    weight           = 10
  }

  name        = "your own value"
  rate_by     = "kbps"
  schedule    = "weighted"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `cos_queue` - Cos-Queue. The structure of `cos_queue` block is documented below.
* `name` - QoS policy name.
* `rate_by` - COS queue rate by kbps or percent. Valid values: `kbps`, `percent`.

* `schedule` - COS queue scheduling. Valid values: `strict`, `round-robin`, `weighted`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `cos_queue` block supports:

* `description` - Description of the COS queue.
* `drop_policy` - COS queue drop policy. Valid values: `taildrop`, `weighted-random-early-detection`.

* `ecn` - Enable/disable ECN packet marking to drop eligible packets. Valid values: `disable`, `enable`.

* `max_rate` - Maximum rate (0 - 4294967295 kbps, 0 to disable).
* `max_rate_percent` - Maximum rate (% of link speed).
* `min_rate` - Minimum rate (0 - 4294967295 kbps, 0 to disable).
* `min_rate_percent` - Minimum rate (% of link speed).
* `name` - Cos queue ID.
* `weight` - Weight of weighted round robin scheduling.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController QosQueuePolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_qos_queuepolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

