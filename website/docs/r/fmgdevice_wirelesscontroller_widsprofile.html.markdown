---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_widsprofile"
description: |-
  Configure wireless intrusion detection system (WIDS) profiles.
---

# fmgdevice_wirelesscontroller_widsprofile
Configure wireless intrusion detection system (WIDS) profiles.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_widsprofile" "trname" {
  ap_auto_suppress            = "enable"
  ap_bgscan_disable_schedules = ["your own value"]
  ap_bgscan_duration          = 10
  ap_bgscan_idle              = 10
  ap_bgscan_intv              = 10
  name                        = "your own value"
  device_name                 = var.device_name # not required if setting is at provider
  device_vdom                 = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `adhoc_network` - Enable/disable adhoc network detection (default = disable). Valid values: `disable`, `enable`.

* `adhoc_valid_ssid` - Enable/disable adhoc using valid SSID detection (default = disable). Valid values: `disable`, `enable`.

* `air_jack` - Enable/disable AirJack detection (default = disable). Valid values: `disable`, `enable`.

* `ap_auto_suppress` - Enable/disable on-wire rogue AP auto-suppression (default = disable). Valid values: `disable`, `enable`.

* `ap_bgscan_disable_schedules` - Firewall schedules for turning off FortiAP radio background scan. Background scan will be disabled when at least one of the schedules is valid. Separate multiple schedule names with a space.
* `ap_bgscan_duration` - Listen time on scanning a channel (10 - 1000 msec, default = 30).
* `ap_bgscan_idle` - Wait time for channel inactivity before scanning this channel (0 - 1000 msec, default = 20).
* `ap_bgscan_intv` - Period between successive channel scans (1 - 600 sec, default = 3).
* `ap_bgscan_period` - Period between background scans (10 - 3600 sec, default = 600).
* `ap_bgscan_report_intv` - Period between background scan reports (15 - 600 sec, default = 30).
* `ap_fgscan_report_intv` - Period between foreground scan reports (15 - 600 sec, default = 15).
* `ap_impersonation` - Enable/disable AP impersonation detection (default = disable). Valid values: `disable`, `enable`.

* `ap_scan` - Enable/disable rogue AP detection. Valid values: `disable`, `enable`.

* `ap_scan_channel_list_2g_5g` - Selected ap scan channel list for 2.4G and 5G bands.
* `ap_scan_channel_list_6g` - Selected ap scan channel list for 6G band.
* `ap_scan_passive` - Enable/disable passive scanning. Enable means do not send probe request on any channels (default = disable). Valid values: `disable`, `enable`.

* `ap_scan_threshold` - Minimum signal level/threshold in dBm required for the AP to report detected rogue AP (-95 to -20, default = -90).
* `ap_spoofing` - Enable/disable AP spoofing detection (default = disable). Valid values: `disable`, `enable`.

* `asleap_attack` - Enable/disable asleap attack detection (default = disable). Valid values: `disable`, `enable`.

* `assoc_flood_thresh` - The threshold value for association frame flooding.
* `assoc_flood_time` - Number of seconds after which a station is considered not connected.
* `assoc_frame_flood` - Enable/disable association frame flooding detection (default = disable). Valid values: `disable`, `enable`.

* `auth_flood_thresh` - The threshold value for authentication frame flooding.
* `auth_flood_time` - Number of seconds after which a station is considered not connected.
* `auth_frame_flood` - Enable/disable authentication frame flooding detection (default = disable). Valid values: `disable`, `enable`.

* `bcn_flood` - Enable/disable bcn flood detection (default = disable). Valid values: `disable`, `enable`.

* `bcn_flood_thresh` - The threshold value for bcn flood.
* `bcn_flood_time` - Detection Window Period.
* `beacon_wrong_channel` - Enable/disable beacon wrong channel detection (default = disable). Valid values: `disable`, `enable`.

* `block_ack_flood` - Enable/disable block_ack flood detection (default = disable). Valid values: `disable`, `enable`.

* `block_ack_flood_thresh` - The threshold value for block_ack flood.
* `block_ack_flood_time` - Detection Window Period.
* `chan_based_mitm` - Enable/disable channel based mitm detection (default = disable). Valid values: `disable`, `enable`.

