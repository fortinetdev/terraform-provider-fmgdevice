---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ssl_settings_authenticationrule_move"
description: |-
  Authentication rule for SSL-VPN.
---

# fmgdevice_vpn_ssl_settings_authenticationrule_move
Authentication rule for SSL-VPN.

## Example Usage

```hcl
resource "fmgdevice_vpn_ssl_settings_authenticationrule_move" "trname" {
  target              = fmgdevice_vpn_ssl_settings_authenticationrule.trname2.fosid
  authentication_rule = fmgdevice_vpn_ssl_settings_authenticationrule.trname3.fosid
  option              = "after"
  device_name         = var.device_name # not required if setting is at provider
  device_vdom         = var.device_vdom # not required if setting is at provider
  depends_on          = [fmgdevice_vpn_ssl_settings_authenticationrule.trname2, fmgdevice_vpn_ssl_settings_authenticationrule.trname3]
}

resource "fmgdevice_vpn_ssl_settings_authenticationrule" "trname2" {
  fosid            = 5
  source_interface = ["any"]
  source_address   = ["all"]
  users            = ["guest"]
  portal           = ["web-access"]
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
}

resource "fmgdevice_vpn_ssl_settings_authenticationrule" "trname3" {
  fosid            = 6
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
* `authentication_rule` - Authentication Rule.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the authentication rule changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of authentication rules.
