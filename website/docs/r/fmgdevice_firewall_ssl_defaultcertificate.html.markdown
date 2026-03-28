---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_ssl_defaultcertificate"
description: |-
  Device FirewallSslDefaultCertificate
---

# fmgdevice_firewall_ssl_defaultcertificate
Device FirewallSslDefaultCertificate

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `default_ca` - Default-Ca.
* `default_server_cert` - Default-Server-Cert.
* `default_untrusted_ca` - Default-Untrusted-Ca.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall SslDefaultCertificate can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_ssl_defaultcertificate.labelname FirewallSslDefaultCertificate
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

