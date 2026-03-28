---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_wildcardfqdn_group"
description: |-
  <i>This object will be purged after policy copy and install.</i> Config global Wildcard FQDN address groups.
---

# fmgdevice_firewall_wildcardfqdn_group
<i>This object will be purged after policy copy and install.</i> Config global Wildcard FQDN address groups.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `color` - GUI icon color.
* `comment` - Comment.
* `member` - Address group members.
* `name` - Address group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall WildcardFqdnGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_wildcardfqdn_group.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

