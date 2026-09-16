---
subcategory: "Firewall"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_customtag"
description: |-
  <i>This object will be purged after policy copy and install.</i> Define custom tag table.
---

# fmgdevice_firewall_customtag
<i>This object will be purged after policy copy and install.</i> Define custom tag table.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `abbreviation` - Custom tag name abbreviation.
* `color` - Color of the custom tag on the GUI.
* `comment` - Comment.
* `fabric_force_sync` - Enable/disable forced synchronization of configuration objects from the root FortiGate unit to the downstream devices.  Configuration conflict check is skipped. Valid values: `disable`, `enable`.

* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `fabric_object_source` - Source of truth for fabric object. Valid values: `member`, `local`, `root`.

* `name` - Custom tag name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall CustomTag can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_customtag.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

