---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ssl_settings_authenticationrule"
description: |-
  Authentication rule for SSL-VPN.
---

# fmgdevice_vpn_ssl_settings_authenticationrule
Authentication rule for SSL-VPN.

~> This resource is a sub resource for variable `authentication_rule` of resource `fmgdevice_vpn_ssl_settings`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_vpn_ssl_settings_authenticationrule" "trname" {
  fosid            = 3
  source_interface = ["any"]
  source_address   = ["all"]
  users            = ["guest"]
  portal           = ["web-access"]
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `auth` - SSL-VPN authentication method restriction. Valid values: `any`, `local`, `radius`, `ldap`, `tacacs+`, `peer`.

* `cipher` - SSL-VPN cipher strength. Valid values: `any`, `high`, `medium`.

* `client_cert` - Enable/disable SSL-VPN client certificate restrictive. Valid values: `disable`, `enable`.

* `groups` - User groups.
* `fosid` - ID (0 - 4294967295).
* `portal` - SSL-VPN portal.
* `realm` - SSL-VPN realm.
* `source_address` - Source address of incoming traffic.
* `source_address_negate` - Enable/disable negated source address match. Valid values: `disable`, `enable`.

* `source_address6` - IPv6 source address of incoming traffic.
* `source_address6_negate` - Enable/disable negated source IPv6 address match. Valid values: `disable`, `enable`.

* `source_interface` - SSL-VPN source interface of incoming traffic.
* `user_peer` - Name of user peer.
* `users` - User name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Vpn SslSettingsAuthenticationRule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ssl_settings_authenticationrule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

