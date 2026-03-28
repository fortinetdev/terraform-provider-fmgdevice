---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_casb_useractivity_match_rules"
description: |-
  <i>This object will be purged after policy copy and install.</i> CASB user activity rules.
---

# fmgdevice_casb_useractivity_match_rules
<i>This object will be purged after policy copy and install.</i> CASB user activity rules.

~> This resource is a sub resource for variable `rules` of resource `fmgdevice_casb_useractivity_match`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `user_activity` - User Activity.
* `match` - Match.

* `body_type` - CASB user activity match rule body type. Valid values: `json`.

* `case_sensitive` - CASB user activity match case sensitive. Valid values: `disable`, `enable`.

* `domains` - CASB user activity domain list.
* `header_name` - CASB user activity rule header name.
* `fosid` - CASB user activity rule ID.
* `jq` - CASB user activity rule match jq script.
* `match_pattern` - CASB user activity rule match pattern. Valid values: `simple`, `substr`, `regexp`.

* `match_value` - CASB user activity rule match value.
* `methods` - CASB user activity method list.
* `negate` - Enable/disable what the matching strategy must not be. Valid values: `disable`, `enable`.

* `type` - CASB user activity rule type. Valid values: `domains`, `host`, `path`, `header`, `header-value`, `method`, `body`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Casb UserActivityMatchRules can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "user_activity=YOUR_VALUE", "match=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_casb_useractivity_match_rules.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

