---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_sslsshprofile_echoutersni"
description: |-
  <i>This object will be purged after policy copy and install.</i> ClientHelloOuter SNIs to be blocked.
---

# fmgdevice_firewall_sslsshprofile_echoutersni
<i>This object will be purged after policy copy and install.</i> ClientHelloOuter SNIs to be blocked.

~> This resource is a sub resource for variable `ech_outer_sni` of resource `fmgdevice_firewall_sslsshprofile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `ssl_ssh_profile` - Ssl Ssh Profile.

* `name` - ClientHelloOuter SNI name.
* `sni` - ClientHelloOuter SNI to be blocked.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall SslSshProfileEchOuterSni can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "ssl_ssh_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_sslsshprofile_echoutersni.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

