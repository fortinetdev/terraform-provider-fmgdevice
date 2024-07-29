---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_qos_dot1pmap"
description: |-
  Configure FortiSwitch QoS 802.1p.
---

# fmgdevice_switchcontroller_qos_dot1pmap
Configure FortiSwitch QoS 802.1p.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_qos_dot1pmap" "trname" {
  description        = "your own value"
  egress_pri_tagging = "enable"
  name               = "your own value"
  priority_0         = "queue-3"
  priority_1         = "queue-6"
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `description` - Description of the 802.1p name.
* `egress_pri_tagging` - Enable/disable egress priority-tag frame. Valid values: `disable`, `enable`.

* `name` - Dot1p map name.
* `priority_0` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.

* `priority_1` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.

* `priority_2` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.

* `priority_3` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.

* `priority_4` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.

* `priority_5` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.

* `priority_6` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.

* `priority_7` - COS queue mapped to dot1p priority number. Valid values: `queue-0`, `queue-1`, `queue-2`, `queue-3`, `queue-4`, `queue-5`, `queue-6`, `queue-7`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController QosDot1PMap can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_qos_dot1pmap.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

