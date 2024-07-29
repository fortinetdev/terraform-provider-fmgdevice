---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_centralmanagement"
description: |-
  Configure central management.
---

# fmgdevice_system_centralmanagement
Configure central management.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `server_list`: `fmgdevice_system_centralmanagement_serverlist`



## Example Usage

```hcl
resource "fmgdevice_system_centralmanagement" "trname" {
  allow_monitor                     = "disable"
  allow_push_configuration          = "enable"
  allow_push_firmware               = "enable"
  allow_remote_firmware_upgrade     = "enable"
  allow_remote_lte_firmware_upgrade = "enable"
  device_name                       = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `allow_monitor` - Enable/disable allowing the central management server to remotely monitor this FortiGate unit. Valid values: `disable`, `enable`.

* `allow_push_configuration` - Enable/disable allowing the central management server to push configuration changes to this FortiGate. Valid values: `disable`, `enable`.

* `allow_push_firmware` - Enable/disable allowing the central management server to push firmware updates to this FortiGate. Valid values: `disable`, `enable`.

* `allow_remote_firmware_upgrade` - Enable/disable remotely upgrading the firmware on this FortiGate from the central management server. Valid values: `disable`, `enable`.

* `allow_remote_lte_firmware_upgrade` - Enable/disable remotely upgrading the lte firmware on this FortiGate from the central management server. Valid values: `disable`, `enable`.

* `allow_remote_modem_firmware_upgrade` - Enable/disable remotely upgrading the internal cellular modem firmware on this FortiGate from the central management server. Valid values: `disable`, `enable`.

* `ca_cert` - CA certificate to be used by FGFM protocol.
* `enc_algorithm` - Encryption strength for communications between the FortiGate and central management. Valid values: `default`, `high`, `low`.

* `fmg` - IP address or FQDN of the FortiManager.
* `fmg_source_ip` - IPv4 source address that this FortiGate uses when communicating with FortiManager.
* `fmg_source_ip6` - IPv6 source address that this FortiGate uses when communicating with FortiManager.
* `fmg_update_port` - Port used to communicate with FortiManager that is acting as a FortiGuard update server. Valid values: `443`, `8890`.

* `fortigate_cloud_sso_default_profile` - Override access profile.
* `include_default_servers` - Enable/disable inclusion of public FortiGuard servers in the override server list. Valid values: `disable`, `enable`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `local_cert` - Certificate to be used by FGFM protocol.
* `ltefw_upgrade_frequency` - Set LTE firmware auto pushdown frequency. Valid values: `everyHour`, `every12hour`, `everyDay`, `everyWeek`.

* `ltefw_upgrade_time` - Schedule next LTE firmware upgrade time (Local Time). Format: YYYY-MM-DD HH:MM:SS
* `mode` - Central management mode. Valid values: `normal`, `backup`.

* `modem_upgrade_frequency` - Set internal cellular modem firmware auto pushdown frequency. Valid values: `everyHour`, `every12hour`, `everyDay`, `everyWeek`.

* `modem_upgrade_time` - Schedule next internal cellular modem firmware upgrade time (Local Time). Format: YYYY-MM-DD HH:MM:SS
* `schedule_config_restore` - Enable/disable allowing the central management server to restore the configuration of this FortiGate. Valid values: `disable`, `enable`.

* `schedule_script_restore` - Enable/disable allowing the central management server to restore the scripts stored on this FortiGate. Valid values: `disable`, `enable`.

* `serial_number` - Serial number.
* `server_list` - Server-List. The structure of `server_list` block is documented below.
* `type` - Central management type. Valid values: `fortimanager`, `fortiguard`, `none`.

* `use_elbc_vdom` - Enable/disable use of special ELBC config sync VDOM to connect to FortiManager. Valid values: `disable`, `enable`.

* `vdom` - Virtual domain (VDOM) name to use when communicating with FortiManager.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `server_list` block supports:

* `addr_type` - Indicate whether the FortiGate communicates with the override server using an IPv4 address, an IPv6 address or a FQDN. Valid values: `fqdn`, `ipv4`, `ipv6`.

* `fqdn` - FQDN address of override server.
* `id` - ID.
* `server_address` - IPv4 address of override server.
* `server_address6` - IPv6 address of override server.
* `server_type` - FortiGuard service type. Valid values: `update`, `rating`, `iot-query`, `iot-collect`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System CentralManagement can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_centralmanagement.labelname SystemCentralManagement
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

