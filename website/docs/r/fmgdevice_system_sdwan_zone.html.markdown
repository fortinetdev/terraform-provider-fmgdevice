---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sdwan_zone"
description: |-
  Configure SD-WAN zones.
---

# fmgdevice_system_sdwan_zone
Configure SD-WAN zones.

~> This resource is a sub resource for variable `zone` of resource `fmgdevice_system_sdwan`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_sdwan_zone" "trname" {
  advpn_health_check       = ["your own value"]
  advpn_select             = "enable"
  minimum_sla_meet_members = 10
  name                     = "your own value"
  service_sla_tie_break    = "cfg-order"
  device_name              = var.device_name # not required if setting is at provider
  device_vdom              = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `advpn_health_check` - Health check for ADVPN local overlay link quality.
* `advpn_select` - Enable/disable selection of ADVPN based on SDWAN information. Valid values: `disable`, `enable`.

* `minimum_sla_meet_members` - Minimum number of members which meet SLA when the neighbor is preferred.
* `name` - Zone name.
* `service_sla_tie_break` - Method of selecting member if more than one meets the SLA. Valid values: `cfg-order`, `fib-best-match`, `input-device`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SdwanZone can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sdwan_zone.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

