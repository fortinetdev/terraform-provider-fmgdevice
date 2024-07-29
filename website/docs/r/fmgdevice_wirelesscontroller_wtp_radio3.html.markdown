---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_wtp_radio3"
description: |-
  Configuration options for radio 3.
---

# fmgdevice_wirelesscontroller_wtp_radio3
Configuration options for radio 3.

~> This resource is a sub resource for variable `radio_3` of resource `fmgdevice_wirelesscontroller_wtp`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_wtp_radio3" "trname" {
  auto_power_high   = 10
  auto_power_level  = "enable"
  auto_power_low    = 10
  auto_power_target = "your own value"
  band              = ["802.11g-only"]
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `wtp` - Wtp.

* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `disable`, `enable`.

* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - Target of automatic transmit power adjustment in dBm (-95 to -20, default = -70).
* `band` - WiFi band that Radio 3 operates on. Valid values: `802.11a`, `802.11b`, `802.11g`, `802.11n`, `802.11n-5G`, `802.11n,g-only`, `802.11g-only`, `802.11n-only`, `802.11n-5G-only`, `802.11ac`, `802.11ac,n-only`, `802.11ac-only`, `802.11ax-5G`, `802.11ax,ac-only`, `802.11ax,ac,n-only`, `802.11ax-5G-only`, `802.11ax`, `802.11ax,n-only`, `802.11ax,n,g-only`, `802.11ax-only`, `802.11ac-2G`, `802.11ax-6G`, `802.11n-2G`, `802.11ac-5G`, `802.11ax-2G`, `802.11be-2G`, `802.11be-5G`, `802.11be-6G`.

* `channel` - Selected list of wireless radio channels.
* `drma_manual_mode` - Radio mode to be used for DRMA manual mode (default = ncf). Valid values: `ap`, `monitor`, `ncf`, `ncf-peek`.

* `override_analysis` - Override-Analysis. Valid values: `disable`, `enable`.

* `override_band` - Enable to override the WTP profile band setting. Valid values: `disable`, `enable`.

* `override_channel` - Enable to override WTP profile channel settings. Valid values: `disable`, `enable`.

* `override_txpower` - Enable to override the WTP profile power level configuration. Valid values: `disable`, `enable`.

* `override_vaps` - Enable to override WTP profile Virtual Access Point (VAP) settings. Valid values: `disable`, `enable`.

* `power_level` - Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities. Valid values: `dBm`, `percentage`.

* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `radio_id` - Radio-Id.
* `spectrum_analysis` - Spectrum-Analysis. Valid values: `disable`, `enable`, `scan-only`.

* `vap_all` - Configure method for assigning SSIDs to this FortiAP (default = automatically assign tunnel SSIDs). Valid values: `disable`, `enable`, `tunnel`, `bridge`, `manual`.

* `vap1` - Virtual Access Point (VAP) for wlan ID 1
* `vap2` - Virtual Access Point (VAP) for wlan ID 2
* `vap3` - Virtual Access Point (VAP) for wlan ID 3
* `vap4` - Virtual Access Point (VAP) for wlan ID 4
* `vap5` - Virtual Access Point (VAP) for wlan ID 5
* `vap6` - Virtual Access Point (VAP) for wlan ID 6
* `vap7` - Virtual Access Point (VAP) for wlan ID 7
* `vap8` - Virtual Access Point (VAP) for wlan ID 8
* `vaps` - Manually selected list of Virtual Access Points (VAPs).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController WtpRadio3 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "wtp=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_wtp_radio3.labelname WirelessControllerWtpRadio3
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

