---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_sslsshprofile_sslclientcertificate"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_firewall_sslsshprofile_sslclientcertificate
<i>This object will be purged after policy copy and install.</i>

~> This resource is a sub resource for variable `ssl_client_certificate` of resource `fmgdevice_firewall_sslsshprofile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `ssl_ssh_profile` - Ssl Ssh Profile.

* `caname` - Caname.
* `cert` - Cert.
* `keyring_list` - Keyring-List.
* `status` - Status. Valid values: `do-not-offer`, `keyring-list`, `ca-sign`, `static`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall SslSshProfileSslClientCertificate can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "ssl_ssh_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_sslsshprofile_sslclientcertificate.labelname FirewallSslSshProfileSslClientCertificate
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

