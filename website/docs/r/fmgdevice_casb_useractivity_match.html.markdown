---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_useractivity_match"
description: |-
  <i>This object will be purged after policy copy and install.</i> CASB user activity match rules.
---

# fmgdevice_casb_useractivity_match
<i>This object will be purged after policy copy and install.</i> CASB user activity match rules.

~> This resource is a sub resource for variable `match` of resource `fmgdevice_casb_useractivity`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `rules`: `fmgdevice_casb_useractivity_match_rules`
>- `tenant_extraction`: `fmgdevice_casb_useractivity_match_tenantextraction`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `user_activity` - User Activity.

* `fosid` - CASB user activity match rules ID.
* `rules` - Rules. The structure of `rules` block is documented below.
* `strategy` - CASB user activity rules strategy. Valid values: `or`, `and`.

* `tenant_extraction` - Tenant-Extraction. The structure of `tenant_extraction` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rules` block supports:

* `body_type` - CASB user activity match rule body type. Valid values: `json`.

* `case_sensitive` - CASB user activity match case sensitive. Valid values: `disable`, `enable`.

* `domains` - CASB user activity domain list.
* `header_name` - CASB user activity rule header name.
* `id` - CASB user activity rule ID.
* `jq` - CASB user activity rule match jq script.
* `match_pattern` - CASB user activity rule match pattern. Valid values: `simple`, `substr`, `regexp`.

* `match_value` - CASB user activity rule match value.
* `methods` - CASB user activity method list.
* `negate` - Enable/disable what the matching strategy must not be. Valid values: `disable`, `enable`.

* `type` - CASB user activity rule type. Valid values: `domains`, `host`, `path`, `header`, `header-value`, `method`, `body`.


The `tenant_extraction` block supports:

* `filters` - Filters. The structure of `filters` block is documented below.
* `jq` - CASB user activity tenant extraction jq script.
* `status` - Enable/disable CASB tenant extraction. Valid values: `disable`, `enable`.

* `type` - CASB user activity tenant extraction type. Valid values: `json-query`.


The `filters` block supports:

* `body_type` - CASB tenant extraction filter body type. Valid values: `json`.

* `direction` - CASB tenant extraction filter direction. Valid values: `request`, `response`.

* `header_name` - CASB tenant extraction filter header name.
* `id` - CASB tenant extraction filter ID.
* `place` - CASB tenant extraction filter place type. Valid values: `path`, `header`, `body`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Casb UserActivityMatch can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "user_activity=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_useractivity_match.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

