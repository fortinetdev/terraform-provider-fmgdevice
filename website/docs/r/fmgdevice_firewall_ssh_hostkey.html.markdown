---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ssh_hostkey"
description: |-
  SSH proxy host public keys.
---

# fmgdevice_firewall_ssh_hostkey
SSH proxy host public keys.

## Example Usage

```hcl
resource "fmgdevice_firewall_ssh_hostkey" "trname" {
  hostname    = "your own value"
  ip          = "your own value"
  name        = "your own value"
  nid         = "521"
  port        = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `hostname` - Hostname of the SSH server to match SSH certificate principals.
* `ip` - IP address of the SSH server.
* `name` - SSH public key name.
* `nid` - Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.

* `port` - Port of the SSH server.
* `public_key` - SSH public key.
* `status` - Set the trust status of the public key. Valid values: `trusted`, `revoked`.

* `type` - Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.

* `usage` - Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall SshHostKey can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ssh_hostkey.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

