---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_log_setting"
description: |-
  Configure general log settings.
---

# fmgdevice_log_setting
Configure general log settings.

## Example Usage

```hcl
resource "fmgdevice_log_setting" "trname" {
  anonymization_hash    = "your own value"
  brief_traffic_format  = "disable"
  custom_log_fields     = ["your own value"]
  daemon_log            = "disable"
  expolicy_implicit_log = "disable"
  device_name           = var.device_name # not required if setting is at provider
  device_vdom           = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `anonymization_hash` - User name anonymization hash salt.
* `brief_traffic_format` - Enable/disable brief format traffic logging. Valid values: `disable`, `enable`.

* `custom_log_fields` - Custom fields to append to all log messages.
* `daemon_log` - Enable/disable daemon logging. Valid values: `disable`, `enable`.

* `expolicy_implicit_log` - Enable/disable explicit proxy firewall implicit policy logging. Valid values: `disable`, `enable`.

* `extended_log` - Enable/disable extended traffic logging. Valid values: `disable`, `enable`.

* `extended_utm_log` - Enable/disable extended UTM logging. Valid values: `disable`, `enable`.

* `faz_override` - Enable/disable override FortiAnalyzer settings. Valid values: `disable`, `enable`.

* `fortiview_weekly_data` - Enable/disable FortiView weekly data. Valid values: `disable`, `enable`.

* `fwpolicy_implicit_log` - Enable/disable implicit firewall policy logging. Valid values: `disable`, `enable`.

* `fwpolicy6_implicit_log` - Enable/disable implicit firewall policy6 logging. Valid values: `disable`, `enable`.

* `local_in_allow` - Enable/disable local-in-allow logging. Valid values: `disable`, `enable`.

* `local_in_deny_broadcast` - Enable/disable local-in-deny-broadcast logging. Valid values: `disable`, `enable`.

* `local_in_deny_unicast` - Enable/disable local-in-deny-unicast logging. Valid values: `disable`, `enable`.

* `local_in_policy_log` - Enable/disable local-in-policy logging. Valid values: `disable`, `enable`.

* `local_out` - Enable/disable local-out logging. Valid values: `disable`, `enable`.

* `local_out_ioc_detection` - Enable/disable local-out traffic IoC detection. Requires local-out to be enabled. Valid values: `disable`, `enable`.

* `log_invalid_packet` - Enable/disable invalid packet traffic logging. Valid values: `disable`, `enable`.

* `log_policy_comment` - Enable/disable inserting policy comments into traffic logs. Valid values: `disable`, `enable`.

* `log_policy_name` - Enable/disable inserting policy name into traffic logs. Valid values: `disable`, `enable`.

* `log_user_in_upper` - Enable/disable logs with user-in-upper. Valid values: `disable`, `enable`.

* `long_live_session_stat` - Enable/disable long-live-session statistics logging. Valid values: `disable`, `enable`.

* `neighbor_event` - Enable/disable neighbor event logging. Valid values: `disable`, `enable`.

* `resolve_ip` - Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `disable`, `enable`.

* `resolve_port` - Enable/disable adding resolved service names to traffic logs. Valid values: `disable`, `enable`.

* `rest_api_get` - Enable/disable REST API GET request logging. Valid values: `disable`, `enable`.

* `rest_api_set` - Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `disable`, `enable`.

* `syslog_override` - Enable/disable override Syslog settings. Valid values: `disable`, `enable`.

* `user_anonymize` - Enable/disable anonymizing user names in log messages. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log Setting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_log_setting.labelname LogSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

