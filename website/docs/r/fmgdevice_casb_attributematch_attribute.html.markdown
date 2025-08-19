---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_attributematch_attribute"
description: |-
  CASB tenant match rules.
---

# fmgdevice_casb_attributematch_attribute
CASB tenant match rules.

~> This resource is a sub resource for variable `attribute` of resource `fmgdevice_casb_attributematch`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `attribute_match` - Attribute Match.

* `case_sensitive` - CASB attribute match case sensitive. Valid values: `disable`, `enable`.

* `match_pattern` - CASB attribute match pattern. Valid values: `simple`, `substr`, `regexp`.

* `match_value` - CASB attribute match value.
* `name` - CASB attribute match name.
* `negate` - Enable/disable what the matching strategy must not be. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Casb AttributeMatchAttribute can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "attribute_match=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_attributematch_attribute.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

