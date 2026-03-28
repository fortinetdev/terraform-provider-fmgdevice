---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_authentication_rule"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure Authentication Rules.
---

# fmgdevice_authentication_rule
<i>This object will be purged after policy copy and install.</i> Configure Authentication Rules.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `active_auth_method` - Select an active authentication method.
* `cert_auth_cookie` - Enable/disable to use device certificate as authentication cookie (default = enable). Valid values: `disable`, `enable`.

* `comments` - Comment.
* `cors_depth` - Depth to allow CORS access (default = 3).
* `cors_stateful` - Enable/disable allowance of CORS access (default = disable). Valid values: `disable`, `enable`.

* `dstaddr` - Select an IPv4 destination address from available options. Required for web proxy authentication.
* `dstaddr6` - Select an IPv6 destination address from available options. Required for web proxy authentication.
* `ip_based` - Enable/disable IP-based authentication. When enabled, previously authenticated users from the same IP address will be exempted. Valid values: `disable`, `enable`.

* `name` - Authentication rule name.
* `protocol` - Authentication is required for the selected protocol (default = HTTP). Valid values: `http`, `ftp`, `socks`, `ssh`, `ztna-portal`.

* `srcaddr` - Authentication is required for the selected IPv4 source address.
* `srcaddr6` - Authentication is required for the selected IPv6 source address.
* `srcintf` - Incoming (ingress) interface.
* `sso_auth_method` - Select a single-sign on (SSO) authentication method.
* `status` - Enable/disable this authentication rule. Valid values: `disable`, `enable`.

* `transaction_based` - Enable/disable transaction based authentication (default = disable). Valid values: `disable`, `enable`.

* `web_auth_cookie` - Enable/disable Web authentication cookies (default = disable). Valid values: `disable`, `enable`.

* `web_portal` - Enable/disable web portal for proxy transparent policy (default = enable). Valid values: `disable`, `enable`.

* `form_auth_fallback` - Form-Auth-Fallback. Valid values: `disable`, `enable`.

* `web_proxy` - Web-Proxy.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Authentication Rule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_authentication_rule.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

