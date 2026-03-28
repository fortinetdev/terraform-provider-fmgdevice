---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_dictionary_entries"
description: |-
  <i>This object will be purged after policy copy and install.</i> DLP dictionary entries.
---

# fmgdevice_dlp_dictionary_entries
<i>This object will be purged after policy copy and install.</i> DLP dictionary entries.

~> This resource is a sub resource for variable `entries` of resource `fmgdevice_dlp_dictionary`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `dictionary` - Dictionary.

* `comment` - Optional comments.
* `fosid` - ID.
* `ignore_case` - Enable/disable ignore case. Valid values: `disable`, `enable`.

* `pattern` - Pattern to match.
* `repeat` - Enable/disable repeat match. Valid values: `disable`, `enable`.

* `status` - Enable/disable this pattern. Valid values: `disable`, `enable`.

* `type` - Pattern type to match.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Dlp DictionaryEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "dictionary=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_dictionary_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

