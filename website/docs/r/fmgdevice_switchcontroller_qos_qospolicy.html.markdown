---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_qos_qospolicy"
description: |-
  Configure FortiSwitch QoS policy.
---

# fmgdevice_switchcontroller_qos_qospolicy
Configure FortiSwitch QoS policy.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_qos_qospolicy" "trname" {
  default_cos       = 10
  name              = "your own value"
  queue_policy      = ["your own value"]
  trust_dot1p_map   = ["your own value"]
  trust_ip_dscp_map = ["your own value"]
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `default_cos` - Default cos queue for untagged packets.
* `name` - QoS policy name.
* `queue_policy` - QoS egress queue policy.
* `trust_dot1p_map` - QoS trust 802.1p map.
* `trust_ip_dscp_map` - QoS trust ip dscp map.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController QosQosPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_qos_qospolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

