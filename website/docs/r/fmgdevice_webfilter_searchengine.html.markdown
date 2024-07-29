---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_searchengine"
description: |-
  Configure web filter search engines.
---

# fmgdevice_webfilter_searchengine
Configure web filter search engines.

## Example Usage

```hcl
resource "fmgdevice_webfilter_searchengine" "trname" {
  charset     = "utf-8"
  hostname    = "your own value"
  name        = "your own value"
  query       = "your own value"
  safesearch  = "header"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `charset` - Search engine charset. Valid values: `utf-8`, `gb2312`.

* `hostname` - Hostname (regular expression).
* `name` - Search engine name.
* `query` - Code used to prefix a query (must end with an equals character).
* `safesearch` - Safe search method. You can disable safe search, add the safe search string to URLs, or insert a safe search header. Valid values: `disable`, `url`, `header`, `translate`, `yt-pattern`, `yt-scan`, `yt-video`, `yt-channel`.

* `safesearch_str` - Safe search parameter used in the URL in URL mode. In translate mode, it provides either the regex to translate the URL or the special case to translate the URL.
* `url` - URL (regular expression).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Webfilter SearchEngine can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_searchengine.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

