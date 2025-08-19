---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_global"
description: |-
  Configure wireless controller global settings.
---

# fmgdevice_wirelesscontroller_global
Configure wireless controller global settings.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_global" "trname" {
  acd_process_count       = 10
  ap_log_server           = "enable"
  ap_log_server_ip        = "your own value"
  ap_log_server_port      = 10
  control_message_offload = ["aeroscout-mu"]
  device_name             = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `acd_process_count` - Configure the number cw_acd daemons for multi-core CPU support (default = 0).
* `ap_log_server` - Enable/disable configuring FortiGate to redirect wireless event log messages or FortiAPs to send UTM log messages to a syslog server (default = disable). Valid values: `disable`, `enable`.

* `ap_log_server_ip` - IP address that FortiGate or FortiAPs send log messages to.
* `ap_log_server_port` - Port that FortiGate or FortiAPs send log messages to.
* `control_message_offload` - Configure CAPWAP control message data channel offload. Valid values: `ebp-frame`, `aeroscout-tag`, `ap-list`, `sta-list`, `sta-cap-list`, `stats`, `aeroscout-mu`, `sta-health`, `spectral-analysis`.

* `data_ethernet_ii` - Configure the wireless controller to use Ethernet II or 802.3 frames with 802.3 data tunnel mode (default = enable). Valid values: `disable`, `enable`.

* `dfs_lab_test` - Enable/disable DFS certificate lab test mode. Valid values: `disable`, `enable`.

* `discovery_mc_addr` - Multicast IP address for AP discovery (default = 244.0.1.140).
* `fiapp_eth_type` - Ethernet type for Fortinet Inter-Access Point Protocol (IAPP), or IEEE 802.11f, packets (0 - 65535, default = 5252).
* `image_download` - Enable/disable WTP image download at join time. Valid values: `disable`, `enable`.

* `ipsec_base_ip` - Base IP address for IPsec VPN tunnels between the access points and the wireless controller (default = 169.254.0.1).
* `link_aggregation` - Enable/disable calculating the CAPWAP transmit hash to load balance sessions to link aggregation nodes (default = disable). Valid values: `disable`, `enable`.

* `local_radio_vdom` - Assign local radio's virtual domain.
* `location` - Description of the location of the wireless controller.
* `max_ble_device` - Maximum number of BLE devices stored on the controller (default = 0).
* `max_clients` - Maximum number of clients that can connect simultaneously (default = 0, meaning no limitation).
* `max_retransmit` - Maximum number of tunnel packet retransmissions (0 - 64, default = 3).
* `max_rogue_ap` - Maximum number of rogue APs stored on the controller (default = 0).
* `max_rogue_ap_wtp` - Maximum number of rogue AP's wtp info stored on the controller (1 - 16, default = 16).
* `max_rogue_sta` - Maximum number of rogue stations stored on the controller (default = 0).
* `max_sta_cap` - Maximum number of station cap stored on the controller (default = 0).
* `max_sta_cap_wtp` - Maximum number of station cap's wtp info stored on the controller (1 - 16, default = 8).
* `max_wids_entry` - Maximum number of wids entries stored on the controller (default = 0).
* `mesh_eth_type` - Mesh Ethernet identifier included in backhaul packets (0 - 65535, default = 8755).
* `nac_interval` - Interval in seconds between two WiFi network access control (NAC) checks (10 - 600, default = 120).
* `name` - Name of the wireless controller.
* `rogue_scan_mac_adjacency` - Maximum numerical difference between an AP's Ethernet and wireless MAC values to match for rogue detection (0 - 31, default = 7).
* `rolling_wtp_upgrade` - Enable/disable rolling WTP upgrade (default = disable). Valid values: `disable`, `enable`.

* `rolling_wtp_upgrade_threshold` - Minimum signal level/threshold in dBm required for the managed WTP to be included in rolling WTP upgrade (-95 to -20, default = -80).
* `tunnel_mode` - Compatible/strict tunnel mode. Valid values: `compatible`, `strict`.

* `wpad_process_count` - Wpad daemon process count for multi-core CPU support.
* `wtp_share` - Enable/disable sharing of WTPs between VDOMs. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController Global can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_global.labelname WirelessControllerGlobal
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

