---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_urlfilter_entries"
description: |-
  <i>This object will be purged after policy copy and install.</i> URL filter entries.
---

# fmgdevice_webfilter_urlfilter_entries
<i>This object will be purged after policy copy and install.</i> URL filter entries.

~> This resource is a sub resource for variable `entries` of resource `fmgdevice_webfilter_urlfilter`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `urlfilter` - Urlfilter.

* `action` - Action to take for URL filter matches. Valid values: `exempt`, `block`, `allow`, `monitor`.

* `antiphish_action` - Action to take for AntiPhishing matches. Valid values: `block`, `log`.

* `comment` - Comment.
* `dns_address_family` - Resolve IPv4 address, IPv6 address, or both from DNS server. Valid values: `ipv4`, `ipv6`, `both`.

* `exempt` - If action is set to exempt, select the security profile operations that exempt URLs skip. Separate multiple options with a space. Valid values: `av`, `web-content`, `activex-java-cookie`, `dlp`, `fortiguard`, `all`, `pass`, `range-block`, `antiphish`.

* `fosid` - Id.
* `referrer_host` - Referrer host name.
* `status` - Enable/disable this URL filter. Valid values: `disable`, `enable`.

* `type` - Filter type (simple, regex, or wildcard). Valid values: `simple`, `regex`, `wildcard`.

* `url` - URL to be filtered.
* `web_proxy_profile` - Web proxy profile.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Webfilter UrlfilterEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "urlfilter=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_urlfilter_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

