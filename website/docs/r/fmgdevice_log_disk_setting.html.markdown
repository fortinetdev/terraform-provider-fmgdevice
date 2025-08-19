---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_disk_setting"
description: |-
  Settings for local disk logging.
---

# fmgdevice_log_disk_setting
Settings for local disk logging.

## Example Usage

```hcl
resource "fmgdevice_log_disk_setting" "trname" {
  diskfull                      = "nolog"
  dlp_archive_quota             = 10
  full_final_warning_threshold  = 10
  full_first_warning_threshold  = 10
  full_second_warning_threshold = 10
  device_name                   = var.device_name # not required if setting is at provider
  device_vdom                   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `diskfull` - Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `blocktraffic`, `nolog`.

* `dlp_archive_quota` - DLP archive quota (MB).
* `full_final_warning_threshold` - Log full final warning threshold as a percent (3 - 100, default = 95).
* `full_first_warning_threshold` - Log full first warning threshold as a percent (1 - 98, default = 75).
* `full_second_warning_threshold` - Log full second warning threshold as a percent (2 - 99, default = 90).
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ips_archive` - Enable/disable IPS packet archiving to the local disk. Valid values: `disable`, `enable`.

* `log_quota` - Disk log quota (MB).
* `max_log_file_size` - Maximum log file size before rolling (1 - 100 Mbytes).
* `max_policy_packet_capture_size` - Maximum size of policy sniffer in MB (0 means unlimited).
* `maximum_log_age` - Delete log files older than (days).
* `report_quota` - Report db quota (MB).
* `roll_day` - Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.

* `roll_schedule` - Frequency to check log file for rolling. Valid values: `daily`, `weekly`.

* `roll_time` - Time of day to roll the log file (hh:mm).
* `source_ip` - Source IP address to use for uploading disk log files.
* `status` - Enable/disable local disk logging. Valid values: `disable`, `enable`.

* `upload` - Enable/disable uploading log files when they are rolled. Valid values: `disable`, `enable`.

* `upload_delete_files` - Delete log files after uploading (default = enable). Valid values: `disable`, `enable`.

* `upload_destination` - The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`, `fortianalyzer`.

* `upload_ssl_conn` - Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.

* `uploaddir` - The remote directory on the FTP server to upload log files to.
* `uploadip` - IP address of the FTP server to upload log files to.
* `uploadpass` - Password required to log into the FTP server to upload disk log files.
* `uploadport` - TCP port to use for communicating with the FTP server (default = 21).
* `uploadsched` - Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.

* `uploadtime` - Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
* `uploadtype` - Types of log files to upload. Separate multiple entries with a space. Valid values: `traffic`, `event`, `virus`, `webfilter`, `attack`, `spamfilter`, `voip`, `dlp`, `app-ctrl`, `netscan`, `dlp-archive`, `IPS`, `anomaly`, `waf`, `gtp`, `dns`, `emailfilter`, `ssh`, `ssl`, `cifs`, `file-filter`, `icap`, `ztna`, `http`, `virtual-patch`.

* `uploaduser` - Username required to log into the FTP server to upload disk log files.
* `vrf_select` - VRF ID used for connection to server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log DiskSetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_disk_setting.labelname LogDiskSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

