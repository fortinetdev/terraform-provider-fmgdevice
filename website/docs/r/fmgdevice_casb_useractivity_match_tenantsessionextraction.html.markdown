---
subcategory: "CASB"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_useractivity_match_tenantsessionextraction"
description: |-
  <i>This object will be purged after policy copy and install.</i> CASB user activity tenant session extraction.
---

# fmgdevice_casb_useractivity_match_tenantsessionextraction
<i>This object will be purged after policy copy and install.</i> CASB user activity tenant session extraction.

~> This resource is a sub resource for variable `tenant_session_extraction` of resource `fmgdevice_casb_useractivity_match`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `filters`: `fmgdevice_casb_useractivity_match_tenantsessionextraction_filters`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `user_activity` - User Activity.
* `match` - Match.

* `filters` - Filters. The structure of `filters` block is documented below.
* `jq` - CASB user activity session extraction jq script.
* `session_match` - CASB user activity session match name.
* `session_source` - Enable/disable CASB session extraction source flag. Valid values: `disable`, `enable`.

* `status` - Enable/disable CASB session extraction. Valid values: `disable`, `enable`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `filters` block supports:

* `body_type` - CASB content extraction filter body type. Valid values: `json`, `form`.

* `cookie_name` - CASB content extraction filter cookie name.
* `direction` - CASB content extraction filter direction. Valid values: `request`, `response`.

* `header_name` - CASB content extraction filter header name.
* `id` - CASB content extraction filter ID.
* `place` - CASB content extraction filter place type. Valid values: `header`, `path`, `body`, `cookie`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Casb UserActivityMatchTenantSessionExtraction can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "user_activity=YOUR_VALUE", "match=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_useractivity_match_tenantsessionextraction.labelname CasbUserActivityMatchTenantSessionExtraction
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