* `client_flood` - Enable/disable client flood detection (default = disable). Valid values: `disable`, `enable`.

* `client_flood_thresh` - The threshold value for client flood.
* `client_flood_time` - Detection Window Period.
* `comment` - Comment.
* `cts_flood` - Enable/disable cts flood detection (default = disable). Valid values: `disable`, `enable`.

* `cts_flood_thresh` - The threshold value for cts flood.
* `cts_flood_time` - Detection Window Period.
* `deauth_broadcast` - Enable/disable broadcasting de-authentication detection (default = disable). Valid values: `disable`, `enable`.

* `deauth_unknown_src_thresh` - Threshold value per second to deauth unknown src for DoS attack (0: no limit).
* `disassoc_broadcast` - Enable/disable broadcast dis-association detection (default = disable). Valid values: `disable`, `enable`.

* `disconnect_station` - Enable/disable disconnect station detection (default = disable). Valid values: `disable`, `enable`.

* `eapol_fail_flood` - Enable/disable EAPOL-Failure flooding (to AP) detection (default = disable). Valid values: `disable`, `enable`.

* `eapol_fail_intv` - The detection interval for EAPOL-Failure flooding (1 - 3600 sec).
* `eapol_fail_thresh` - The threshold value for EAPOL-Failure flooding in specified interval.
* `eapol_key_overflow` - Enable/disable overflow EAPOL key detection (default = disable). Valid values: `disable`, `enable`.

* `eapol_logoff_flood` - Enable/disable EAPOL-Logoff flooding (to AP) detection (default = disable). Valid values: `disable`, `enable`.

* `eapol_logoff_intv` - The detection interval for EAPOL-Logoff flooding (1 - 3600 sec).
* `eapol_logoff_thresh` - The threshold value for EAPOL-Logoff flooding in specified interval.
* `eapol_pre_fail_flood` - Enable/disable premature EAPOL-Failure flooding (to STA) detection (default = disable). Valid values: `disable`, `enable`.

* `eapol_pre_fail_intv` - The detection interval for premature EAPOL-Failure flooding (1 - 3600 sec).
* `eapol_pre_fail_thresh` - The threshold value for premature EAPOL-Failure flooding in specified interval.
* `eapol_pre_succ_flood` - Enable/disable premature EAPOL-Success flooding (to STA) detection (default = disable). Valid values: `disable`, `enable`.

* `eapol_pre_succ_intv` - The detection interval for premature EAPOL-Success flooding (1 - 3600 sec).
* `eapol_pre_succ_thresh` - The threshold value for premature EAPOL-Success flooding in specified interval.
* `eapol_start_flood` - Enable/disable EAPOL-Start flooding (to AP) detection (default = disable). Valid values: `disable`, `enable`.

* `eapol_start_intv` - The detection interval for EAPOL-Start flooding (1 - 3600 sec).
* `eapol_start_thresh` - The threshold value for EAPOL-Start flooding in specified interval.
* `eapol_succ_flood` - Enable/disable EAPOL-Success flooding (to AP) detection (default = disable). Valid values: `disable`, `enable`.

* `eapol_succ_intv` - The detection interval for EAPOL-Success flooding (1 - 3600 sec).
* `eapol_succ_thresh` - The threshold value for EAPOL-Success flooding in specified interval.
* `fata_jack` - Enable/disable FATA-Jack detection (default = disable). Valid values: `disable`, `enable`.

* `fuzzed_beacon` - Enable/disable fuzzed beacon detection (default = disable). Valid values: `disable`, `enable`.

* `fuzzed_probe_request` - Enable/disable fuzzed probe request detection (default = disable). Valid values: `disable`, `enable`.

* `fuzzed_probe_response` - Enable/disable fuzzed probe response detection (default = disable). Valid values: `disable`, `enable`.

* `hotspotter_attack` - Enable/disable hotspotter attack detection (default = disable). Valid values: `disable`, `enable`.

* `ht_40mhz_intolerance` - Enable/disable HT 40 MHz intolerance detection (default = disable). Valid values: `disable`, `enable`.

