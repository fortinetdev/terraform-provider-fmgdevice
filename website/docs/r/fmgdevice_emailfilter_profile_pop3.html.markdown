---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_emailfilter_profile_pop3"
description: |-
  <i>This object will be purged after policy copy and install.</i> POP3.
---

# fmgdevice_emailfilter_profile_pop3
<i>This object will be purged after policy copy and install.</i> POP3.

~> This resource is a sub resource for variable `pop3` of resource `fmgdevice_emailfilter_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `action` - Action for spam email. Valid values: `pass`, `tag`.

* `log` - Log. Valid values: `disable`, `enable`.

* `log_all` - Enable/disable logging of all email traffic. Valid values: `disable`, `enable`.

* `tag_msg` - Subject text or header added to spam email.
* `tag_type` - Tag subject or header for spam email. Valid values: `subject`, `header`, `spaminfo`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Emailfilter ProfilePop3 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_emailfilter_profile_pop3.labelname EmailfilterProfilePop3
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

