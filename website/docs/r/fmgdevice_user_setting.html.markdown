---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_setting"
description: |-
  Configure user authentication setting.
---

# fmgdevice_user_setting
Configure user authentication setting.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `auth_ports`: `fmgdevice_user_setting_authports`



## Example Usage

```hcl
resource "fmgdevice_user_setting" "trname" {
  auth_blackout_time = 10
  auth_ca_cert       = ["your own value"]
  auth_cert          = ["your own value"]
  auth_http_basic    = "enable"
  auth_invalid_max   = 10
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auth_blackout_time` - Time in seconds an IP address is denied access after failing to authenticate five times within one minute.
* `auth_ca_cert` - HTTPS CA certificate for policy authentication.
* `auth_cert` - HTTPS server certificate for policy authentication.
* `auth_http_basic` - Enable/disable use of HTTP basic authentication for identity-based firewall policies. Valid values: `disable`, `enable`.

* `auth_invalid_max` - Maximum number of failed authentication attempts before the user is blocked.
* `auth_lockout_duration` - Lockout period in seconds after too many login failures.
* `auth_lockout_threshold` - Maximum number of failed login attempts before login lockout is triggered.
* `auth_on_demand` - Always/implicitly trigger firewall authentication on demand. Valid values: `always`, `implicitly`.

* `auth_portal_timeout` - Time in minutes before captive portal user have to re-authenticate (1 - 30 min, default 3 min).
* `auth_ports` - Auth-Ports. The structure of `auth_ports` block is documented below.
* `auth_secure_http` - Enable/disable redirecting HTTP user authentication to more secure HTTPS. Valid values: `disable`, `enable`.

* `auth_src_mac` - Enable/disable source MAC for user identity. Valid values: `disable`, `enable`.

* `auth_ssl_allow_renegotiation` - Allow/forbid SSL re-negotiation for HTTPS authentication. Valid values: `disable`, `enable`.

* `auth_ssl_max_proto_version` - Maximum supported protocol version for SSL/TLS connections (default is no limit). Valid values: `tlsv1-1`, `tlsv1-2`, `sslv3`, `tlsv1`, `tlsv1-3`.

* `auth_ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `TLSv1-1`, `TLSv1-2`, `SSLv3`, `TLSv1`, `TLSv1-3`.

* `auth_ssl_sigalgs` - Set signature algorithms related to HTTPS authentication (affects TLS version &lt;= 1.2 only, default is to enable all). Valid values: `no-rsa-pss`, `all`.

* `auth_timeout` - Time in minutes before the firewall user authentication timeout requires the user to re-authenticate.
* `auth_timeout_type` - Control if authenticated users have to login again after a hard timeout, after an idle timeout, or after a session timeout. Valid values: `idle-timeout`, `hard-timeout`, `new-session`.

* `auth_type` - Supported firewall policy authentication protocols/methods. Valid values: `http`, `https`, `ftp`, `telnet`.

* `default_user_password_policy` - Default password policy to apply to all local users unless otherwise specified, as defined in config user password-policy.
* `per_policy_disclaimer` - Enable/disable per policy disclaimer. Valid values: `disable`, `enable`.

* `radius_ses_timeout_act` - Set the RADIUS session timeout to a hard timeout or to ignore RADIUS server session timeouts. Valid values: `hard-timeout`, `ignore-timeout`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `auth_ports` block supports:

* `id` - ID.
* `port` - Non-standard port for firewall user authentication.
* `type` - Service type. Valid values: `http`, `https`, `ftp`, `telnet`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

User Setting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_setting.labelname UserSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

