---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extendercontroller_extender"
description: |-
  Device ExtenderControllerExtender
---

# fmgdevice_extendercontroller_extender
Device ExtenderControllerExtender

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `controller_report`: `fmgdevice_extendercontroller_extender_controllerreport`
>- `modem1`: `fmgdevice_extendercontroller_extender_modem1`
>- `modem2`: `fmgdevice_extendercontroller_extender_modem2`
>- `wan_extension`: `fmgdevice_extendercontroller_extender_wanextension`



## Example Usage

```hcl
resource "fmgdevice_extendercontroller_extender" "trname" {
  _dataplan         = ["your own value"]
  _template         = ["your own value"]
  aaa_shared_secret = ["your own value"]
  access_point_name = "your own value"
  admin             = "disable"
  name              = "your own value"
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `_dataplan` - _Dataplan.
* `_template` - _Template.
* `aaa_shared_secret` - Aaa-Shared-Secret.
* `access_point_name` - Access-Point-Name.
* `admin` - Admin. Valid values: `disable`, `enable`, `discovered`.

* `at_dial_script` - At-Dial-Script.
* `allowaccess` - Allowaccess. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`.

* `authorized` - Authorized. Valid values: `disable`, `enable`.

* `billing_start_day` - Billing-Start-Day.
* `cdma_aaa_spi` - Cdma-Aaa-Spi.
* `cdma_ha_spi` - Cdma-Ha-Spi.
* `cdma_nai` - Cdma-Nai.
* `conn_status` - Conn-Status.
* `bandwidth_limit` - Bandwidth-Limit.
* `controller_report` - Controller-Report. The structure of `controller_report` block is documented below.
* `description` - Description.
* `dial_mode` - Dial-Mode. Valid values: `dial-on-demand`, `always-connect`.

* `dial_status` - Dial-Status.
* `device_id` - Device-Id.
* `enforce_bandwidth` - Enforce-Bandwidth. Valid values: `disable`, `enable`.

* `ext_name` - Ext-Name.
* `ha_shared_secret` - Ha-Shared-Secret.
* `extension_type` - Extension-Type. Valid values: `wan-extension`, `lan-extension`.

* `fosid` - Id.
* `ifname` - Ifname.
* `initiated_update` - Initiated-Update. Valid values: `disable`, `enable`.

* `login_password` - Login-Password.
* `mode` - Mode. Valid values: `standalone`, `redundant`.

* `modem_passwd` - Modem-Passwd.
* `modem_type` - Modem-Type. Valid values: `cdma`, `gsm/lte`, `wimax`.

* `login_password_change` - Login-Password-Change. Valid values: `no`, `yes`, `default`.

* `modem1` - Modem1. The structure of `modem1` block is documented below.
* `modem2` - Modem2. The structure of `modem2` block is documented below.
* `multi_mode` - Multi-Mode. Valid values: `auto`, `auto-3g`, `force-lte`, `force-3g`, `force-2g`.

* `name` - Name.
* `ppp_auth_protocol` - Ppp-Auth-Protocol. Valid values: `auto`, `pap`, `chap`.

* `ppp_echo_request` - Ppp-Echo-Request. Valid values: `disable`, `enable`.

