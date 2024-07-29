---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_5gmodem_modem1"
description: |-
  Configure 5G Modem1.
---

# fmgdevice_system_5gmodem_modem1
Configure 5G Modem1.

~> This resource is a sub resource for variable `modem1` of resource `fmgdevice_system_5gmodem`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `sim_switch`: `fmgdevice_system_5gmodem_modem1_simswitch`



## Example Usage

```hcl
resource "fmgdevice_system_5gmodem_modem1" "trname" {
  carrier_config      = "auto-gcf"
  custom_ipv4_netmask = "your own value"
  default_gateway     = "none"
  default_netmask     = "auto"
  gps_service         = "enable"
  device_name         = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `carrier_config` - carrier-config selection mode. Valid values: `manual`, `auto-gcf`, `auto-ptcrb`.

* `custom_ipv4_netmask` - Netmask assigned by the DHCP server.
* `default_gateway` - Modem interface default gateway. Valid values: `auto`, `none`.

* `default_netmask` - Modem interface default netmask. Valid values: `auto`, `custom`.

* `gps_service` - Enable/disable Modem online mode. Valid values: `disable`, `enable`.

* `intferface` - Modem interface.
* `modem_id` - Modem ID.
* `network_type` - Modem network type. Valid values: `auto`, `4g|5g`, `3g|4g`, `3g|5g`, `5g`, `4g`, `3g`.

* `sim_data_plan` - data plan for SIM.
* `sim_pin` - PIN code for SIM(if applicable).
* `sim_switch` - Sim-Switch. The structure of `sim_switch` block is documented below.
* `sim1_data_plan` - data plan for SIM #1.
* `sim1_pin` - PIN code for SIM #1 (if applicable).
* `sim2_data_plan` - data plan for SIM #2.
* `sim2_pin` - PIN code for SIM #2 (if applicable).
* `status` - Enable/disable Modem online mode. Valid values: `online`, `low-power`.


The `sim_switch` block supports:

* `active_sim_slot` - set active SIM card slot to slot-1 or slot-2. Valid values: `slot-1`, `slot-2`.

* `by_connection_state` - Enable/disable automatic switch of SIM by MODEM connection state (mobile data usage charges not incurred). Valid values: `disable`, `enable`.

* `by_link_monitor` - Enable/disable automatic switch of SIM by link monitor (mobile data usage charges incurred). Valid values: `disable`, `enable`.

* `by_sim_state` - Enable/disable automatic switch of SIM when MODEM SIM state is empty or in error. Valid values: `disable`, `enable`.

* `link_monitor` - Set link monitor name.
* `modem_disconnection_time` - Configure connection-based automatic switch of SIM time interval in seconds (30 - 86400).
* `sim_switch_log_alert_interval` - When sim-switch &gt; X threshold within Y interval, trigger log event (Y:1 - 99999 min, default = 15).
* `sim_switch_log_alert_threshold` - When sim-switch &gt; X threshold within Y interval, trigger log event (X:1 - 1000 times, default = 5).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System 5GModemModem1 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_5gmodem_modem1.labelname System5GModemModem1
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

