---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ssl_web_userbookmark_bookmarks_formdata"
description: |-
  Form data.
---

# fmgdevice_vpn_ssl_web_userbookmark_bookmarks_formdata
Form data.

~> This resource is a sub resource for variable `form_data` of resource `fmgdevice_vpn_ssl_web_userbookmark_bookmarks`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_vpn_ssl_web_userbookmark_bookmarks_formdata" "trname" {
  name        = "your own value"
  value       = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `user_bookmark` - User Bookmark.
* `bookmarks` - Bookmarks.

* `name` - Name.
* `value` - Value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn SslWebUserBookmarkBookmarksFormData can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "user_bookmark=YOUR_VALUE", "bookmarks=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ssl_web_userbookmark_bookmarks_formdata.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

