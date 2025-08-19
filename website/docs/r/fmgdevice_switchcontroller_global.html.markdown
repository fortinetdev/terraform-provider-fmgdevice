---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_global"
description: |-
  Configure FortiSwitch global settings.
---

# fmgdevice_switchcontroller_global
Configure FortiSwitch global settings.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `custom_command`: `fmgdevice_switchcontroller_global_customcommand`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_global" "trname" {
  allow_multiple_interfaces = "disable"
  bounce_quarantined_link   = "enable"
  custom_command {
    command_entry = "your own value"
    command_name  = ["your own value"]
  }

  default_virtual_switch_vlan = ["your own value"]
  dhcp_option82_circuit_id    = ["hostname"]
  device_name                 = var.device_name # not required if setting is at provider
  device_vdom                 = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `allow_multiple_interfaces` - Allow-Multiple-Interfaces. Valid values: `disable`, `enable`.

* `bounce_quarantined_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.

* `custom_command` - Custom-Command. The structure of `custom_command` block is documented below.
* `default_virtual_switch_vlan` - Default VLAN for ports when added to the virtual-switch.
* `dhcp_option82_circuit_id` - List the parameters to be included to inform about client identification. Valid values: `intfname`, `vlan`, `hostname`, `mode`, `description`.

* `dhcp_option82_format` - DHCP option-82 format string. Valid values: `ascii`, `legacy`.

* `dhcp_option82_remote_id` - List the parameters to be included to inform about client identification. Valid values: `mac`, `hostname`, `ip`.

* `dhcp_server_access_list` - Enable/disable DHCP snooping server access list. Valid values: `disable`, `enable`.

* `dhcp_snoop_client_db_exp` - Expiry time for DHCP snooping server database entries (300 - 259200 sec, default = 86400 sec).
* `dhcp_snoop_client_req` - Client DHCP packet broadcast mode. Valid values: `drop-untrusted`, `forward-untrusted`.

* `dhcp_snoop_db_per_port_learn_limit` - Per Interface dhcp-server entries learn limit (0 - 1024, default = 64).
* `disable_discovery` - Prevent this FortiSwitch from discovering.
* `fips_enforce` - Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.

* `firmware_provision_on_authorization` - Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `disable`, `enable`.

* `https_image_push` - Enable/disable image push to FortiSwitch using HTTPS. Valid values: `disable`, `enable`.

* `log_mac_limit_violations` - Enable/disable logs for Learning Limit Violations. Valid values: `disable`, `enable`.

* `mac_aging_interval` - Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
* `mac_event_logging` - Enable/disable MAC address event logging. Valid values: `disable`, `enable`.

* `mac_retention_period` - Time in hours after which an inactive MAC is removed from client DB (0 = aged out based on mac-aging-interval).
* `mac_violation_timer` - Set timeout for Learning Limit Violations (0 = disabled).
* `quarantine_mode` - Quarantine mode. Valid values: `by-vlan`, `by-redirect`.

* `sn_dns_resolution` - Enable/disable DNS resolution of the FortiSwitch unit's IP address with switch name. Valid values: `disable`, `enable`.

* `switch_on_deauth` - No-operation/Factory-reset the managed FortiSwitch on deauthorization. Valid values: `no-op`, `factory-reset`.

* `update_user_device` - Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.

* `vlan_all_mode` - VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `defined`, `all`.

* `vlan_identity` - Identity of the VLAN. Commonly used for RADIUS Tunnel-Private-Group-Id. Valid values: `description`, `name`.

* `vlan_optimization` - FortiLink VLAN optimization. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `custom_command` block supports:

* `command_entry` - List of FortiSwitch commands.
* `command_name` - Name of custom command to push to all FortiSwitches in VDOM.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController Global can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_global.labelname SwitchControllerGlobal
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

