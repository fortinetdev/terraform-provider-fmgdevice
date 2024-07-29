---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ipsecaggregate"
description: |-
  Configure an aggregate of IPsec tunnels.
---

# fmgdevice_system_ipsecaggregate
Configure an aggregate of IPsec tunnels.

## Example Usage

```hcl
resource "fmgdevice_system_ipsecaggregate" "trname" {
  add_slbc_route = "disable"
  algorithm      = "round-robin"
  member         = ["your own value"]
  name           = "your own value"
  device_name    = var.device_name # not required if setting is at provider
  device_vdom    = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `add_slbc_route` - Enable/disable add IPsec routes to the SLBC load-balancer. Valid values: `disable`, `enable`.

* `algorithm` - Frame distribution algorithm. Valid values: `L3`, `L4`, `round-robin`, `redundant`, `weighted-round-robin`.

* `member` - Member tunnels of the aggregate.
* `name` - IPsec aggregate name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System IpsecAggregate can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ipsecaggregate.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

