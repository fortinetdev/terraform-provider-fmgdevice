---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_voip_profile_msrp"
description: |-
  <i>This object will be purged after policy copy and install.</i> MSRP.
---

# fmgdevice_voip_profile_msrp
<i>This object will be purged after policy copy and install.</i> MSRP.

~> This resource is a sub resource for variable `msrp` of resource `fmgdevice_voip_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `log_violations` - Enable/disable logging of MSRP violations. Valid values: `disable`, `enable`.

* `max_msg_size` - Maximum allowable MSRP message size (1-65535).
* `max_msg_size_action` - Action for violation of max-msg-size. Valid values: `pass`, `block`, `reset`, `monitor`.

* `status` - Enable/disable MSRP. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Voip ProfileMsrp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_voip_profile_msrp.labelname VoipProfileMsrp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

