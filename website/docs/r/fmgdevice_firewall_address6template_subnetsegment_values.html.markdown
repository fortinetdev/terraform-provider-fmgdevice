---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_address6template_subnetsegment_values"
description: |-
  <i>This object will be purged after policy copy and install.</i> Subnet segment values.
---

# fmgdevice_firewall_address6template_subnetsegment_values
<i>This object will be purged after policy copy and install.</i> Subnet segment values.

~> This resource is a sub resource for variable `values` of resource `fmgdevice_firewall_address6template_subnetsegment`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `address6_template` - Address6 Template.
* `subnet_segment` - Subnet Segment.

* `name` - Subnet segment value name.
* `value` - Subnet segment value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Address6TemplateSubnetSegmentValues can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "address6_template=YOUR_VALUE", "subnet_segment=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_address6template_subnetsegment_values.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

