---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dnsfilter_domainfilter_entries"
description: |-
  <i>This object will be purged after policy copy and install.</i> DNS domain filter entries.
---

# fmgdevice_dnsfilter_domainfilter_entries
<i>This object will be purged after policy copy and install.</i> DNS domain filter entries.

~> This resource is a sub resource for variable `entries` of resource `fmgdevice_dnsfilter_domainfilter`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `domain_filter` - Domain Filter.

* `action` - Action to take for domain filter matches. Valid values: `block`, `allow`, `monitor`.

* `comment` - Comment.
* `domain` - Domain entries to be filtered.
* `fosid` - Id.
* `status` - Enable/disable this domain filter. Valid values: `disable`, `enable`.

* `type` - DNS domain filter type. Valid values: `simple`, `regex`, `wildcard`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Dnsfilter DomainFilterEntries can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "domain_filter=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dnsfilter_domainfilter_entries.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

