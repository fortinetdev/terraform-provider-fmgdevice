---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_fortiguard_setting"
description: |-
  Configure logging to FortiCloud.
---

# fmgdevice_log_fortiguard_setting
Configure logging to FortiCloud.

## Example Usage

```hcl
resource "fmgdevice_log_fortiguard_setting" "trname" {
  access_config           = "enable"
  conn_timeout            = 10
  enc_algorithm           = "disable"
  interface               = ["port2"]
  interface_select_method = "auto"
  device_name             = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `access_config` - Enable/disable FortiCloud access to configuration and data. Valid values: `disable`, `enable`.

* `conn_timeout` - FortiGate Cloud connection timeout in seconds.
* `enc_algorithm` - Configure the level of SSL protection for secure communication with FortiCloud. Valid values: `default`, `high`, `low`, `disable`, `high-medium`, `low-medium`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `max_log_rate` - FortiCloud maximum log rate in MBps (0 = unlimited).
* `priority` - Set log transmission priority. Valid values: `low`, `default`.

* `source_ip` - Source IP address used to connect FortiCloud.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1`, `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1-3`.

* `status` - Enable/disable logging to FortiCloud. Valid values: `disable`, `enable`.

* `upload_day` - Day of week to roll logs.
* `upload_interval` - Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.

* `upload_option` - Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.

* `upload_time` - Time of day to roll logs (hh:mm).
* `vrf_select` - VRF ID used for connection to server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log FortiguardSetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_fortiguard_setting.labelname LogFortiguardSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

