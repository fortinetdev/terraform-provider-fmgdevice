---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_sdwan_duplication"
description: |-
  Create SD-WAN duplication rule.
---

# fmgdevice_system_sdwan_duplication
Create SD-WAN duplication rule.

~> This resource is a sub resource for variable `duplication` of resource `fmgdevice_system_sdwan`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_sdwan_duplication" "trname" {
  dstaddr               = ["your own value"]
  dstaddr6              = ["your own value"]
  dstintf               = ["your own value"]
  fosid                 = 10
  packet_de_duplication = "enable"
  device_name           = var.device_name # not required if setting is at provider
  device_vdom           = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `dstaddr` - Destination address or address group names.
* `dstaddr6` - Destination address6 or address6 group names.
* `dstintf` - Outgoing (egress) interfaces or zones.
* `fosid` - Duplication rule ID (1 - 255).
* `packet_de_duplication` - Enable/disable discarding of packets that have been duplicated. Valid values: `disable`, `enable`.

* `packet_duplication` - Configure packet duplication method. Valid values: `disable`, `force`, `on-demand`.

* `service` - Service and service group name.
* `service_id` - SD-WAN service rule ID list.
* `sla_match_service` - Enable/disable packet duplication matching health-check SLAs in service rule. Valid values: `disable`, `enable`.

* `srcaddr` - Source address or address group names.
* `srcaddr6` - Source address6 or address6 group names.
* `srcintf` - Incoming (ingress) interfaces or zones.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SdwanDuplication can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_sdwan_duplication.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

