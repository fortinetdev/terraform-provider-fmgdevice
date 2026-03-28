---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_profile_antiphish_inspectionentries"
description: |-
  <i>This object will be purged after policy copy and install.</i> AntiPhishing entries.
---

# fmgdevice_webfilter_profile_antiphish_inspectionentries
<i>This object will be purged after policy copy and install.</i> AntiPhishing entries.

~> This resource is a sub resource for variable `inspection_entries` of resource `fmgdevice_webfilter_profile_antiphish`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `action` - Action to be taken upon an AntiPhishing match. Valid values: `log`, `block`, `exempt`.

* `fortiguard_category` - FortiGuard category to match.
* `name` - Inspection target name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Webfilter ProfileAntiphishInspectionEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_profile_antiphish_inspectionentries.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

