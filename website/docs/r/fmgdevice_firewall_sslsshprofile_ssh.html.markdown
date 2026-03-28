---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_sslsshprofile_ssh"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure SSH options.
---

# fmgdevice_firewall_sslsshprofile_ssh
<i>This object will be purged after policy copy and install.</i> Configure SSH options.

~> This resource is a sub resource for variable `ssh` of resource `fmgdevice_firewall_sslsshprofile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `ssl_ssh_profile` - Ssl Ssh Profile.

* `inspect_all` - Level of SSL inspection. Valid values: `disable`, `deep-inspection`.

* `ports` - Ports to use for scanning (1 - 65535, default = 443).
* `proxy_after_tcp_handshake` - Proxy traffic after the TCP 3-way handshake has been established (not before). Valid values: `disable`, `enable`.

* `ssh_algorithm` - Relative strength of encryption algorithms accepted during negotiation. Valid values: `compatible`, `high-encryption`.

* `ssh_tun_policy_check` - Enable/disable SSH tunnel policy check. Valid values: `disable`, `enable`.

* `status` - Configure protocol inspection status. Valid values: `disable`, `deep-inspection`.

* `unsupported_version` - Action based on SSH version being unsupported. Valid values: `block`, `bypass`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall SslSshProfileSsh can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "ssl_ssh_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_sslsshprofile_ssh.labelname FirewallSslSshProfileSsh
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

