---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_wtpprofile_radio2"
description: |-
  Configuration options for radio 2.
---

# fmgdevice_wirelesscontroller_wtpprofile_radio2
Configuration options for radio 2.

~> This resource is a sub resource for variable `radio_2` of resource `fmgdevice_wirelesscontroller_wtpprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_wtpprofile_radio2" "trname" {
  n80211d          = "enable"
  airtime_fairness = "disable"
  amsdu            = "enable"
  ap_handoff       = "disable"
  ap_sniffer_addr  = "your own value"
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `wtp_profile` - Wtp Profile.

* `n80211d` - Enable/disable 802.11d countryie(default = enable). Valid values: `disable`, `enable`.

* `n80211mc` - Enable/disable 802.11mc responder mode (default = disable). Valid values: `disable`, `enable`.

* `airtime_fairness` - Enable/disable airtime fairness (default = disable). Valid values: `disable`, `enable`.

* `amsdu` - Enable/disable 802.11n AMSDU support. AMSDU can improve performance if supported by your WiFi clients (default = enable). Valid values: `disable`, `enable`.

* `ap_handoff` - Enable/disable AP handoff of clients to other APs (default = disable). Valid values: `disable`, `enable`.

* `ap_sniffer_addr` - MAC address to monitor.
* `ap_sniffer_bufsize` - Sniffer buffer size (1 - 32 MB, default = 16).
* `ap_sniffer_chan` - Channel on which to operate the sniffer (default = 6).
* `ap_sniffer_chan_width` - Channel bandwidth for sniffer. Valid values: `320MHz`, `240MHz`, `160MHz`, `80MHz`, `40MHz`, `20MHz`.

* `ap_sniffer_ctl` - Enable/disable sniffer on WiFi control frame (default = enable). Valid values: `disable`, `enable`.

* `ap_sniffer_data` - Enable/disable sniffer on WiFi data frame (default = enable). Valid values: `disable`, `enable`.

* `ap_sniffer_mgmt_beacon` - Enable/disable sniffer on WiFi management Beacon frames (default = enable). Valid values: `disable`, `enable`.

* `ap_sniffer_mgmt_other` - Enable/disable sniffer on WiFi management other frames  (default = enable). Valid values: `disable`, `enable`.

* `ap_sniffer_mgmt_probe` - Enable/disable sniffer on WiFi management probe frames (default = enable). Valid values: `disable`, `enable`.

* `arrp_profile` - Distributed Automatic Radio Resource Provisioning (DARRP) profile name to assign to the radio.
* `auto_power_high` - The upper bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_level` - Enable/disable automatic power-level adjustment to prevent co-channel interference (default = enable). Valid values: `disable`, `enable`.

* `auto_power_low` - The lower bound of automatic transmit power adjustment in dBm (the actual range of transmit power depends on the AP platform type).
* `auto_power_target` - Target of automatic transmit power adjustment in dBm (-95 to -20, default = -70).
* `band` - WiFi band that Radio 2 operates on. Valid values: `802.11a`, `802.11b`, `802.11g`, `802.11n`, `802.11n-5G`, `802.11n,g-only`, `802.11g-only`, `802.11n-only`, `802.11n-5G-only`, `802.11ac`, `802.11ac,n-only`, `802.11ac-only`, `802.11ax-5G`, `802.11ax,ac-only`, `802.11ax,ac,n-only`, `802.11ax-5G-only`, `802.11ax`, `802.11ax,n-only`, `802.11ax,n,g-only`, `802.11ax-only`, `802.11ac-2G`, `802.11ax-6G`, `802.11n-2G`, `802.11ac-5G`, `802.11ax-2G`, `802.11be-2G`, `802.11be-5G`, `802.11be-6G`.

* `band_5g_type` - WiFi 5G band type. Valid values: `5g-full`, `5g-high`, `5g-low`.

