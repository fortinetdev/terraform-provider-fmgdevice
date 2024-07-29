---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_policy6"
description: |-
  Configure IPv6 routing policies.
---

# fmgdevice_router_policy6
Configure IPv6 routing policies.

## Example Usage

```hcl
resource "fmgdevice_router_policy6" "trname" {
  action      = "deny"
  comments    = "your own value"
  dst         = ["your own value"]
  dst_negate  = "disable"
  dstaddr     = ["your own value"]
  seq_num     = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `action` - Action of the policy route. Valid values: `permit`, `deny`.

* `comments` - Optional comments.
* `dst` - Destination IPv6 prefix.
* `dst_negate` - Enable/disable negating destination address match. Valid values: `disable`, `enable`.

* `dstaddr` - Destination address name.
* `end_port` - End destination port number (1 - 65535).
* `end_source_port` - End source port number (1 - 65535).
* `gateway` - IPv6 address of the gateway.
* `input_device` - Incoming interface name.
* `input_device_negate` - Enable/disable negation of input device match. Valid values: `disable`, `enable`.

* `internet_service_custom` - Custom Destination Internet Service name.
* `internet_service_id` - Destination Internet Service ID.
* `output_device` - Outgoing interface name.
* `protocol` - Protocol number (0 - 255).
* `seq_num` - Sequence number(1-65535).
* `src` - Source IPv6 prefix.
* `src_negate` - Enable/disable negating source address match. Valid values: `disable`, `enable`.

* `srcaddr` - Source address name.
* `start_port` - Start destination port number (1 - 65535).
* `start_source_port` - Start source port number (1 - 65535).
* `status` - Enable/disable this policy route. Valid values: `disable`, `enable`.

* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Router Policy6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_policy6.labelname {{seq_num}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

