---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_profile"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure web proxy profiles.
---

# fmgdevice_webproxy_profile
<i>This object will be purged after policy copy and install.</i> Configure web proxy profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `headers`: `fmgdevice_webproxy_profile_headers`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `header_client_ip` - Action to take on the HTTP client-IP header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `header_front_end_https` - Action to take on the HTTP front-end-HTTPS header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `header_via_request` - Action to take on the HTTP via header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `header_via_response` - Action to take on the HTTP via header in forwarded responses: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `header_x_authenticated_groups` - Action to take on the HTTP x-authenticated-groups header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `header_x_authenticated_user` - Action to take on the HTTP x-authenticated-user header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `header_x_forwarded_client_cert` - Action to take on the HTTP x-forwarded-client-cert header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `header_x_forwarded_for` - Action to take on the HTTP x-forwarded-for header in forwarded requests: forwards (pass), adds, or removes the HTTP header. Valid values: `pass`, `add`, `remove`.

* `headers` - Headers. The structure of `headers` block is documented below.
* `log_header_change` - Enable/disable logging HTTP header changes. Valid values: `disable`, `enable`.

* `name` - Profile name.
* `strip_encoding` - Enable/disable stripping unsupported encoding from the request header. Valid values: `disable`, `enable`.

* `max_cache_object_size` - Max-Cache-Object-Size.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `headers` block supports:

* `action` - Configure adding, removing, or logging of the HTTP header entry in HTTP requests and responses. Valid values: `add-to-request`, `add-to-response`, `remove-from-request`, `remove-from-response`, `monitor-request`, `monitor-response`.

* `add_option` - Configure options to append content to existing HTTP header or add new HTTP header. Valid values: `append`, `new-on-not-found`, `new`, `replace`, `replace-when-match`.

* `base64_encoding` - Enable/disable use of base64 encoding of HTTP content. Valid values: `disable`, `enable`.

* `content` - HTTP header content (max length: 3999 characters).
* `dstaddr` - Destination address and address group names.
* `dstaddr6` - Destination address and address group names (IPv6).
* `id` - HTTP forwarded header id.
* `name` - HTTP forwarded header name.
* `protocol` - Configure protocol(s) to take add-option action on (HTTP, HTTPS, or both). Valid values: `https`, `http`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy Profile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

