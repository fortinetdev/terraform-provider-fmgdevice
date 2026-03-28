---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_attributematch_match_rule"
description: |-
  CASB attribute match rule.
---

# fmgdevice_casb_attributematch_match_rule
CASB attribute match rule.

~> This resource is a sub resource for variable `rule` of resource `fmgdevice_casb_attributematch_match`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `attribute_match` - Attribute Match.
* `match` - Match.

* `attribute` - CASB attribute match name.
* `case_sensitive` - CASB attribute match case sensitive. Valid values: `disable`, `enable`.

* `fosid` - CASB attribute rule ID.
* `match_pattern` - CASB attribute match pattern. Valid values: `simple`, `substr`, `regexp`.

* `match_value` - CASB attribute match value.
* `negate` - Enable/disable what the matching strategy must not be. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Casb AttributeMatchMatchRule can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "attribute_match=YOUR_VALUE", "match=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_attributematch_match_rule.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

