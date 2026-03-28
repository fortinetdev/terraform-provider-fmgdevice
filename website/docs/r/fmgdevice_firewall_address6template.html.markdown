---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_address6template"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure IPv6 address templates.
---

# fmgdevice_firewall_address6template
<i>This object will be purged after policy copy and install.</i> Configure IPv6 address templates.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `subnet_segment`: `fmgdevice_firewall_address6template_subnetsegment`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `_image_base64` - _Image-Base64.
* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `ip6` - IPv6 address prefix.
* `name` - IPv6 address template name.
* `subnet_segment` - Subnet-Segment. The structure of `subnet_segment` block is documented below.
* `subnet_segment_count` - Number of IPv6 subnet segments.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `subnet_segment` block supports:

* `bits` - Number of bits.
* `exclusive` - Enable/disable exclusive value. Valid values: `disable`, `enable`.

* `id` - Subnet segment ID.
* `name` - Subnet segment name.
* `values` - Values. The structure of `values` block is documented below.

The `values` block supports:

* `name` - Subnet segment value name.
* `value` - Subnet segment value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Address6Template can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_address6template.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

