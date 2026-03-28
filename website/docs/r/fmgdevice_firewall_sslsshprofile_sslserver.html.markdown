---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_sslsshprofile_sslserver"
description: |-
  <i>This object will be purged after policy copy and install.</i> SSL server settings used for client certificate request.
---

# fmgdevice_firewall_sslsshprofile_sslserver
<i>This object will be purged after policy copy and install.</i> SSL server settings used for client certificate request.

~> This resource is a sub resource for variable `ssl_server` of resource `fmgdevice_firewall_sslsshprofile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `ssl_ssh_profile` - Ssl Ssh Profile.

* `ftps_client_certificate` - Action based on received client certificate during the FTPS handshake. Valid values: `bypass`, `inspect`, `block`.

* `https_client_certificate` - Action based on received client certificate during the HTTPS handshake. Valid values: `bypass`, `inspect`, `block`.

* `fosid` - SSL server ID.
* `imaps_client_certificate` - Action based on received client certificate during the IMAPS handshake. Valid values: `bypass`, `inspect`, `block`.

* `ip` - IPv4 address of the SSL server.
* `pop3s_client_certificate` - Action based on received client certificate during the POP3S handshake. Valid values: `bypass`, `inspect`, `block`.

* `smtps_client_certificate` - Action based on received client certificate during the SMTPS handshake. Valid values: `bypass`, `inspect`, `block`.

* `ssl_other_client_certificate` - Action based on received client certificate during an SSL protocol handshake. Valid values: `bypass`, `inspect`, `block`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall SslSshProfileSslServer can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "ssl_ssh_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_sslsshprofile_sslserver.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

