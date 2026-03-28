---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_isolator_profile_entries"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_isolator_profile_entries
<i>This object will be purged after policy copy and install.</i>

~> This resource is a sub resource for variable `entries` of resource `fmgdevice_isolator_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `action` - Action. Valid values: `isolate`, `freeze`, `block`, `allow`.

* `copy_paste` - Copy-Paste. Valid values: `disable`, `enable`.

* `fosid` - Id.
* `proxy_address` - Proxy-Address.
* `right_click` - Right-Click. Valid values: `disable`, `enable`.

* `status` - Status. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Isolator ProfileEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_isolator_profile_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

