---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_fortianalyzer3_setting"
description: |-
  Global FortiAnalyzer settings.
---

# fmgdevice_log_fortianalyzer3_setting
Global FortiAnalyzer settings.

## Example Usage

```hcl
resource "fmgdevice_log_fortianalyzer3_setting" "trname" {
  __change_ip              = 10
  access_config            = "enable"
  alt_server               = "your own value"
  certificate              = ["your own value"]
  certificate_verification = "enable"
  device_name              = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `__change_ip` - Hidden attribute.
* `access_config` - Enable/disable FortiAnalyzer access to configuration and data. Valid values: `disable`, `enable`.

* `alt_server` - Alternate FortiAnalyzer.
* `certificate` - Certificate used to communicate with FortiAnalyzer.
* `certificate_verification` - Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `disable`, `enable`.

* `conn_timeout` - FortiAnalyzer connection time-out in seconds (for status and log buffer).
* `enc_algorithm` - Configure the level of SSL protection for secure communication with FortiAnalyzer. Valid values: `default`, `high`, `low`, `disable`, `high-medium`, `low-medium`.

* `fallback_to_primary` - Enable/disable this FortiGate unit to fallback to the primary FortiAnalyzer when it is available. Valid values: `disable`, `enable`.

* `hmac_algorithm` - OFTP login hash algorithm. Valid values: `sha256`, `sha1`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ips_archive` - Enable/disable IPS packet archive logging. Valid values: `disable`, `enable`.

* `max_log_rate` - FortiAnalyzer maximum log rate in MBps (0 = unlimited).
* `monitor_failure_retry_period` - Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
* `monitor_keepalive_period` - Time between OFTP keepalives in seconds (for status and log buffer).
* `preshared_key` - Preshared-key used for auto-authorization on FortiAnalyzer.
* `priority` - Set log transmission priority. Valid values: `low`, `default`.

* `reliable` - Enable/disable reliable logging to FortiAnalyzer. Valid values: `disable`, `enable`.

* `serial` - Serial numbers of the FortiAnalyzer.
* `server` - The remote FortiAnalyzer.
* `server_cert_ca` - Mandatory CA on FortiGate in certificate chain of server.
* `source_ip` - Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1-3`.

* `status` - Enable/disable logging to FortiAnalyzer. Valid values: `disable`, `enable`.

* `upload_day` - Day of week (month) to upload logs.
* `upload_interval` - Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.

* `upload_option` - Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.

* `upload_time` - Time to upload logs (hh:mm).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log Fortianalyzer3Setting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_fortianalyzer3_setting.labelname LogFortianalyzer3Setting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

