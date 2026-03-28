---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dnsfilter_domainfilter"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure DNS domain filters.
---

# fmgdevice_dnsfilter_domainfilter
<i>This object will be purged after policy copy and install.</i> Configure DNS domain filters.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fmgdevice_dnsfilter_domainfilter_entries`



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

* `action` - Action to take for domain filter matches. Valid values: `block`, `allow`, `monitor`.

* `comment` - Comment.
* `domain` - Domain entries to be filtered.
* `id` - Id.
* `status` - Enable/disable this domain filter. Valid values: `disable`, `enable`.

* `type` - DNS domain filter type. Valid values: `simple`, `regex`, `wildcard`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Dnsfilter DomainFilter can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dnsfilter_domainfilter.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

