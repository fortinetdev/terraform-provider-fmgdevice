---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_voip_profile_sccp"
description: |-
  <i>This object will be purged after policy copy and install.</i> SCCP.
---

# fmgdevice_voip_profile_sccp
<i>This object will be purged after policy copy and install.</i> SCCP.

~> This resource is a sub resource for variable `sccp` of resource `fmgdevice_voip_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `block_mcast` - Enable/disable block multicast RTP connections. Valid values: `disable`, `enable`.

* `log_call_summary` - Enable/disable log summary of SCCP calls. Valid values: `disable`, `enable`.

* `log_violations` - Enable/disable logging of SCCP violations. Valid values: `disable`, `enable`.

* `max_calls` - Maximum calls per minute per SCCP client (max 65535).
* `status` - Enable/disable SCCP. Valid values: `disable`, `enable`.

* `verify_header` - Enable/disable verify SCCP header content. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Voip ProfileSccp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_voip_profile_sccp.labelname VoipProfileSccp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

