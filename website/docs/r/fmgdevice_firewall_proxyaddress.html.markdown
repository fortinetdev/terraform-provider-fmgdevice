---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_proxyaddress"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure web proxy address.
---

# fmgdevice_firewall_proxyaddress
<i>This object will be purged after policy copy and install.</i> Configure web proxy address.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `header_group`: `fmgdevice_firewall_proxyaddress_headergroup`
>- `tagging`: `fmgdevice_firewall_proxyaddress_tagging`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `_image_base64` - _Image-Base64.
* `application` - SaaS application.
* `case_sensitivity` - Enable to make the pattern case sensitive. Valid values: `disable`, `enable`.

* `category` - FortiGuard category ID.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `comment` - Optional comments.
* `header` - HTTP header name as a regular expression.
* `header_group` - Header-Group. The structure of `header_group` block is documented below.
* `header_name` - Name of HTTP header.
* `host` - Address object for the host.
* `host_regex` - Host name as a regular expression.
* `method` - HTTP request methods to be used. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `other`, `connect`, `patch`, `update`.

* `name` - Address name.
* `path` - URL path as a regular expression.
* `query` - Match the query part of the URL as a regular expression.
* `referrer` - Enable/disable use of referrer field in the HTTP header to match the address. Valid values: `disable`, `enable`.

* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `type` - Proxy address type. Valid values: `host-regex`, `url`, `category`, `method`, `ua`, `header`, `src-advanced`, `dst-advanced`, `saas`.

* `ua` - Names of browsers to be used as user agent. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`, `ie`, `edge`.

* `ua_max_ver` - Maximum version of the user agent specified in dotted notation. For example, use 120 with the ua field set to "chrome" to require Google Chrome's maximum version must be 120.
* `ua_min_ver` - Minimum version of the user agent specified in dotted notation. For example, use 90.0.1 with the ua field set to "chrome" to require Google Chrome's minimum version must be 90.0.1.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `post_arg` - Post-Arg. Valid values: `disable`, `enable`.

* `url_list` - Url-List.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `header_group` block supports:

* `case_sensitivity` - Case sensitivity in pattern. Valid values: `disable`, `enable`.

* `header` - HTTP header regular expression.
* `header_name` - HTTP header.
* `id` - ID.

The `tagging` block supports:

* `category` - Tag category.
* `name` - Tagging entry name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ProxyAddress can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_proxyaddress.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

