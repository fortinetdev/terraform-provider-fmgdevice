---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_useractivity_match_tenantextraction_filters"
description: |-
  <i>This object will be purged after policy copy and install.</i> CASB user activity tenant extraction filters.
---

# fmgdevice_casb_useractivity_match_tenantextraction_filters
<i>This object will be purged after policy copy and install.</i> CASB user activity tenant extraction filters.

~> This resource is a sub resource for variable `filters` of resource `fmgdevice_casb_useractivity_match_tenantextraction`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `user_activity` - User Activity.
* `match` - Match.

* `body_type` - CASB tenant extraction filter body type. Valid values: `json`.

* `direction` - CASB tenant extraction filter direction. Valid values: `request`, `response`.

* `header_name` - CASB tenant extraction filter header name.
* `fosid` - CASB tenant extraction filter ID.
* `place` - CASB tenant extraction filter place type. Valid values: `path`, `header`, `body`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Casb UserActivityMatchTenantExtractionFilters can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "user_activity=YOUR_VALUE", "match=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_useractivity_match_tenantextraction_filters.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