* `bandwidth_admission_control` - Enable/disable WiFi multimedia (WMM) bandwidth admission control to optimize WiFi bandwidth use. A request to join the wireless network is only allowed if the access point has enough bandwidth to support it. Valid values: `disable`, `enable`.

* `bandwidth_capacity` - Maximum bandwidth capacity allowed (1 - 600000 Kbps, default = 2000).
* `beacon_interval` - Beacon interval. The time between beacon frames in milliseconds. Actual range of beacon interval depends on the AP platform type (default = 100).
* `bss_color` - BSS color value for this 11ax radio (0 - 63, disable = 0, default = 0).
* `bss_color_mode` - BSS color mode for this 11ax radio (default = auto). Valid values: `auto`, `static`.

* `call_admission_control` - Enable/disable WiFi multimedia (WMM) call admission control to optimize WiFi bandwidth use for VoIP calls. New VoIP calls are only accepted if there is enough bandwidth available to support them. Valid values: `disable`, `enable`.

* `call_capacity` - Maximum number of Voice over WLAN (VoWLAN) phones supported by the radio (0 - 60, default = 10).
* `channel` - Selected list of wireless radio channels.
* `channel_bonding` - Channel bandwidth: 320, 240, 160, 80, 40, or 20MHz. Channels may use both 20 and 40 by enabling coexistence. Valid values: `disable`, `enable`, `80MHz`, `40MHz`, `20MHz`, `160MHz`, `320MHz`, `240MHz`.

* `channel_bonding_ext` - Channel bandwidth extension: 320 MHz-1 and 320 MHz-2 (default = 320 MHz-2). Valid values: `320MHz-1`, `320MHz-2`.

* `channel_utilization` - Enable/disable measuring channel utilization. Valid values: `disable`, `enable`.

* `coexistence` - Enable/disable allowing both HT20 and HT40 on the same radio (default = enable). Valid values: `disable`, `enable`.

* `darrp` - Enable/disable Distributed Automatic Radio Resource Provisioning (DARRP) to make sure the radio is always using the most optimal channel (default = disable). Valid values: `disable`, `enable`.

* `drma` - Enable/disable dynamic radio mode assignment (DRMA) (default = disable). Valid values: `disable`, `enable`.

* `drma_sensitivity` - Network Coverage Factor (NCF) percentage required to consider a radio as redundant (default = low). Valid values: `low`, `medium`, `high`.

* `dtim` - Delivery Traffic Indication Map (DTIM) period (1 - 255, default = 1). Set higher to save battery life of WiFi client in power-save mode.
* `frag_threshold` - Maximum packet size that can be sent without fragmentation (800 - 2346 bytes, default = 2346).
* `frequency_handoff` - Enable/disable frequency handoff of clients to other channels (default = disable). Valid values: `disable`, `enable`.

* `iperf_protocol` - Iperf test protocol (default = "UDP"). Valid values: `udp`, `tcp`.

* `iperf_server_port` - Iperf service port number.
* `max_clients` - Maximum number of stations (STAs) or WiFi clients supported by the radio. Range depends on the hardware.
* `max_distance` - Maximum expected distance between the AP and clients (0 - 54000 m, default = 0).
* `mimo_mode` - Configure radio MIMO mode (default = default). Valid values: `default`, `1x1`, `2x2`, `3x3`, `4x4`, `8x8`.

* `mode` - Mode of radio 2. Radio 2 can be disabled, configured as an access point, a rogue AP monitor, a sniffer, or a station. Valid values: `disabled`, `ap`, `monitor`, `sniffer`, `sam`.

* `optional_antenna` - Optional antenna used on FAP (default = none). Valid values: `none`, `FANT-04ABGN-0606-O-N`, `FANT-04ABGN-1414-P-N`, `FANT-04ABGN-8065-P-N`, `FANT-04ABGN-0606-O-R`, `FANT-04ABGN-0606-P-R`, `FANT-10ACAX-1213-D-N`, `FANT-08ABGN-1213-D-R`, `custom`.

