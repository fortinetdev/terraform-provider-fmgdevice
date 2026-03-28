---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_urllist"
description: |-
  Device WebProxyUrlList
---

# fmgdevice_webproxy_urllist
Device WebProxyUrlList

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fmgdevice_webproxy_urllist_entries`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `entries` - Entries. The structure of `entries` block is documented below.
* `name` - Name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `id` - Id.
* `status` - Status. Valid values: `disable`, `enable`.

* `type` - Type. Valid values: `simple`, `wildcard`.

* `url` - Url.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy UrlList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_urllist.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

