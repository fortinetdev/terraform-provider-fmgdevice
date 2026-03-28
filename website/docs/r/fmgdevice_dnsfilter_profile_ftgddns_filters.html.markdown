---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dnsfilter_profile_ftgddns_filters"
description: |-
  <i>This object will be purged after policy copy and install.</i> FortiGuard DNS domain filters.
---

# fmgdevice_dnsfilter_profile_ftgddns_filters
<i>This object will be purged after policy copy and install.</i> FortiGuard DNS domain filters.

~> This resource is a sub resource for variable `filters` of resource `fmgdevice_dnsfilter_profile_ftgddns`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `action` - Action to take for DNS requests matching the category. Valid values: `monitor`, `block`.

* `category` - Category number.
* `fosid` - ID number.
* `log` - Enable/disable DNS filter logging for this DNS profile. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Dnsfilter ProfileFtgdDnsFilters can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dnsfilter_profile_ftgddns_filters.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

