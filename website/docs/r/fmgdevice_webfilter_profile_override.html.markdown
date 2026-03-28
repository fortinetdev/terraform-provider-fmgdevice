---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_profile_override"
description: |-
  <i>This object will be purged after policy copy and install.</i> Web Filter override settings.
---

# fmgdevice_webfilter_profile_override
<i>This object will be purged after policy copy and install.</i> Web Filter override settings.

~> This resource is a sub resource for variable `override` of resource `fmgdevice_webfilter_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `ovrd_cookie` - Allow/deny browser-based (cookie) overrides. Valid values: `deny`, `allow`.

* `ovrd_dur` - Override duration.
* `ovrd_dur_mode` - Override duration mode. Valid values: `constant`, `ask`.

* `ovrd_scope` - Override scope. Valid values: `user`, `user-group`, `ip`, `ask`, `browser`.

* `ovrd_user_group` - User groups with permission to use the override.
* `profile` - Web filter profile with permission to create overrides.
* `profile_attribute` - Profile attribute to retrieve from the RADIUS server. Valid values: `User-Name`, `NAS-IP-Address`, `Framed-IP-Address`, `Framed-IP-Netmask`, `Filter-Id`, `Login-IP-Host`, `Reply-Message`, `Callback-Number`, `Callback-Id`, `Framed-Route`, `Framed-IPX-Network`, `Class`, `Called-Station-Id`, `Calling-Station-Id`, `NAS-Identifier`, `Proxy-State`, `Login-LAT-Service`, `Login-LAT-Node`, `Login-LAT-Group`, `Framed-AppleTalk-Zone`, `Acct-Session-Id`, `Acct-Multi-Session-Id`.

* `profile_type` - Override profile type. Valid values: `list`, `radius`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Webfilter ProfileOverride can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_profile_override.labelname WebfilterProfileOverride
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