* `ht_greenfield` - Enable/disable HT greenfield detection (default = disable). Valid values: `disable`, `enable`.

* `invalid_addr_combination` - Enable/disable invalid address combination detection (default = disable). Valid values: `disable`, `enable`.

* `invalid_mac_oui` - Enable/disable invalid MAC OUI detection. Valid values: `disable`, `enable`.

* `long_duration_attack` - Enable/disable long duration attack detection based on user configured threshold (default = disable). Valid values: `disable`, `enable`.

* `long_duration_thresh` - Threshold value for long duration attack detection (1000 - 32767 usec, default = 8200).
* `malformed_association` - Enable/disable malformed association request detection (default = disable). Valid values: `disable`, `enable`.

* `malformed_auth` - Enable/disable malformed auth frame detection (default = disable). Valid values: `disable`, `enable`.

* `malformed_ht_ie` - Enable/disable malformed HT IE detection (default = disable). Valid values: `disable`, `enable`.

* `name` - WIDS profile name.
* `netstumbler` - Enable/disable netstumbler detection (default = disable). Valid values: `disable`, `enable`.

* `netstumbler_thresh` - The threshold value for netstumbler.
* `netstumbler_time` - Detection Window Period.
* `null_ssid_probe_resp` - Enable/disable null SSID probe response detection (default = disable). Valid values: `disable`, `enable`.

* `omerta_attack` - Enable/disable omerta attack detection (default = disable). Valid values: `disable`, `enable`.

* `overflow_ie` - Enable/disable overflow IE detection (default = disable). Valid values: `disable`, `enable`.

* `probe_flood` - Enable/disable probe flood detection (default = disable). Valid values: `disable`, `enable`.

* `probe_flood_thresh` - The threshold value for probe flood.
* `probe_flood_time` - Detection Window Period.
* `pspoll_flood` - Enable/disable pspoll flood detection (default = disable). Valid values: `disable`, `enable`.

* `pspoll_flood_thresh` - The threshold value for pspoll flood.
* `pspoll_flood_time` - Detection Window Period.
* `pwsave_dos_attack` - Enable/disable power save DOS attack detection (default = disable). Valid values: `disable`, `enable`.

* `reassoc_flood` - Enable/disable reassociation flood detection (default = disable). Valid values: `disable`, `enable`.

* `reassoc_flood_thresh` - The threshold value for reassociation flood.
* `reassoc_flood_time` - Detection Window Period.
* `risky_encryption` - Enable/disable Risky Encryption detection (default = disable). Valid values: `disable`, `enable`.

* `rts_flood` - Enable/disable rts flood detection (default = disable). Valid values: `disable`, `enable`.

* `rts_flood_thresh` - The threshold value for rts flood.
* `rts_flood_time` - Detection Window Period.
* `sensor_mode` - Scan nearby WiFi stations (default = disable). Valid values: `disable`, `foreign`, `both`.

* `spoofed_deauth` - Enable/disable spoofed de-authentication attack detection (default = disable). Valid values: `disable`, `enable`.

* `unencrypted_valid` - Enable/disable unencrypted valid detection (default = disable). Valid values: `disable`, `enable`.

* `valid_client_misassociation` - Enable/disable valid client misassociation detection (default = disable). Valid values: `disable`, `enable`.

* `valid_ssid_misuse` - Enable/disable valid SSID misuse detection (default = disable). Valid values: `disable`, `enable`.

* `weak_wep_iv` - Enable/disable weak WEP IV (Initialization Vector) detection (default = disable). Valid values: `disable`, `enable`.

* `wellenreiter` - Enable/disable wellenreiter detection (default = disable). Valid values: `disable`, `enable`.

* `wellenreiter_thresh` - The threshold value for wellenreiter.
* `wellenreiter_time` - Detection Window Period.
* `windows_bridge` - Enable/disable windows bridge detection (default = disable). Valid values: `disable`, `enable`.

* `wireless_bridge` - Enable/disable wireless bridge detection (default = disable). Valid values: `disable`, `enable`.

* `wpa_ft_attack` - Enable/disable WPA FT attack detection (default = disable). Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController WidsProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_widsprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

