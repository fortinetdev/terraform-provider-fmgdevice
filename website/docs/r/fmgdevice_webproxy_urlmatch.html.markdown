---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webproxy_urlmatch"
description: |-
  Exempt URLs from web proxy forwarding and caching.
---

# fmgdevice_webproxy_urlmatch
Exempt URLs from web proxy forwarding and caching.

## Example Usage

```hcl
resource "fmgdevice_webproxy_urlmatch" "trname" {
  cache_exemption = "enable"
  comment         = "your own value"
  fast_fallback   = ["your own value"]
  forward_server  = ["your own value"]
  name            = "your own value"
  device_name     = var.device_name # not required if setting is at provider
  device_vdom     = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `cache_exemption` - Enable/disable exempting this URL pattern from caching. Valid values: `disable`, `enable`.

* `comment` - Comment.
* `fast_fallback` - Fast fallback configuration entry name.
* `forward_server` - Forward server name.
* `name` - Configure a name for the URL to be exempted.
* `status` - Enable/disable exempting the URLs matching the URL pattern from web proxy forwarding and caching. Valid values: `disable`, `enable`.

* `url_pattern` - URL pattern to be exempted from web proxy forwarding and caching.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy UrlMatch can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webproxy_urlmatch.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

