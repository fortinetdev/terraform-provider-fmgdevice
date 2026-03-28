---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_isolator_profile"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_isolator_profile
<i>This object will be purged after policy copy and install.</i>

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fmgdevice_isolator_profile_entries`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comments` - Comments.
* `disclaimer` - Disclaimer.
* `entries` - Entries. The structure of `entries` block is documented below.
* `name` - Name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `action` - Action. Valid values: `isolate`, `freeze`, `block`, `allow`.

* `copy_paste` - Copy-Paste. Valid values: `disable`, `enable`.

* `id` - Id.
* `proxy_address` - Proxy-Address.
* `right_click` - Right-Click. Valid values: `disable`, `enable`.

* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Isolator Profile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_isolator_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

