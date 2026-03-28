---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ssh_localca"
description: |-
  <i>This object will be purged after policy copy and install.</i> SSH proxy local CA.
---

# fmgdevice_firewall_ssh_localca
<i>This object will be purged after policy copy and install.</i> SSH proxy local CA.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `name` - SSH proxy local CA name.
* `password` - Password for SSH private key.
* `private_key` - SSH proxy private key, encrypted with a password.
* `public_key` - SSH proxy public key.
* `source` - SSH proxy local CA source type. Valid values: `built-in`, `user`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall SshLocalCa can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ssh_localca.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

