---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_fortiguard"
description: |-
  Configure FortiGuard services.
---

# fmgdevice_system_fortiguard
Configure FortiGuard services.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fds_license_expiring_days` - Threshold for number of days before FortiGuard license expiration to generate license expiring event log (1 - 100 days, default = 15).
* `antispam_cache` - Enable/disable FortiGuard antispam request caching. Uses a small amount of memory but improves performance. Valid values: `disable`, `enable`.

* `antispam_cache_mpermille` - Maximum permille of FortiGate memory the antispam cache is allowed to use (1 - 150).
* `antispam_cache_ttl` - Time-to-live for antispam cache entries in seconds (300 - 86400). Lower times reduce the cache size. Higher times may improve performance since the cache will have more entries.
* `antispam_expiration` - Antispam-Expiration.
* `antispam_force_off` - Enable/disable turning off the FortiGuard antispam service. Valid values: `disable`, `enable`.

* `antispam_license` - Antispam-License.
* `antispam_timeout` - Antispam query time out (1 - 30 sec, default = 7).
* `anycast_sdns_server_ip` - IP address of the FortiGuard anycast DNS rating server.
* `anycast_sdns_server_port` - Port to connect to on the FortiGuard anycast DNS rating server.
* `auto_firmware_upgrade` - Enable/disable automatic patch-level firmware upgrade from FortiGuard. The FortiGate unit searches for new patches only in the same major and minor version. Valid values: `disable`, `enable`.

* `auto_firmware_upgrade_day` - Allowed day(s) of the week to install an automatic patch-level firmware upgrade from FortiGuard (default is none). Disallow any day of the week to use auto-firmware-upgrade-delay instead, which waits for designated days before installing an automatic patch-level firmware upgrade. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.

* `auto_firmware_upgrade_delay` - Delay of day(s) before installing an automatic patch-level firmware upgrade from FortiGuard (default = 3). Set it 0 to use auto-firmware-upgrade-day instead, which selects allowed day(s) of the week for installing an automatic patch-level firmware upgrade.
* `auto_firmware_upgrade_end_hour` - End time in the designated time window for automatic patch-level firmware upgrade from FortiGuard in 24 hour time (0 ~ 23, default = 4). When the end time is smaller than the start time, the end time is interpreted as the next day. The actual upgrade time is selected randomly within the time window.
* `auto_firmware_upgrade_start_hour` - Start time in the designated time window for automatic patch-level firmware upgrade from FortiGuard in 24 hour time (0 ~ 23, default = 2). The actual upgrade time is selected randomly within the time window.
* `auto_join_forticloud` - Automatically connect to and login to FortiCloud. Valid values: `disable`, `enable`.

* `ddns_server_ip` - IP address of the FortiDDNS server.
* `ddns_server_ip6` - IPv6 address of the FortiDDNS server.
* `ddns_server_port` - Port used to communicate with FortiDDNS servers.
* `fortiguard_anycast` - Enable/disable use of FortiGuard's Anycast network. Valid values: `disable`, `enable`.

* `fortiguard_anycast_source` - Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet. Valid values: `fortinet`, `aws`, `debug`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `load_balance_servers` - Number of servers to alternate between as first FortiGuard option.
* `outbreak_prevention_cache` - Enable/disable FortiGuard Virus Outbreak Prevention cache. Valid values: `disable`, `enable`.

* `outbreak_prevention_cache_mpermille` - Maximum permille of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 150 permille, default = 1).
* `outbreak_prevention_cache_ttl` - Time-to-live for FortiGuard Virus Outbreak Prevention cache entries (300 - 86400 sec, default = 300).
* `outbreak_prevention_expiration` - Outbreak-Prevention-Expiration.
* `outbreak_prevention_force_off` - Turn off FortiGuard Virus Outbreak Prevention service. Valid values: `disable`, `enable`.

* `outbreak_prevention_license` - Outbreak-Prevention-License.
* `outbreak_prevention_timeout` - FortiGuard Virus Outbreak Prevention time out (1 - 30 sec, default = 7).
* `persistent_connection` - Enable/disable use of persistent connection to receive update notification from FortiGuard. Valid values: `disable`, `enable`.

* `port` - Port used to communicate with the FortiGuard servers. Valid values: `53`, `80`, `443`, `8888`.

* `protocol` - Protocol used to communicate with the FortiGuard servers. Valid values: `udp`, `http`, `https`.

* `proxy_password` - Proxy user password.
* `proxy_server_ip` - Hostname or IPv4 address of the proxy server.
* `proxy_server_port` - Port used to communicate with the proxy server.
* `proxy_username` - Proxy user name.
* `sandbox_inline_scan` - Enable/disable FortiCloud Sandbox inline-scan. Valid values: `disable`, `enable`.

* `sandbox_region` - FortiCloud Sandbox region.
* `sdns_options` - Customization options for the FortiGuard DNS service. Valid values: `include-question-section`.

* `sdns_server_ip` - IP address of the FortiGuard DNS rating server.
* `sdns_server_port` - Port to connect to on the FortiGuard DNS rating server.
* `service_account_id` - Service account ID.
* `source_ip` - Source IPv4 address used to communicate with FortiGuard.
* `source_ip6` - Source IPv6 address used to communicate with FortiGuard.
* `subscribe_update_notification` - Enable/disable subscription to receive update notification from FortiGuard. Valid values: `disable`, `enable`.

* `update_build_proxy` - Enable/disable proxy dictionary rebuild. Valid values: `disable`, `enable`.

* `update_dldb` - Enable/disable DLP signature update. Valid values: `disable`, `enable`.

* `update_extdb` - Enable/disable external resource update. Valid values: `disable`, `enable`.

* `update_ffdb` - Enable/disable Internet Service Database update. Valid values: `disable`, `enable`.

* `update_server_location` - Location from which to receive FortiGuard updates. Valid values: `usa`, `automatic`, `eu`.

* `update_uwdb` - Enable/disable allowlist update. Valid values: `disable`, `enable`.

* `vdom` - FortiGuard Service virtual domain name.
* `vrf_select` - VRF ID used for connection to server.
* `webfilter_cache` - Enable/disable FortiGuard web filter caching. Valid values: `disable`, `enable`.

* `webfilter_cache_ttl` - Time-to-live for web filter cache entries in seconds (300 - 86400).
* `webfilter_expiration` - Webfilter-Expiration.
* `webfilter_force_off` - Enable/disable turning off the FortiGuard web filtering service. Valid values: `disable`, `enable`.

* `webfilter_license` - Webfilter-License.
* `webfilter_timeout` - Web filter query time out (1 - 30 sec, default = 15).
* `antispam_cache_mpercent` - Antispam-Cache-Mpercent.
* `dlp_expiration` - Dlp-Expiration.
* `dlp_license` - Dlp-License.
* `fnbi_expiration` - Fnbi-Expiration.
* `fnbi_license` - Fnbi-License.
* `gui_prompt_auto_upgrade` - Gui-Prompt-Auto-Upgrade. Valid values: `disable`, `enable`.

* `ia_expiration` - Ia-Expiration.
* `ia_license` - Ia-License.
* `outbreak_prevention_cache_mpercent` - Outbreak-Prevention-Cache-Mpercent.
* `videofilter_expiration` - Videofilter-Expiration.
* `videofilter_license` - Videofilter-License.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fortiguard can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_fortiguard.labelname SystemFortiguard
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

