---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_waf_profile_addresslist"
description: |-
  <i>This object will be purged after policy copy and install.</i> Address block and allow lists.
---

# fmgdevice_waf_profile_addresslist
<i>This object will be purged after policy copy and install.</i> Address block and allow lists.

~> This resource is a sub resource for variable `address_list` of resource `fmgdevice_waf_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `blocked_address` - Blocked address.
* `blocked_log` - Enable/disable logging on blocked addresses. Valid values: `disable`, `enable`.

* `severity` - Severity. Valid values: `low`, `medium`, `high`.

* `status` - Status. Valid values: `disable`, `enable`.

* `trusted_address` - Trusted address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Waf ProfileAddressList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_waf_profile_addresslist.labelname WafProfileAddressList
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

