---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_profile_web"
description: |-
  <i>This object will be purged after policy copy and install.</i> Web content filtering settings.
---

# fmgdevice_webfilter_profile_web
<i>This object will be purged after policy copy and install.</i> Web content filtering settings.

~> This resource is a sub resource for variable `web` of resource `fmgdevice_webfilter_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `blacklist` - Enable/disable automatic addition of URLs detected by FortiSandbox to blacklist. Valid values: `disable`, `enable`.

* `allowlist` - FortiGuard allowlist settings. Valid values: `exempt-av`, `exempt-webcontent`, `exempt-activex-java-cookie`, `exempt-dlp`, `exempt-rangeblock`, `extended-log-others`.

* `blocklist` - Enable/disable automatic addition of URLs detected by FortiSandbox to blocklist. Valid values: `disable`, `enable`.

* `bword_table` - Banned word table ID.
* `bword_threshold` - Banned word score threshold.
* `content_header_list` - Content header list.
* `keyword_match` - Search keywords to log when match is found.
* `log_search` - Enable/disable logging all search phrases. Valid values: `disable`, `enable`.

* `safe_search` - Safe search type. Valid values: `url`, `header`.

* `urlfilter_table` - URL filter table ID.
* `whitelist` - FortiGuard whitelist settings. Valid values: `exempt-av`, `exempt-webcontent`, `exempt-activex-java-cookie`, `exempt-dlp`, `exempt-rangeblock`, `extended-log-others`.

* `vimeo_restrict` - Set Vimeo-restrict ("7" = don't show mature content, "134" = don't show unrated and mature content). A value of cookie "content_rating".
* `youtube_restrict` - YouTube EDU filter level. Valid values: `strict`, `none`, `moderate`.

* `qwant_restrict` - Qwant-Restrict. Valid values: `strict`, `none`, `moderate`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Webfilter ProfileWeb can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_profile_web.labelname WebfilterProfileWeb
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

