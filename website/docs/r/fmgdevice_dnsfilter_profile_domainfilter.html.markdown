---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dnsfilter_profile_domainfilter"
description: |-
  <i>This object will be purged after policy copy and install.</i> Domain filter settings.
---

# fmgdevice_dnsfilter_profile_domainfilter
<i>This object will be purged after policy copy and install.</i> Domain filter settings.

~> This resource is a sub resource for variable `domain_filter` of resource `fmgdevice_dnsfilter_profile`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `profile` - Profile.

* `domain_filter_table` - DNS domain filter table ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dnsfilter ProfileDomainFilter can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dnsfilter_profile_domainfilter.labelname DnsfilterProfileDomainFilter
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

