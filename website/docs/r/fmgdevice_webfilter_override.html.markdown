---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_override"
description: |-
  Configure FortiGuard Web Filter administrative overrides.
---

# fmgdevice_webfilter_override
Configure FortiGuard Web Filter administrative overrides.

## Example Usage

```hcl
resource "fmgdevice_webfilter_override" "trname" {
  # fosid       = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `expires` - Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
* `fosid` - Override rule ID.
* `initiator` - Initiating user of override (read-only setting).
* `ip` - IPv4 address which the override applies.
* `ip6` - IPv6 address which the override applies.
* `new_profile` - Name of the new web filter profile used by the override.
* `old_profile` - Name of the web filter profile which the override applies.
* `scope` - Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.

* `status` - Enable/disable override rule. Valid values: `disable`, `enable`.

* `user` - Name of the user which the override applies.
* `user_group` - Specify the user group for which the override applies.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Webfilter Override can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_override.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

