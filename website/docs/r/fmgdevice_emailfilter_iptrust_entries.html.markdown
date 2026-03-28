---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_emailfilter_iptrust_entries"
description: |-
  Spam filter trusted IP addresses.
---

# fmgdevice_emailfilter_iptrust_entries
Spam filter trusted IP addresses.

~> This resource is a sub resource for variable `entries` of resource `fmgdevice_emailfilter_iptrust`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `iptrust` - Iptrust.

* `addr_type` - Type of address. Valid values: `ipv4`, `ipv6`.

* `fosid` - Trusted IP entry ID.
* `ip4_subnet` - IPv4 network address or network address/subnet mask bits.
* `ip6_subnet` - IPv6 network address/subnet mask bits.
* `status` - Enable/disable status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Emailfilter IptrustEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "iptrust=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_emailfilter_iptrust_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

