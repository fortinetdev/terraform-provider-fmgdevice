---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_sctpfilter_profile_ppidfilters"
description: |-
  <i>This object will be purged after policy copy and install.</i> PPID filters list.
---

# fmgdevice_sctpfilter_profile_ppidfilters
<i>This object will be purged after policy copy and install.</i> PPID filters list.

~> This resource is a sub resource for variable `ppid_filters` of resource `fmgdevice_sctpfilter_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `action` - Action taken when PPID is matched. Valid values: `pass`, `reset`, `replace`.

* `comment` - Comment.
* `fosid` - ID.
* `ppid` - Payload protocol identifier.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SctpFilter ProfilePpidFilters can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_sctpfilter_profile_ppidfilters.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

