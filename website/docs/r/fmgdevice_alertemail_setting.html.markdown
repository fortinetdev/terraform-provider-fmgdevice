---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_alertemail_setting"
description: |-
  Configure alert email settings.
---

# fmgdevice_alertemail_setting
Configure alert email settings.

## Example Usage

```hcl
resource "fmgdevice_alertemail_setting" "trname" {
  fds_license_expiring_days    = 10
  fds_license_expiring_warning = "enable"
  fds_update_logs              = "enable"
  fips_cc_errors               = "disable"
  fsso_disconnect_logs         = "disable"
  device_name                  = var.device_name # not required if setting is at provider
  device_vdom                  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `fds_license_expiring_days` - Number of days to send alert email prior to FortiGuard license expiration (1 - 100 days, default = 15).
* `fds_license_expiring_warning` - Enable/disable FortiGuard license expiration warnings in alert email. Valid values: `disable`, `enable`.

* `fds_update_logs` - Enable/disable FortiGuard update logs in alert email. Valid values: `disable`, `enable`.

* `fips_cc_errors` - Enable/disable FIPS and Common Criteria error logs in alert email. Valid values: `disable`, `enable`.

* `fsso_disconnect_logs` - Enable/disable logging of FSSO collector agent disconnect. Valid values: `disable`, `enable`.

* `ha_logs` - Enable/disable HA logs in alert email. Valid values: `disable`, `enable`.

* `ips_logs` - Enable/disable IPS logs in alert email. Valid values: `disable`, `enable`.

* `ipsec_errors_logs` - Enable/disable IPsec error logs in alert email. Valid values: `disable`, `enable`.

* `ppp_errors_logs` - Enable/disable PPP error logs in alert email. Valid values: `disable`, `enable`.

* `admin_login_logs` - Enable/disable administrator login/logout logs in alert email. Valid values: `disable`, `enable`.

* `alert_interval` - Alert alert interval in minutes.
* `amc_interface_bypass_mode` - Enable/disable Fortinet Advanced Mezzanine Card (AMC) interface bypass mode logs in alert email. Valid values: `disable`, `enable`.

* `antivirus_logs` - Enable/disable antivirus logs in alert email. Valid values: `disable`, `enable`.

* `configuration_changes_logs` - Enable/disable configuration change logs in alert email. Valid values: `disable`, `enable`.

* `critical_interval` - Critical alert interval in minutes.
* `debug_interval` - Debug alert interval in minutes.
* `email_interval` - Interval between sending alert emails (1 - 99999 min, default = 5).
* `emergency_interval` - Emergency alert interval in minutes.
* `error_interval` - Error alert interval in minutes.
* `filter_mode` - How to filter log messages that are sent to alert emails. Valid values: `category`, `threshold`.

* `firewall_authentication_failure_logs` - Enable/disable firewall authentication failure logs in alert email. Valid values: `disable`, `enable`.

* `fortiguard_log_quota_warning` - Enable/disable FortiCloud log quota warnings in alert email. Valid values: `disable`, `enable`.

* `information_interval` - Information alert interval in minutes.
* `local_disk_usage` - Disk usage percentage at which to send alert email (1 - 99 percent, default = 75).
* `log_disk_usage_warning` - Enable/disable disk usage warnings in alert email. Valid values: `disable`, `enable`.

* `mailto1` - Email address to send alert email to (usually a system administrator) (max. 63 characters).
* `mailto2` - Optional second email address to send alert email to (max. 63 characters).
* `mailto3` - Optional third email address to send alert email to (max. 63 characters).
* `notification_interval` - Notification alert interval in minutes.
* `severity` - Lowest severity level to log. Valid values: `emergency`, `alert`, `critical`, `error`, `warning`, `notification`, `information`, `debug`.

* `ssh_logs` - Enable/disable SSH logs in alert email. Valid values: `disable`, `enable`.

* `sslvpn_authentication_errors_logs` - Enable/disable SSL-VPN authentication error logs in alert email. Valid values: `disable`, `enable`.

* `username` - Name that appears in the From: field of alert emails (max. 63 characters).
* `violation_traffic_logs` - Enable/disable violation traffic logs in alert email. Valid values: `disable`, `enable`.

* `warning_interval` - Warning alert interval in minutes.
* `webfilter_logs` - Enable/disable web filter logs in alert email. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Alertemail Setting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_alertemail_setting.labelname AlertemailSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

