---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_multicastaddress"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure multicast addresses.
---

# fmgdevice_firewall_multicastaddress
<i>This object will be purged after policy copy and install.</i> Configure multicast addresses.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `tagging`: `fmgdevice_firewall_multicastaddress_tagging`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `associated_interface` - Interface associated with the address object. When setting up a policy, only addresses associated with this interface are available.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `comment` - Comment.
* `end_ip` - Final IPv4 address (inclusive) in the range for the address.
* `name` - Multicast address name.
* `start_ip` - First IPv4 address (inclusive) in the range for the address.
* `subnet` - Broadcast address and subnet.
* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `type` - Type of address object: multicast IP address range or broadcast IP/mask to be treated as a multicast address. Valid values: `multicastrange`, `broadcastmask`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `tagging` block supports:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall MulticastAddress can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_multicastaddress.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

