---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_icap_profile_icapheaders"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure ICAP forwarded request headers.
---

# fmgdevice_icap_profile_icapheaders
<i>This object will be purged after policy copy and install.</i> Configure ICAP forwarded request headers.

~> This resource is a sub resource for variable `icap_headers` of resource `fmgdevice_icap_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `base64_encoding` - Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.

* `content` - HTTP header content.
* `fosid` - HTTP forwarded header ID.
* `name` - HTTP forwarded header name.
* `http_header` - Http-Header.
* `sesson_info_type` - Sesson-Info-Type. Valid values: `client-ip`, `user`, `upn`, `domain`, `local-grp`, `remote-grp`, `proxy-name`, `auth-user-uri`, `auth-group-uri`.

* `source` - Source. Valid values: `content`, `http-header`, `session`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Icap ProfileIcapHeaders can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_icap_profile_icapheaders.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

