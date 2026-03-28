---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_address6_list"
description: |-
  <i>This object will be purged after policy copy and install.</i> IP address list.
---

# fmgdevice_firewall_address6_list
<i>This object will be purged after policy copy and install.</i> IP address list.

~> This resource is a sub resource for variable `list` of resource `fmgdevice_firewall_address6`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `address6` - Address6.

* `ip` - IP.
* `net_id` - Network ID.
* `obj_id` - Object ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{ip}}.

## Import

Firewall Address6List can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "address6=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_address6_list.labelname {{ip}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

