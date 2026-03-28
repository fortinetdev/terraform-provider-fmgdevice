---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_sslsshprofile_sslexempt"
description: |-
  <i>This object will be purged after policy copy and install.</i> Servers to exempt from SSL inspection.
---

# fmgdevice_firewall_sslsshprofile_sslexempt
<i>This object will be purged after policy copy and install.</i> Servers to exempt from SSL inspection.

~> This resource is a sub resource for variable `ssl_exempt` of resource `fmgdevice_firewall_sslsshprofile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `ssl_ssh_profile` - Ssl Ssh Profile.

* `address` - IPv4 address object.
* `address6` - IPv6 address object.
* `fortiguard_category` - FortiGuard category ID.
* `fosid` - ID number.
* `regex` - Exempt servers by regular expression.
* `type` - Type of address object (IPv4 or IPv6) or FortiGuard category. Valid values: `fortiguard-category`, `address`, `address6`, `wildcard-fqdn`, `regex`.

* `wildcard_fqdn` - Exempt servers by wildcard FQDN.
* `finger_print_category` - Finger-Print-Category. Valid values: `unknown`, `firefox`, `chrome`, `safari`, `edge`, `ie`, `android`, `ios`, `windows`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall SslSshProfileSslExempt can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "ssl_ssh_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_sslsshprofile_sslexempt.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

