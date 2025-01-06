---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_syslogd3_overridesetting"
description: |-
  Override settings for remote syslog server.
---

# fmgdevice_log_syslogd3_overridesetting
Override settings for remote syslog server.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `custom_field_name`: `fmgdevice_log_syslogd3_overridesetting_customfieldname`



## Example Usage

```hcl
resource "fmgdevice_log_syslogd3_overridesetting" "trname" {
  certificate = ["your own value"]
  custom_field_name {
    custom = "your own value"
    fosid  = 10
    name   = "your own value"
  }

  enc_algorithm = "high"
  facility      = "ftp"
  format        = "rfc5424"
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `certificate` - Certificate used to communicate with Syslog server.
* `custom_field_name` - Custom-Field-Name. The structure of `custom_field_name` block is documented below.
* `enc_algorithm` - Enable/disable reliable syslogging with TLS encryption. Valid values: `high`, `low`, `disable`, `high-medium`.

* `facility` - Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.

* `format` - Log format. Valid values: `default`, `csv`, `cef`, `rfc5424`, `json`.

* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `max_log_rate` - Syslog maximum log rate in MBps (0 = unlimited).
* `mode` - Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.

* `override` - Enable/disable override syslog settings. Valid values: `disable`, `enable`.

* `port` - Server listen port.
* `priority` - Set log transmission priority. Valid values: `low`, `default`.

* `server` - Address of remote syslog server.
* `source_ip` - Source IP address of syslog.
* `source_ip_interface` - Source interface of syslog.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1`, `TLSv1-3`.

* `status` - Enable/disable remote syslog logging. Valid values: `disable`, `enable`.

* `use_management_vdom` - Enable/disable use of management VDOM as source VDOM for logs sent to syslog server. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `custom_field_name` block supports:

* `custom` - Field custom name [A-Za-z0-9_].
* `id` - Entry ID.
* `name` - Field name [A-Za-z0-9_].


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log Syslogd3OverrideSetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_syslogd3_overridesetting.labelname LogSyslogd3OverrideSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

