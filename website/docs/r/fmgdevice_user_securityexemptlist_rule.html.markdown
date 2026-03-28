---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_user_securityexemptlist_rule"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure rules for exempting users from captive portal authentication.
---

# fmgdevice_user_securityexemptlist_rule
<i>This object will be purged after policy copy and install.</i> Configure rules for exempting users from captive portal authentication.

~> This resource is a sub resource for variable `rule` of resource `fmgdevice_user_securityexemptlist`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `security_exempt_list` - Security Exempt List.

* `dstaddr` - Destination addresses or address groups.
* `fosid` - ID.
* `service` - Destination services.
* `srcaddr` - Source addresses or address groups.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

User SecurityExemptListRule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "security_exempt_list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_user_securityexemptlist_rule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

