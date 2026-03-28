---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_sctpfilter_profile"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure SCTP filter profiles.
---

# fmgdevice_sctpfilter_profile
<i>This object will be purged after policy copy and install.</i> Configure SCTP filter profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `ppid_filters`: `fmgdevice_sctpfilter_profile_ppidfilters`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `name` - Profile name.
* `ppid_filters` - Ppid-Filters. The structure of `ppid_filters` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ppid_filters` block supports:

* `action` - Action taken when PPID is matched. Valid values: `pass`, `reset`, `replace`.

* `comment` - Comment.
* `id` - ID.
* `ppid` - Payload protocol identifier.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SctpFilter Profile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_sctpfilter_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

