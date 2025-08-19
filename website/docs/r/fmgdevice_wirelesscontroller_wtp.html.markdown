---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_wtp"
description: |-
  Configure Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.
---

# fmgdevice_wirelesscontroller_wtp
Configure Wireless Termination Points (WTPs), that is, FortiAPs or APs to be managed by FortiGate.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `lan`: `fmgdevice_wirelesscontroller_wtp_lan`
>- `radio_1`: `fmgdevice_wirelesscontroller_wtp_radio1`
>- `radio_2`: `fmgdevice_wirelesscontroller_wtp_radio2`
>- `radio_3`: `fmgdevice_wirelesscontroller_wtp_radio3`
>- `radio_4`: `fmgdevice_wirelesscontroller_wtp_radio4`
>- `split_tunneling_acl`: `fmgdevice_wirelesscontroller_wtp_splittunnelingacl`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_wtp" "trname" {
  admin         = "discovery"
  allowaccess   = ["telnet"]
  apcfg_profile = ["your own value"]
  ble_major_id  = 10
  ble_minor_id  = 10
  wtp_id        = "your own value"
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `admin` - Configure how the FortiGate operating as a wireless controller discovers and manages this WTP, AP or FortiAP. Valid values: `discovery`, `disable`, `enable`, `discovered`.

* `allowaccess` - Control management access to the managed WTP, FortiAP, or AP. Separate entries with a space. Valid values: `https`, `ssh`, `snmp`, `http`, `telnet`.

* `apcfg_profile` - AP local configuration profile name.
* `ble_major_id` - Override BLE Major ID.
* `ble_minor_id` - Override BLE Minor ID.
* `bonjour_profile` - Bonjour profile name.
* `comment` - Comment.
* `coordinate_latitude` - WTP latitude coordinate.
* `coordinate_longitude` - WTP longitude coordinate.
* `firmware_provision` - Firmware version to provision to this FortiAP on bootup (major.minor.build, i.e. 6.2.1234).
* `firmware_provision_latest` - Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.

* `image_download` - Enable/disable WTP image download. Valid values: `disable`, `enable`.

* `index` - Index.
* `ip_fragment_preventing` - Method(s) by which IP fragmentation is prevented for control and data packets through CAPWAP tunnel (default = tcp-mss-adjust). Valid values: `tcp-mss-adjust`, `icmp-unreachable`.

* `lan` - Lan. The structure of `lan` block is documented below.
* `led_state` - Enable to allow the FortiAPs LEDs to light. Disable to keep the LEDs off. You may want to keep the LEDs off so they are not distracting in low light areas etc. Valid values: `disable`, `enable`.

* `location` - Field for describing the physical location of the WTP, AP or FortiAP.
* `login_passwd` - Set the managed WTP, FortiAP, or AP's administrator password.
* `login_passwd_change` - Change or reset the administrator password of a managed WTP, FortiAP or AP (yes, default, or no, default = no). Valid values: `no`, `yes`, `default`.

* `mesh_bridge_enable` - Enable/disable mesh Ethernet bridge when WTP is configured as a mesh branch/leaf AP. Valid values: `disable`, `enable`, `default`.

* `name` - WTP, AP or FortiAP configuration name.
* `override_allowaccess` - Enable to override the WTP profile management access configuration. Valid values: `disable`, `enable`.

* `override_ip_fragment` - Enable/disable overriding the WTP profile IP fragment prevention setting. Valid values: `disable`, `enable`.

* `override_lan` - Enable to override the WTP profile LAN port setting. Valid values: `disable`, `enable`.

* `override_led_state` - Enable to override the profile LED state setting for this FortiAP. You must enable this option to use the led-state command to turn off the FortiAP's LEDs. Valid values: `disable`, `enable`.

* `override_login_passwd_change` - Enable to override the WTP profile login-password (administrator password) setting. Valid values: `disable`, `enable`.

* `override_split_tunnel` - Enable/disable overriding the WTP profile split tunneling setting. Valid values: `disable`, `enable`.

* `override_wan_port_mode` - Enable/disable overriding the wan-port-mode in the WTP profile. Valid values: `disable`, `enable`.

* `purdue_level` - Purdue Level of this WTP. Valid values: `1`, `2`, `3`, `4`, `5`, `1.5`, `2.5`, `3.5`, `5.5`.

* `radio_1` - Radio-1. The structure of `radio_1` block is documented below.
* `radio_2` - Radio-2. The structure of `radio_2` block is documented below.
* `radio_3` - Radio-3. The structure of `radio_3` block is documented below.
* `radio_4` - Radio-4. The structure of `radio_4` block is documented below.
* `region` - Region name WTP is associated with.
* `region_x` - Relative horizontal region coordinate (between 0 and 1).
* `region_y` - Relative vertical region coordinate (between 0 and 1).
* `split_tunneling_acl` - Split-Tunneling-Acl. The structure of `split_tunneling_acl` block is documented below.
* `split_tunneling_acl_local_ap_subnet` - Enable/disable automatically adding local subnetwork of FortiAP to split-tunneling ACL (default = disable). Valid values: `disable`, `enable`.

