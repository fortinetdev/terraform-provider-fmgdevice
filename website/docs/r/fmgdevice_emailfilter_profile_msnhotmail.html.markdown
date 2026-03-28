---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_emailfilter_profile_msnhotmail"
description: |-
  <i>This object will be purged after policy copy and install.</i> MSN Hotmail.
---

# fmgdevice_emailfilter_profile_msnhotmail
<i>This object will be purged after policy copy and install.</i> MSN Hotmail.

~> This resource is a sub resource for variable `msn_hotmail` of resource `fmgdevice_emailfilter_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `log` - Log. Valid values: `disable`, `enable`.

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Emailfilter ProfileMsnHotmail can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_emailfilter_profile_msnhotmail.labelname EmailfilterProfileMsnHotmail
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

