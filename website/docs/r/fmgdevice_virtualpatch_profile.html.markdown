---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_virtualpatch_profile"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure virtual-patch profile.
---

# fmgdevice_virtualpatch_profile
<i>This object will be purged after policy copy and install.</i> Configure virtual-patch profile.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `exemption`: `fmgdevice_virtualpatch_profile_exemption`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `action` - Action (pass/block). Valid values: `pass`, `block`.

* `comment` - Comment.
* `exemption` - Exemption. The structure of `exemption` block is documented below.
* `log` - Enable/disable logging of detection. Valid values: `disable`, `enable`.

* `name` - Profile name.
* `severity` - Relative severity of the signature (low, medium, high, critical). Valid values: `low`, `medium`, `high`, `critical`, `info`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `exemption` block supports:

* `device` - Device MAC addresses.
* `id` - IDs.
* `rule` - Patch signature rule IDs.
* `status` - Enable/disable exemption. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VirtualPatch Profile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_virtualpatch_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

