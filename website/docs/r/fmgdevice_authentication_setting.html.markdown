---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_authentication_setting"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure authentication setting.
---

# fmgdevice_authentication_setting
<i>This object will be purged after policy copy and install.</i> Configure authentication setting.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `active_auth_scheme` - Active authentication method (scheme name).
* `auth_https` - Enable/disable redirecting HTTP user authentication to HTTPS. Valid values: `disable`, `enable`.

* `captive_portal` - Captive portal host name.
* `captive_portal_ip` - Captive portal IP address.
* `captive_portal_ip6` - Captive portal IPv6 address.
* `captive_portal_port` - Captive portal port number (1 - 65535, default = 7830).
* `captive_portal_ssl_port` - Captive portal SSL port number (1 - 65535, default = 7831).
* `captive_portal_type` - Captive portal type. Valid values: `fqdn`, `ip`.

* `captive_portal6` - IPv6 captive portal host name.
* `cert_auth` - Enable/disable redirecting certificate authentication to HTTPS portal. Valid values: `disable`, `enable`.

* `cert_captive_portal` - Certificate captive portal host name.
* `cert_captive_portal_ip` - Certificate captive portal IP address.
* `cert_captive_portal_port` - Certificate captive portal port number (1 - 65535, default = 7832).
* `cookie_max_age` - Persistent web portal cookie maximum age in minutes (30 - 10080 (1 week), default = 480 (8 hours)).
* `cookie_refresh_div` - Refresh rate divider of persistent web portal cookie (default = 2). Refresh value = cookie-max-age/cookie-refresh-div.
* `dev_range` - Address range for the IP based device query.
* `ip_auth_cookie` - Enable/disable persistent cookie on IP based web portal authentication (default = disable). Valid values: `disable`, `enable`.

* `persistent_cookie` - Enable/disable persistent cookie on web portal authentication (default = enable). Valid values: `disable`, `enable`.

* `sso_auth_scheme` - Single-Sign-On authentication method (scheme name).
* `update_time` - Time of the last update.
* `user_cert_ca` - CA certificate used for client certificate verification.
* `log_auth_request` - Log-Auth-Request. Valid values: `disable`, `enable`.

* `rewrite_https_port` - Rewrite-Https-Port.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Authentication Setting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_authentication_setting.labelname AuthenticationSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

