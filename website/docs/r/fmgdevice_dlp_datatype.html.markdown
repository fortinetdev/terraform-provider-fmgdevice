---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_datatype"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure predefined data type used by DLP blocking.
---

# fmgdevice_dlp_datatype
<i>This object will be purged after policy copy and install.</i> Configure predefined data type used by DLP blocking.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Optional comments.
* `fgd_id` - ID of object in FortiGuard database.
* `look_ahead` - Number of characters to obtain in advance for verification (1 - 255, default = 1).
* `look_back` - Number of characters required to save for verification (1 - 255, default = 1).
* `match_ahead` - Number of characters behind for match-around (1 - 4096, default = 1).
* `match_around` - Dictionary to check whether it has a match around (Only support match-any and basic types, no repeat supported).
* `match_back` - Number of characters in front for match-around (1 - 4096, default = 1).
* `name` - Name of table containing the data type.
* `pattern` - Regular expression pattern string without look around.
* `transform` - Template to transform user input to a pattern using capture group from 'pattern'.
* `verify` - Regular expression pattern string used to verify the data type.
* `verify_transformed_pattern` - Enable/disable verification for transformed pattern. Valid values: `disable`, `enable`.

* `verify2` - Extra regular expression pattern string used to verify the data type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp DataType can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_datatype.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

