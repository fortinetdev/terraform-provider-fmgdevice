---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_federatedupgrade"
description: |-
  Coordinate federated upgrades within the Security Fabric.
---

# fmgdevice_system_federatedupgrade
Coordinate federated upgrades within the Security Fabric.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `known_ha_members`: `fmgdevice_system_federatedupgrade_knownhamembers`
>- `node_list`: `fmgdevice_system_federatedupgrade_nodelist`



## Example Usage

```hcl
resource "fmgdevice_system_federatedupgrade" "trname" {
  failure_device        = "your own value"
  failure_reason        = "internal"
  ha_reboot_controller  = "your own value"
  ignore_signing_errors = "enable"
  known_ha_members {
    serial = "your own value"
  }

  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `dry_run` - Perform federated upgrades as a dry run. Valid values: `disable`, `enable`.

* `failure_device` - Serial number of the node to include.
* `failure_reason` - Reason for upgrade failure. Valid values: `none`, `internal`, `timeout`, `device-type-unsupported`, `download-failed`, `device-missing`, `version-unavailable`, `staging-failed`, `reboot-failed`, `device-not-reconnected`, `node-not-ready`, `no-final-confirmation`, `no-confirmation-query`, `config-error-log-nonempty`, `node-failed`, `csf-tree-not-supported`.

* `ha_reboot_controller` - Serial number of the FortiGate unit that will control the reboot process for the federated upgrade of the HA cluster.
* `ignore_signing_errors` - Allow/reject use of FortiGate firmware images that are unsigned. Valid values: `disable`, `enable`.

* `known_ha_members` - Known-Ha-Members. The structure of `known_ha_members` block is documented below.
* `next_path_index` - The index of the next image to upgrade to.
* `node_list` - Node-List. The structure of `node_list` block is documented below.
* `source` - Source that set up the federated upgrade config. Valid values: `user`, `auto-firmware-upgrade`.

* `status` - Current status of the upgrade. Valid values: `disabled`, `initialized`, `downloading`, `download-failed`, `ready`, `cancelled`, `confirmed`, `done`, `failed`, `device-disconnected`, `staging`, `final-check`, `upgrade-devices`, `coordinating`.

* `upgrade_id` - Unique identifier for this upgrade.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `known_ha_members` block supports:

* `serial` - Serial number of HA member

The `node_list` block supports:

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
* `id` - an identifier for the resource.

## Import

System FederatedUpgrade can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_federatedupgrade.labelname SystemFederatedUpgrade
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

