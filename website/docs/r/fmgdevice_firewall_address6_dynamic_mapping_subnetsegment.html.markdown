---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_address6_dynamic_mapping_subnetsegment"
description: |-
  <i>This object will be purged after policy copy and install.</i> IPv6 subnet segments.
---

# fmgdevice_firewall_address6_dynamic_mapping_subnetsegment
<i>This object will be purged after policy copy and install.</i> IPv6 subnet segments.

~> This resource is a sub resource for variable `subnet_segment` of resource `fmgdevice_firewall_address6_dynamic_mapping`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `address6` - Address6.
* `dynamic_mapping_name` - Dynamic Mapping Name.
* `dynamic_mapping_vdom` - Dynamic Mapping Vdom.

* `name` - Name.
* `type` - Subnet segment type. Valid values: `any`, `specific`.

* `value` - Subnet segment value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Address6DynamicMappingSubnetSegment can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "address6=YOUR_VALUE", "dynamic_mapping_name=YOUR_VALUE", "dynamic_mapping_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_address6_dynamic_mapping_subnetsegment.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

