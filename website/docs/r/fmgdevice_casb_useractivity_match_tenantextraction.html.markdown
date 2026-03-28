---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_useractivity_match_tenantextraction"
description: |-
  <i>This object will be purged after policy copy and install.</i> CASB user activity tenant extraction.
---

# fmgdevice_casb_useractivity_match_tenantextraction
<i>This object will be purged after policy copy and install.</i> CASB user activity tenant extraction.

~> This resource is a sub resource for variable `tenant_extraction` of resource `fmgdevice_casb_useractivity_match`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `filters`: `fmgdevice_casb_useractivity_match_tenantextraction_filters`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `user_activity` - User Activity.
* `match` - Match.

* `filters` - Filters. The structure of `filters` block is documented below.
* `jq` - CASB user activity tenant extraction jq script.
* `status` - Enable/disable CASB tenant extraction. Valid values: `disable`, `enable`.

* `type` - CASB user activity tenant extraction type. Valid values: `json-query`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `filters` block supports:

* `body_type` - CASB tenant extraction filter body type. Valid values: `json`.

* `direction` - CASB tenant extraction filter direction. Valid values: `request`, `response`.

* `header_name` - CASB tenant extraction filter header name.
* `id` - CASB tenant extraction filter ID.
* `place` - CASB tenant extraction filter place type. Valid values: `path`, `header`, `body`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Casb UserActivityMatchTenantExtraction can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "user_activity=YOUR_VALUE", "match=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_useractivity_match_tenantextraction.labelname CasbUserActivityMatchTenantExtraction
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