* `ppp_password` - Ppp-Password.
* `ppp_username` - Ppp-Username.
* `primary_ha` - Primary-Ha.
* `quota_limit_mb` - Quota-Limit-Mb.
* `redial` - Redial. Valid values: `none`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`.

* `redundant_intf` - Redundant-Intf.
* `roaming` - Roaming. Valid values: `disable`, `enable`.

* `role` - Role. Valid values: `none`, `primary`, `secondary`.

* `secondary_ha` - Secondary-Ha.
* `sim_pin` - Sim-Pin.
* `override_allowaccess` - Override-Allowaccess. Valid values: `disable`, `enable`.

* `override_enforce_bandwidth` - Override-Enforce-Bandwidth. Valid values: `disable`, `enable`.

* `override_login_password_change` - Override-Login-Password-Change. Valid values: `disable`, `enable`.

* `profile` - Profile.
* `vdom` - Vdom.
* `wimax_auth_protocol` - Wimax-Auth-Protocol. Valid values: `tls`, `ttls`.

* `wimax_carrier` - Wimax-Carrier.
* `wimax_realm` - Wimax-Realm.
* `wan_extension` - Wan-Extension. The structure of `wan_extension` block is documented below.

The `controller_report` block supports:

* `interval` - Controller report interval.
* `signal_threshold` - Controller report signal threshold.
* `status` - FortiExtender controller report status. Valid values: `disable`, `enable`.


The `modem1` block supports:

* `_sim_profile` - _Sim_Profile.
* `auto_switch` - Auto-Switch. The structure of `auto_switch` block is documented below.
* `conn_status` - Conn-Status.
* `default_sim` - Default SIM selection. Valid values: `sim1`, `sim2`, `carrier`, `cost`.

* `gps` - FortiExtender GPS enable/disable. Valid values: `disable`, `enable`.

* `ifname` - FortiExtender interface name.
* `modem_id` - Modem-Id.
* `preferred_carrier` - Preferred carrier.
* `redundant_intf` - Redundant interface.
* `redundant_mode` - FortiExtender mode. Valid values: `disable`, `enable`.

* `sim1_pin` - SIM #1 PIN status. Valid values: `disable`, `enable`.

* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin` - SIM #2 PIN status. Valid values: `disable`, `enable`.

* `sim2_pin_code` - SIM #2 PIN password.
* `status` - FortiExtender modem status. Valid values: `disable`, `enable`.


The `auto_switch` block supports:

* `dataplan` - Automatically switch based on data usage. Valid values: `disable`, `enable`.

* `disconnect` - Auto switch by disconnect. Valid values: `disable`, `enable`.

* `disconnect_period` - Automatically switch based on disconnect period.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `signal` - Automatically switch based on signal strength. Valid values: `disable`, `enable`.

* `status` - FortiExtender automatic switch status. Valid values: `disable`, `enable`.

* `switch_back` - Auto switch with switch back multi-options. Valid values: `time`, `timer`.

* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).

The `modem2` block supports:

* `_sim_profile` - _Sim_Profile.
* `auto_switch` - Auto-Switch. The structure of `auto_switch` block is documented below.
* `conn_status` - Conn-Status.
* `default_sim` - Default SIM selection. Valid values: `sim1`, `sim2`, `carrier`, `cost`.

* `gps` - FortiExtender GPS enable/disable. Valid values: `disable`, `enable`.

* `ifname` - FortiExtender interface name.
* `modem_id` - Modem-Id.
* `preferred_carrier` - Preferred carrier.
* `redundant_intf` - Redundant interface.
* `redundant_mode` - FortiExtender mode. Valid values: `disable`, `enable`.

* `sim1_pin` - SIM #1 PIN status. Valid values: `disable`, `enable`.

* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin` - SIM #2 PIN status. Valid values: `disable`, `enable`.

* `sim2_pin_code` - SIM #2 PIN password.
* `status` - FortiExtender modem status. Valid values: `disable`, `enable`.


The `auto_switch` block supports:

* `dataplan` - Automatically switch based on data usage. Valid values: `disable`, `enable`.

* `disconnect` - Auto switch by disconnect. Valid values: `disable`, `enable`.

* `disconnect_period` - Automatically switch based on disconnect period.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `signal` - Automatically switch based on signal strength. Valid values: `disable`, `enable`.

* `status` - FortiExtender automatic switch status. Valid values: `disable`, `enable`.

* `switch_back` - Auto switch with switch back multi-options. Valid values: `time`, `timer`.

* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).

The `wan_extension` block supports:

* `modem1_extension` - Modem1-Extension.
* `modem2_extension` - Modem2-Extension.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtenderController Extender can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extendercontroller_extender.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

