---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_switchinterface"
description: |-
  Configure software switch interfaces by grouping physical and WiFi interfaces.
---

# fmgdevice_system_switchinterface
Configure software switch interfaces by grouping physical and WiFi interfaces.

## Example Usage

```hcl
resource "fmgdevice_system_switchinterface" "trname" {
  intra_switch_policy = "explicit"
  mac_ttl             = 10
  member              = ["your own value"]
  name                = "your own value"
  span                = "disable"
  device_name         = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `intra_switch_policy` - Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.

* `mac_ttl` - Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
* `member` - Names of the interfaces that belong to the virtual switch.
* `name` - Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
* `span` - Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.

* `span_dest_port` - SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
* `span_direction` - The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.

* `span_source_port` - Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port.
* `type` - Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`, `hardware-switch`.

* `vdom` - VDOM that the software switch belongs to.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SwitchInterface can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_switchinterface.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

