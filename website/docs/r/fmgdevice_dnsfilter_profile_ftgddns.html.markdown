---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dnsfilter_profile_ftgddns"
description: |-
  <i>This object will be purged after policy copy and install.</i> FortiGuard DNS Filter settings.
---

# fmgdevice_dnsfilter_profile_ftgddns
<i>This object will be purged after policy copy and install.</i> FortiGuard DNS Filter settings.

~> This resource is a sub resource for variable `ftgd_dns` of resource `fmgdevice_dnsfilter_profile`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `filters`: `fmgdevice_dnsfilter_profile_ftgddns_filters`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `filters` - Filters. The structure of `filters` block is documented below.
* `options` - FortiGuard DNS filter options. Valid values: `error-allow`, `ftgd-disable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `filters` block supports:

* `action` - Action to take for DNS requests matching the category. Valid values: `monitor`, `block`.

* `category` - Category number.
* `id` - ID number.
* `log` - Enable/disable DNS filter logging for this DNS profile. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dnsfilter ProfileFtgdDns can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dnsfilter_profile_ftgddns.labelname DnsfilterProfileFtgdDns
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

