---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ngfwsettings"
description: |-
  Configure IPS NGFW policy-mode VDOM settings.
---

# fmgdevice_system_ngfwsettings
Configure IPS NGFW policy-mode VDOM settings.

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `match_timeout` - Number of seconds to wait before a security policy match for an idle non-TCP session (0 - 1800, default = 300, 0 means unlimited).
* `tcp_halfopen_match_timeout` - Number of seconds to wait before a security policy match for a session after one peer has sent an open session packet but the other has not responded (0 - 300, default = 8, 0 means unlimited).
* `tcp_match_timeout` - Number of seconds to wait before a security policy match for an idle TCP session (0 - 1800, default = 300, 0 means unlimited).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System NgfwSettings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ngfwsettings.labelname SystemNgfwSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

