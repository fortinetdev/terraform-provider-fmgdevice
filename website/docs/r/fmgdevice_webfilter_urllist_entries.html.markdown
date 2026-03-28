---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_urllist_entries"
description: |-
  <i>This object will be purged after policy copy and install.</i>
---

# fmgdevice_webfilter_urllist_entries
<i>This object will be purged after policy copy and install.</i>

~> This resource is a sub resource for variable `entries` of resource `fmgdevice_webfilter_urllist`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `url_list` - Url List.

* `url` - Url.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{url}}.

## Import

Webfilter UrlListEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "url_list=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_urllist_entries.labelname {{url}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

