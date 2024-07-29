---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_dynamicportpolicy_policy"
description: |-
  Port policies with matching criteria and actions.
---

# fmgdevice_switchcontroller_dynamicportpolicy_policy
Port policies with matching criteria and actions.

~> This resource is a sub resource for variable `policy` of resource `fmgdevice_switchcontroller_dynamicportpolicy`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_dynamicportpolicy_policy" "trname" {
  n802_1x          = ["your own value"]
  bounce_port_link = "disable"
  category         = "device"
  description      = "your own value"
  family           = "your own value"
  name             = "your own value"
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `dynamic_port_policy` - Dynamic Port Policy.

* `n802_1x` - 802.1x security policy to be applied when using this policy.
* `bounce_port_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where this policy is applied. Helps to clear and reassign VLAN from lldp-profile. Valid values: `disable`, `enable`.

* `category` - Category of Dynamic port policy. Valid values: `device`, `interface-tag`.

* `description` - Description for the policy.
* `family` - Match policy based on family.
* `host` - Match policy based on host.
* `hw_vendor` - Match policy based on hardware vendor.
* `interface_tags` - Match policy based on the FortiSwitch interface object tags.
* `lldp_profile` - LLDP profile to be applied when using this policy.
* `mac` - Match policy based on MAC address.
* `match_period` - Number of days the matched devices will be retained (0 - 120, 0 = always retain).
* `match_type` - Match and retain the devices based on the type. Valid values: `dynamic`, `override`.

* `name` - Policy name.
* `qos_policy` - QoS policy to be applied when using this policy.
* `status` - Enable/disable policy. Valid values: `disable`, `enable`.

* `type` - Match policy based on type.
* `vlan_policy` - VLAN policy to be applied when using this policy.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController DynamicPortPolicyPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "dynamic_port_policy=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_dynamicportpolicy_policy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