* `split_tunneling_acl_path` - Split tunneling ACL path is local/tunnel. Valid values: `tunnel`, `local`.

* `tun_mtu_downlink` - The MTU of downlink CAPWAP tunnel (576 - 1500 bytes or 0; 0 means the local MTU of FortiAP; default = 0).
* `tun_mtu_uplink` - The maximum transmission unit (MTU) of uplink CAPWAP tunnel (576 - 1500 bytes or 0; 0 means the local MTU of FortiAP; default = 0).
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `wan_port_mode` - Enable/disable using the FortiAP WAN port as a LAN port. Valid values: `wan-lan`, `wan-only`.

* `wtp_id` - WTP ID.
* `wtp_mode` - WTP, AP, or FortiAP operating mode (default = normal). Valid values: `normal`, `remote`.

* `wtp_profile` - WTP profile name to apply to this WTP, AP or FortiAP.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `lan` block supports:

* `port_esl_mode` - ESL port mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port_esl_ssid` - Bridge ESL port to SSID.
* `port_mode` - LAN port mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port_ssid` - Bridge LAN port to SSID.
* `port1_mode` - LAN port 1 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port1_ssid` - Bridge LAN port 1 to SSID.
* `port2_mode` - LAN port 2 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port2_ssid` - Bridge LAN port 2 to SSID.
* `port3_mode` - LAN port 3 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port3_ssid` - Bridge LAN port 3 to SSID.
* `port4_mode` - LAN port 4 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port4_ssid` - Bridge LAN port 4 to SSID.
* `port5_mode` - LAN port 5 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port5_ssid` - Bridge LAN port 5 to SSID.
* `port6_mode` - LAN port 6 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port6_ssid` - Bridge LAN port 6 to SSID.
* `port7_mode` - LAN port 7 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port7_ssid` - Bridge LAN port 7 to SSID.
* `port8_mode` - LAN port 8 mode. Valid values: `offline`, `bridge-to-wan`, `bridge-to-ssid`, `nat-to-wan`.

* `port8_ssid` - Bridge LAN port 8 to SSID.

The `radio_1` block supports:

* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `disable`, `enable`.

* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - Target of automatic transmit power adjustment in dBm (-95 to -20, default = -70).
* `band` - WiFi band that Radio 1 operates on. Valid values: `802.11a`, `802.11b`, `802.11g`, `802.11n`, `802.11n-5G`, `802.11n,g-only`, `802.11g-only`, `802.11n-only`, `802.11n-5G-only`, `802.11ac`, `802.11ac,n-only`, `802.11ac-only`, `802.11ax-5G`, `802.11ax,ac-only`, `802.11ax,ac,n-only`, `802.11ax-5G-only`, `802.11ax`, `802.11ax,n-only`, `802.11ax,n,g-only`, `802.11ax-only`, `802.11ac-2G`, `802.11ax-6G`, `802.11n-2G`, `802.11ac-5G`, `802.11ax-2G`, `802.11be-2G`, `802.11be-5G`, `802.11be-6G`.

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

The `radio_2` block supports:

* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `disable`, `enable`.

* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - Target of automatic transmit power adjustment in dBm (-95 to -20, default = -70).
* `band` - WiFi band that Radio 2 operates on. Valid values: `802.11a`, `802.11b`, `802.11g`, `802.11n`, `802.11n-5G`, `802.11n,g-only`, `802.11g-only`, `802.11n-only`, `802.11n-5G-only`, `802.11ac`, `802.11ac,n-only`, `802.11ac-only`, `802.11ax-5G`, `802.11ax,ac-only`, `802.11ax,ac,n-only`, `802.11ax-5G-only`, `802.11ax`, `802.11ax,n-only`, `802.11ax,n,g-only`, `802.11ax-only`, `802.11ac-2G`, `802.11ax-6G`, `802.11n-2G`, `802.11ac-5G`, `802.11ax-2G`, `802.11be-2G`, `802.11be-5G`, `802.11be-6G`.

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

The `radio_3` block supports:

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

The `radio_4` block supports:

* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `disable`, `enable`.

* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - Target of automatic transmit power adjustment in dBm (-95 to -20, default = -70).
* `band` - WiFi band that Radio 4 operates on. Valid values: `802.11a`, `802.11b`, `802.11g`, `802.11n`, `802.11n-5G`, `802.11n,g-only`, `802.11g-only`, `802.11n-only`, `802.11n-5G-only`, `802.11ac`, `802.11ac,n-only`, `802.11ac-only`, `802.11ax-5G`, `802.11ax,ac-only`, `802.11ax,ac,n-only`, `802.11ax-5G-only`, `802.11ax`, `802.11ax,n-only`, `802.11ax,n,g-only`, `802.11ax-only`, `802.11ac-2G`, `802.11ax-6G`, `802.11n-2G`, `802.11ac-5G`, `802.11ax-2G`, `802.11be-2G`, `802.11be-5G`, `802.11be-6G`.

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

The `split_tunneling_acl` block supports:

* `dest_ip` - Destination IP and mask for the split-tunneling subnet.
* `id` - ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{wtp_id}}.

## Import

WirelessController Wtp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_wtp.labelname {{wtp_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

