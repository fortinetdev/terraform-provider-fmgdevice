---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_federatedupgrade_nodelist"
description: |-
  Nodes which will be included in the upgrade.
---

# fmgdevice_system_federatedupgrade_nodelist
Nodes which will be included in the upgrade.

~> This resource is a sub resource for variable `node_list` of resource `fmgdevice_system_federatedupgrade`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_federatedupgrade_nodelist" "trname" {
  coordinating_fortigate = "your own value"
  device_type            = "fortigate"
  maximum_minutes        = 10
  serial                 = "your own value"
  setup_time             = ["your own value"]
  device_name            = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `coordinating_fortigate` - Serial number of the FortiGate unit that controls this device.
* `device_type` - Fortinet device type. Valid values: `fortigate`, `fortiswitch`, `fortiap`, `fortiproxy`, `fortiextender`.

* `maximum_minutes` - Maximum number of minutes to allow for immediate upgrade preparation.
* `serial` - Serial number of the node to include.
* `setup_time` - Upgrade preparation start time in UTC (hh:mm yyyy/mm/dd UTC).
* `time` - Scheduled upgrade execution time in UTC (hh:mm yyyy/mm/dd UTC).
* `timing` - Run immediately or at a scheduled time. Valid values: `immediate`, `scheduled`.

* `upgrade_path` - Fortinet OS image versions to upgrade through in major-minor-patch format, such as 7-0-4.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{serial}}.

## Import

System FederatedUpgradeNodeList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_federatedupgrade_nodelist.labelname {{serial}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

