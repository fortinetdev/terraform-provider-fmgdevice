---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_attributematch"
description: |-
  Configure CASB SaaS application.
---

# fmgdevice_casb_attributematch
Configure CASB SaaS application.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `attribute`: `fmgdevice_casb_attributematch_attribute`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `application` - CASB tenant application name.
* `attribute` - Attribute. The structure of `attribute` block is documented below.
* `match_strategy` - CASB tenant match strategy. Valid values: `or`, `and`.

* `name` - CASB tenant match name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `attribute` block supports:

* `case_sensitive` - CASB attribute match case sensitive. Valid values: `disable`, `enable`.

* `match_pattern` - CASB attribute match pattern. Valid values: `simple`, `substr`, `regexp`.

* `match_value` - CASB attribute match value.
* `name` - CASB attribute match name.
* `negate` - Enable/disable what the matching strategy must not be. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Casb AttributeMatch can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_attributematch.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

