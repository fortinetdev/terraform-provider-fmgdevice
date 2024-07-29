---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wanopt_webcache"
description: |-
  Configure global Web cache settings.
---

# fmgdevice_wanopt_webcache
Configure global Web cache settings.

## Example Usage

```hcl
resource "fmgdevice_wanopt_webcache" "trname" {
  always_revalidate = "enable"
  cache_by_default  = "enable"
  cache_cookie      = "enable"
  cache_expired     = "disable"
  default_ttl       = 10
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `always_revalidate` - Enable/disable revalidation of requested cached objects, which have content on the server, before serving it to the client. Valid values: `disable`, `enable`.

* `cache_by_default` - Enable/disable caching content that lacks explicit caching policies from the server. Valid values: `disable`, `enable`.

* `cache_cookie` - Enable/disable caching cookies. Since cookies contain information for or about individual users, they not usually cached. Valid values: `disable`, `enable`.

* `cache_expired` - Enable/disable caching type-1 objects that are already expired on arrival. Valid values: `disable`, `enable`.

* `default_ttl` - Default object expiry time (default = 1440 min (1 day); maximum = 5256000 min (10 years)). This only applies to those objects that do not have an expiry time set by the web server.
* `external` - Enable/disable external Web caching. Valid values: `disable`, `enable`.

* `fresh_factor` - Frequency that the server is checked to see if any objects have expired (1 - 100, default = 100). The higher the fresh factor, the less often the checks occur.
* `host_validate` - Enable/disable validating "Host:" with original server IP. Valid values: `disable`, `enable`.

* `ignore_conditional` - Enable/disable controlling the behavior of cache-control HTTP 1.1 header values. Valid values: `disable`, `enable`.

* `ignore_ie_reload` - Enable/disable ignoring the PNC-interpretation of Internet Explorer's Accept: / header. Valid values: `disable`, `enable`.

* `ignore_ims` - Enable/disable ignoring the if-modified-since (IMS) header. Valid values: `disable`, `enable`.

* `ignore_pnc` - Enable/disable ignoring the pragma no-cache (PNC) header. Valid values: `disable`, `enable`.

* `max_object_size` - Maximum cacheable object size in kB (1 - 2147483 kb (2GB). All objects that exceed this are delivered to the client but not stored in the web cache.
* `max_ttl` - Maximum time an object can stay in the web cache without checking to see if it has expired on the server (default = 7200 min (5 days); maximum = 5256000 min (10 years)).
* `min_ttl` - Minimum time an object can stay in the web cache without checking to see if it has expired on the server (default = 5 min; maximum = 5256000 (10 years)).
* `neg_resp_time` - Time in minutes to cache negative responses or errors (0 - 4294967295, default = 0  which means negative responses are not cached).
* `reval_pnc` - Enable/disable revalidation of pragma-no-cache (PNC) to address bandwidth concerns. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Wanopt Webcache can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wanopt_webcache.labelname WanoptWebcache
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

