---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_dictionary"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure dictionaries used by DLP blocking.
---

# fmgdevice_dlp_dictionary
<i>This object will be purged after policy copy and install.</i> Configure dictionaries used by DLP blocking.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fmgdevice_dlp_dictionary_entries`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Optional comments.
* `entries` - Entries. The structure of `entries` block is documented below.
* `fgd_id` - ID of object in FortiGuard database.
* `match_around` - Enable/disable match-around support. Valid values: `disable`, `enable`.

* `match_type` - Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`.

* `name` - Name of table containing the dictionary.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `comment` - Optional comments.
* `id` - ID.
* `ignore_case` - Enable/disable ignore case. Valid values: `disable`, `enable`.

* `pattern` - Pattern to match.
* `repeat` - Enable/disable repeat match. Valid values: `disable`, `enable`.

* `status` - Enable/disable this pattern. Valid values: `disable`, `enable`.

* `type` - Pattern type to match.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp Dictionary can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_dictionary.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

