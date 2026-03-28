---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_webfilter_urlfilter"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure URL filter lists.
---

# fmgdevice_webfilter_urlfilter
<i>This object will be purged after policy copy and install.</i> Configure URL filter lists.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fmgdevice_webfilter_urlfilter_entries`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Optional comments.
* `entries` - Entries. The structure of `entries` block is documented below.
* `fosid` - ID.
* `include_subdomains` - Enable/disable matching subdomains. Applies only to simple type (default = enable). Valid values: `disable`, `enable`.

* `ip_addr_block` - Enable/disable blocking URLs when the hostname appears as an IP address. Valid values: `disable`, `enable`.

* `ip4_mapped_ip6` - Enable/disable matching of IPv4 mapped IPv6 URLs. Valid values: `disable`, `enable`.

* `name` - Name of URL filter list.
* `one_arm_ips_urlfilter` - Enable/disable DNS resolver for one-arm IPS URL filter operation. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `action` - Action to take for URL filter matches. Valid values: `exempt`, `block`, `allow`, `monitor`.

* `antiphish_action` - Action to take for AntiPhishing matches. Valid values: `block`, `log`.

* `comment` - Comment.
* `dns_address_family` - Resolve IPv4 address, IPv6 address, or both from DNS server. Valid values: `ipv4`, `ipv6`, `both`.

* `exempt` - If action is set to exempt, select the security profile operations that exempt URLs skip. Separate multiple options with a space. Valid values: `av`, `web-content`, `activex-java-cookie`, `dlp`, `fortiguard`, `all`, `pass`, `range-block`, `antiphish`.

* `id` - Id.
* `referrer_host` - Referrer host name.
* `status` - Enable/disable this URL filter. Valid values: `disable`, `enable`.

* `type` - Filter type (simple, regex, or wildcard). Valid values: `simple`, `regex`, `wildcard`.

* `url` - URL to be filtered.
* `web_proxy_profile` - Web proxy profile.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Webfilter Urlfilter can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_webfilter_urlfilter.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

