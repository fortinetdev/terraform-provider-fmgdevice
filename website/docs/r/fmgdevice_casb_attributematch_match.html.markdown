---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_attributematch_match"
description: |-
  CASB tenant match rules.
---

# fmgdevice_casb_attributematch_match
CASB tenant match rules.

~> This resource is a sub resource for variable `match` of resource `fmgdevice_casb_attributematch`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `rule`: `fmgdevice_casb_attributematch_match_rule`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `attribute_match` - Attribute Match.

* `fosid` - CASB attribute match rule ID.
* `rule` - Rule. The structure of `rule` block is documented below.
* `rule_strategy` - CASB attribute match rule strategy. Valid values: `or`, `and`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rule` block supports:

* `attribute` - CASB attribute match name.
* `case_sensitive` - CASB attribute match case sensitive. Valid values: `disable`, `enable`.

* `id` - CASB attribute rule ID.
* `match_pattern` - CASB attribute match pattern. Valid values: `simple`, `substr`, `regexp`.

* `match_value` - CASB attribute match value.
* `negate` - Enable/disable what the matching strategy must not be. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Casb AttributeMatchMatch can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "attribute_match=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_attributematch_match.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

