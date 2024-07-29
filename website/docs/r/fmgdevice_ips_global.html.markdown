---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ips_global"
description: |-
  Configure IPS global parameter.
---

# fmgdevice_ips_global
Configure IPS global parameter.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `tls_active_probe`: `fmgdevice_ips_global_tlsactiveprobe`



## Example Usage

```hcl
resource "fmgdevice_ips_global" "trname" {
  anomaly_mode           = "periodical"
  av_mem_limit           = 10
  cp_accel_mode          = "basic"
  database               = "extended"
  deep_app_insp_db_limit = 10
  device_name            = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `anomaly_mode` - Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.

* `av_mem_limit` - Maximum percentage of system memory allowed for use on AV scanning (10 - 50, default = zero). To disable set to zero. When disabled, there is no limit on the AV memory usage.
* `cp_accel_mode` - IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.

* `database` - Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.

* `deep_app_insp_db_limit` - Limit on number of entries in deep application inspection database (1 - 2147483647, use recommended setting = 0).
* `deep_app_insp_timeout` - Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
* `engine_count` - Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
* `exclude_signatures` - Excluded signatures. Valid values: `none`, `industrial`, `ot`.

* `fail_open` - Enable to allow traffic if the IPS buffer is full. Default is disable and IPS traffic is blocked when the IPS buffer is full. Valid values: `disable`, `enable`.

* `intelligent_mode` - Intelligent-Mode. Valid values: `disable`, `enable`.

* `ips_reserve_cpu` - Enable/disable IPS daemon's use of CPUs other than CPU 0. Valid values: `disable`, `enable`.

* `ngfw_max_scan_range` - NGFW policy-mode app detection threshold.
* `np_accel_mode` - Acceleration mode for IPS processing by NPx processors. Valid values: `none`, `basic`.

* `packet_log_queue_depth` - Packet/pcap log queue depth per IPS engine.
* `session_limit_mode` - Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics). Valid values: `accurate`, `heuristic`.

* `skype_client_public_ipaddr` - Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
* `socket_size` - IPS socket buffer size. Max and default value depend on available memory. Can be changed to tune performance.
* `sync_session_ttl` - Enable/disable use of kernel session TTL for IPS sessions. Valid values: `disable`, `enable`.

* `tls_active_probe` - Tls-Active-Probe. The structure of `tls_active_probe` block is documented below.
* `traffic_submit` - Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `disable`, `enable`.


The `tls_active_probe` block supports:

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `source_ip` - Source IP address used for TLS active probe.
* `source_ip6` - Source IPv6 address used for TLS active probe.
* `vdom` - Virtual domain name for TLS active probe.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Ips Global can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ips_global.labelname IpsGlobal
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

