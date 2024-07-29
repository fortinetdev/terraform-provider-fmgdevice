---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_timers"
description: |-
  Configure CAPWAP timers.
---

# fmgdevice_wirelesscontroller_timers
Configure CAPWAP timers.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_timers" "trname" {
  ap_reboot_wait_interval1 = 10
  ap_reboot_wait_interval2 = 10
  ap_reboot_wait_time      = "your own value"
  auth_timeout             = 10
  ble_device_cleanup       = 10
  device_name              = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `ap_reboot_wait_interval1` - Time in minutes to wait before AP reboots when there is no controller detected (5 - 65535, default = 0, 0 for no reboot).
* `ap_reboot_wait_interval2` - Time in minutes to wait before AP reboots when there is no controller detected and standalone SSIDs are pushed to the AP in the previous session (5 - 65535, default = 0, 0 for no reboot).
* `ap_reboot_wait_time` - Time to reboot the AP when there is no controller detected and standalone SSIDs are pushed to the AP in the previous session, format hh:mm.
* `auth_timeout` - Time after which a client is considered failed in RADIUS authentication and times out (5 - 30 sec, default = 5).
* `ble_device_cleanup` - Time period in minutes to keep BLE device after it is gone (default = 60).
* `ble_scan_report_intv` - Time between running Bluetooth Low Energy (BLE) reports (10 - 3600 sec, default = 30).
* `client_idle_rehome_timeout` - Time after which a client is considered idle and disconnected from the home controller (2 - 3600 sec, default = 20, 0 for no timeout).
* `client_idle_timeout` - Time after which a client is considered idle and times out (20 - 3600 sec, default = 300, 0 for no timeout).
* `discovery_interval` - Time between discovery requests (2 - 180 sec, default = 5).
* `drma_interval` - Dynamic radio mode assignment (DRMA) schedule interval in minutes (10 - 1440, default = 60).
* `echo_interval` - Time between echo requests sent by the managed WTP, AP, or FortiAP (1 - 255 sec, default = 30).
* `fake_ap_log` - Time between recording logs about fake APs if periodic fake AP logging is configured (0 - 1440 min, default = 1).
* `ipsec_intf_cleanup` - Time period to keep IPsec VPN interfaces up after WTP sessions are disconnected (30 - 3600 sec, default = 120).
* `nat_session_keep_alive` - Maximal time in seconds between control requests sent by the managed WTP, AP, or FortiAP (0 - 255 sec, default = 0).
* `radio_stats_interval` - Time between running radio reports (1 - 255 sec, default = 15).
* `rogue_ap_cleanup` - Time period in minutes to keep rogue AP after it is gone (default = 0).
* `rogue_ap_log` - Time between logging rogue AP messages if periodic rogue AP logging is configured (0 - 1440 min, default = 0).
* `rogue_sta_cleanup` - Time period in minutes to keep rogue station after it is gone (default = 0).
* `sta_cap_cleanup` - Time period in minutes to keep station capability data after it is gone (default = 0).
* `sta_capability_interval` - Time between running station capability reports (1 - 255 sec, default = 30).
* `sta_locate_timer` - Time between running client presence flushes to remove clients that are listed but no longer present (0 - 86400 sec, default = 1800).
* `sta_stats_interval` - Time between running client (station) reports (1 - 255 sec, default = 10).
* `vap_stats_interval` - Time between running Virtual Access Point (VAP) reports (1 - 255 sec, default = 15).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController Timers can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_timers.labelname WirelessControllerTimers
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

