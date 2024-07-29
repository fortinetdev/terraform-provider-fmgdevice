---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ssh_setting"
description: |-
  SSH proxy settings.
---

# fmgdevice_firewall_ssh_setting
SSH proxy settings.

## Example Usage

```hcl
resource "fmgdevice_firewall_ssh_setting" "trname" {
  caname                = "your own value"
  host_trusted_checking = "enable"
  hostkey_dsa1024       = ["your own value"]
  hostkey_ecdsa256      = ["your own value"]
  hostkey_ecdsa384      = ["your own value"]
  device_name           = var.device_name # not required if setting is at provider
  device_vdom           = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `caname` - CA certificate used by SSH Inspection.
* `host_trusted_checking` - Enable/disable host trusted checking. Valid values: `disable`, `enable`.

* `hostkey_dsa1024` - DSA certificate used by SSH proxy.
* `hostkey_ecdsa256` - ECDSA nid256 certificate used by SSH proxy.
* `hostkey_ecdsa384` - ECDSA nid384 certificate used by SSH proxy.
* `hostkey_ecdsa521` - ECDSA nid384 certificate used by SSH proxy.
* `hostkey_ed25519` - ED25519 hostkey used by SSH proxy.
* `hostkey_rsa2048` - RSA certificate used by SSH proxy.
* `untrusted_caname` - Untrusted CA certificate used by SSH Inspection.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall SshSetting can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ssh_setting.labelname FirewallSshSetting
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