* `optional_antenna_gain` - Optional antenna gain in dBi (0 to 20, default = 0).
* `power_level` - Radio EIRP power level as a percentage of the maximum EIRP power (0 - 100, default = 100).
* `power_mode` - Set radio effective isotropic radiated power (EIRP) in dBm or by a percentage of the maximum EIRP (default = percentage). This power takes into account both radio transmit power and antenna gain. Higher power level settings may be constrained by local regulatory requirements and AP capabilities. Valid values: `dBm`, `percentage`.

* `power_value` - Radio EIRP power in dBm (1 - 33, default = 27).
* `powersave_optimize` - Enable client power-saving features such as TIM, AC VO, and OBSS etc. Valid values: `tim`, `ac-vo`, `no-obss-scan`, `no-11b-rate`, `client-rate-follow`.

* `protection_mode` - Enable/disable 802.11g protection modes to support backwards compatibility with older clients (rtscts, ctsonly, disable). Valid values: `rtscts`, `ctsonly`, `disable`.

* `radio_id` - Radio-Id.
* `rts_threshold` - Maximum packet size for RTS transmissions, specifying the maximum size of a data packet before RTS/CTS (256 - 2346 bytes, default = 2346).
* `sam_bssid` - BSSID for WiFi network.
* `sam_ca_certificate` - CA certificate for WPA2/WPA3-ENTERPRISE.
* `sam_captive_portal` - Enable/disable Captive Portal Authentication (default = disable). Valid values: `disable`, `enable`.

* `sam_client_certificate` - Client certificate for WPA2/WPA3-ENTERPRISE.
* `sam_cwp_failure_string` - Failure identification on the page after an incorrect login.
* `sam_cwp_match_string` - Identification string from the captive portal login form.
* `sam_cwp_password` - Password for captive portal authentication.
* `sam_cwp_success_string` - Success identification on the page after a successful login.
* `sam_cwp_test_url` - Website the client is trying to access.
* `sam_cwp_username` - Username for captive portal authentication.
* `sam_eap_method` - Select WPA2/WPA3-ENTERPRISE EAP Method (default = PEAP). Valid values: `tls`, `peap`, `both`.

* `sam_password` - Passphrase for WiFi network connection.
* `sam_private_key` - Private key for WPA2/WPA3-ENTERPRISE.
* `sam_private_key_password` - Password for private key file for WPA2/WPA3-ENTERPRISE.
* `sam_report_intv` - SAM report interval (sec), 0 for a one-time report.
* `sam_security_type` - Select WiFi network security type (default = "wpa-personal"). Valid values: `open`, `wpa-personal`, `wpa-enterprise`, `owe`, `wpa3-sae`.

* `sam_server` - Sam-Server.
* `sam_server_fqdn` - SAM test server domain name.
* `sam_server_ip` - SAM test server IP address.
* `sam_server_type` - Select SAM server type (default = "IP"). Valid values: `ip`, `fqdn`.

* `sam_ssid` - SSID for WiFi network.
* `sam_test` - Select SAM test type (default = "PING"). Valid values: `ping`, `iperf`.

* `sam_username` - Username for WiFi network connection.
* `short_guard_interval` - Use either the short guard interval (Short GI) of 400 ns or the long guard interval (Long GI) of 800 ns. Valid values: `disable`, `enable`.

* `spectrum_analysis` - Spectrum-Analysis. Valid values: `disable`, `enable`, `scan-only`.

* `transmit_optimize` - Packet transmission optimization options including power saving, aggregation limiting, retry limiting, etc. All are enabled by default. Valid values: `disable`, `power-save`, `aggr-limit`, `retry-limit`, `send-bar`.

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
* `wids_profile` - Wireless Intrusion Detection System (WIDS) profile name to assign to the radio.
* `zero_wait_dfs` - Enable/disable zero wait DFS on radio (default = enable). Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController WtpProfileRadio2 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "wtp_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_wtpprofile_radio2.labelname WirelessControllerWtpProfileRadio2
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

