---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_proxyaddress6"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_firewall_proxyaddress6
<i>This object will be purged after policy copy and install.</i>

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `header_group`: `fmgdevice_firewall_proxyaddress6_headergroup`
>- `tagging`: `fmgdevice_firewall_proxyaddress6_tagging`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `application` - Application.
* `case_sensitivity` - Case-Sensitivity. Valid values: `disable`, `enable`.

* `category` - Category.
* `color` - Color.
* `comment` - Comment.
* `header` - Header.
* `header_group` - Header-Group. The structure of `header_group` block is documented below.
* `header_name` - Header-Name.
* `host` - Host.
* `host_regex` - Host-Regex.
* `method` - Method. Valid values: `delete`, `get`, `head`, `options`, `post`, `put`, `trace`, `connect`.

* `name` - Name.
* `path` - Path.
* `post_arg` - Post-Arg. Valid values: `disable`, `enable`.

* `query` - Query.
* `referrer` - Referrer. Valid values: `disable`, `enable`.

* `tagging` - Tagging. The structure of `tagging` block is documented below.
* `type` - Type. Valid values: `host-regex`, `url`, `category`, `method`, `ua`, `header`, `src-advanced`, `dst-advanced`, `url-list`, `saas`, `response-header`.

* `ua` - Ua. Valid values: `chrome`, `ms`, `firefox`, `safari`, `other`, `ie`, `edge`.

* `ua_max_ver` - Ua-Max-Ver.
* `ua_min_ver` - Ua-Min-Ver.
* `url_list` - Url-List.
* `uuid` - Uuid.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `header_group` block supports:

* `case_sensitivity` - Case-Sensitivity. Valid values: `disable`, `enable`.

* `header` - Header.
* `header_name` - Header-Name.
* `id` - Id.

The `tagging` block supports:

* `category` - Category.
* `name` - Name.
* `tags` - Tags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ProxyAddress6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_proxyaddress6.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

