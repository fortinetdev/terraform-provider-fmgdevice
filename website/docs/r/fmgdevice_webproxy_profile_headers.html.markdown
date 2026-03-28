---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_profile_headers"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure HTTP forwarded requests headers.
---

# fmgdevice_webproxy_profile_headers
<i>This object will be purged after policy copy and install.</i> Configure HTTP forwarded requests headers.

~> This resource is a sub resource for variable `headers` of resource `fmgdevice_webproxy_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `action` - Configure adding, removing, or logging of the HTTP header entry in HTTP requests and responses. Valid values: `add-to-request`, `add-to-response`, `remove-from-request`, `remove-from-response`, `monitor-request`, `monitor-response`.

* `add_option` - Configure options to append content to existing HTTP header or add new HTTP header. Valid values: `append`, `new-on-not-found`, `new`, `replace`, `replace-when-match`.

* `base64_encoding` - Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.

* `content` - HTTP header content (max length: 3999 characters).
* `dstaddr` - Destination address and address group names.
* `dstaddr6` - Destination address and address group names (IPv6).
* `fosid` - HTTP forwarded header id.
* `name` - HTTP forwarded header name.
* `protocol` - Configure protocol(s) to take add-option action on (HTTP, HTTPS, or both). Valid values: `https`, `http`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WebProxy ProfileHeaders can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_profile_headers.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

