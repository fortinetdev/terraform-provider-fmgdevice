---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_emailfilter_blockallowlist"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure anti-spam block/allow list.
---

# fmgdevice_emailfilter_blockallowlist
<i>This object will be purged after policy copy and install.</i> Configure anti-spam block/allow list.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fmgdevice_emailfilter_blockallowlist_entries`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Optional comments.
* `entries` - Entries. The structure of `entries` block is documented below.
* `fosid` - ID.
* `name` - Name of table.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `action` - Reject, mark as spam or good email. Valid values: `spam`, `clear`, `reject`.

* `addr_type` - IP address type. Valid values: `ipv4`, `ipv6`.

* `id` - Entry ID.
* `ip4_subnet` - IPv4 network address/subnet mask bits.
* `ip6_subnet` - IPv6 network address/subnet mask bits.
* `pattern` - Pattern to match.
* `pattern_type` - Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.

* `status` - Enable/disable status. Valid values: `disable`, `enable`.

* `type` - Entry type. Valid values: `ip`, `email-to`, `email-from`, `subject`.

* `email_pattern` - Email-Pattern.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Emailfilter BlockAllowList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_emailfilter_blockallowlist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

