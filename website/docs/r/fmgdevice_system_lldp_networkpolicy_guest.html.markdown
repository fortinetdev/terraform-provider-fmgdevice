---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_lldp_networkpolicy_guest"
description: |-
  Guest.
---

# fmgdevice_system_lldp_networkpolicy_guest
Guest.

~> This resource is a sub resource for variable `guest` of resource `fmgdevice_system_lldp_networkpolicy`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_lldp_networkpolicy_guest" "trname" {
  dscp        = 10
  priority    = 10
  status      = "disable"
  tag         = "dot1q"
  vlan        = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `network_policy` - Network Policy.

* `dscp` - Differentiated Services Code Point (DSCP) value to advertise.
* `priority` - 802.1P CoS/PCP to advertise (0 - 7; from lowest to highest priority).
* `status` - Enable/disable advertising this policy. Valid values: `disable`, `enable`.

* `tag` - Advertise tagged or untagged traffic. Valid values: `none`, `dot1q`, `dot1p`.

* `vlan` - 802.1Q VLAN ID to advertise (1 - 4094).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System LldpNetworkPolicyGuest can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "network_policy=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_lldp_networkpolicy_guest.labelname SystemLldpNetworkPolicyGuest
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

