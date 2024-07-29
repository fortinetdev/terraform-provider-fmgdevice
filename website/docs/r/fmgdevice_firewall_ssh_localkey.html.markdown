---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ssh_localkey"
description: |-
  SSH proxy local keys.
---

# fmgdevice_firewall_ssh_localkey
SSH proxy local keys.

## Example Usage

```hcl
resource "fmgdevice_firewall_ssh_localkey" "trname" {
  name        = "your own value"
  password    = ["your own value"]
  private_key = "your own value"
  public_key  = "your own value"
  source      = "user"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `name` - SSH proxy local key name.
* `password` - Password for SSH private key.
* `private_key` - SSH proxy private key, encrypted with a password.
* `public_key` - SSH proxy public key.
* `source` - SSH proxy local key source type. Valid values: `built-in`, `user`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall SshLocalKey can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ssh_localkey.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

