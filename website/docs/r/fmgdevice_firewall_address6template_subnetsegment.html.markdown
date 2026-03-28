---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_address6template_subnetsegment"
description: |-
  <i>This object will be purged after policy copy and install.</i> IPv6 subnet segments.
---

# fmgdevice_firewall_address6template_subnetsegment
<i>This object will be purged after policy copy and install.</i> IPv6 subnet segments.

~> This resource is a sub resource for variable `subnet_segment` of resource `fmgdevice_firewall_address6template`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `values`: `fmgdevice_firewall_address6template_subnetsegment_values`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `address6_template` - Address6 Template.

* `bits` - Number of bits.
* `exclusive` - Enable/disable exclusive value. Valid values: `disable`, `enable`.

* `fosid` - Subnet segment ID.
* `name` - Subnet segment name.
* `values` - Values. The structure of `values` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `values` block supports:

* `name` - Subnet segment value name.
* `value` - Subnet segment value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall Address6TemplateSubnetSegment can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "address6_template=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_address6template_subnetsegment.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

