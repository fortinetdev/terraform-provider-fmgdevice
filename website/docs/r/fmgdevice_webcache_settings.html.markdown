---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webcache_settings"
description: |-
  Device WebcacheSettings
---

# fmgdevice_webcache_settings
Device WebcacheSettings

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `add_x_cache` - Add-X-Cache. Valid values: `disable`, `enable`.

* `always_revalidate` - Always-Revalidate. Valid values: `disable`, `enable`.

* `cache_by_default` - Cache-By-Default. Valid values: `disable`, `enable`.

* `cache_cookie` - Cache-Cookie. Valid values: `disable`, `enable`.

* `cache_expired` - Cache-Expired. Valid values: `disable`, `enable`.

* `crawler_user_agent` - Crawler-User-Agent.
* `default_ttl` - Default-Ttl.
* `external` - External. Valid values: `disable`, `enable`.

* `fresh_factor` - Fresh-Factor.
* `host_validate` - Host-Validate. Valid values: `disable`, `enable`.

* `ignore_conditional` - Ignore-Conditional. Valid values: `disable`, `enable`.

* `ignore_ie_reload` - Ignore-Ie-Reload. Valid values: `disable`, `enable`.

* `ignore_ims` - Ignore-Ims. Valid values: `disable`, `enable`.

* `ignore_pnc` - Ignore-Pnc. Valid values: `disable`, `enable`.

* `max_object_size` - Max-Object-Size.
* `max_ttl` - Max-Ttl.
* `min_ttl` - Min-Ttl.
* `neg_resp_time` - Neg-Resp-Time.
* `reval_pnc` - Reval-Pnc. Valid values: `disable`, `enable`.

* `x_cache_message` - X-Cache-Message.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Webcache Settings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webcache_settings.labelname WebcacheSettings
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

