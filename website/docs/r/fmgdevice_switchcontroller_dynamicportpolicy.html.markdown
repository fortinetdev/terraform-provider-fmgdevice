---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_dynamicportpolicy"
description: |-
  Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.
---

# fmgdevice_switchcontroller_dynamicportpolicy
Configure Dynamic port policy to be applied on the managed FortiSwitch ports through DPP device.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `policy`: `fmgdevice_switchcontroller_dynamicportpolicy_policy`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_dynamicportpolicy" "trname" {
  description = "your own value"
  fortilink   = ["your own value"]
  name        = "your own value"
  policy {
    n802_1x          = ["your own value"]
    bounce_port_link = "enable"
    category         = "device"
    description      = "your own value"
    family           = "your own value"
    host             = "your own value"
    hw_vendor        = "your own value"
    interface_tags   = ["port2"]
    lldp_profile     = ["your own value"]
    mac              = "your own value"
    match_period     = 10
    match_type       = "override"
    name             = "your own value"
    qos_policy       = ["your own value"]
    status           = "enable"
    type             = "your own value"
    vlan_policy      = ["your own value"]
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `description` - Description for the Dynamic port policy.
* `fortilink` - FortiLink interface for which this Dynamic port policy belongs to.
* `name` - Dynamic port policy name.
* `policy` - Policy. The structure of `policy` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `policy` block supports:

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

SwitchController DynamicPortPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_dynamicportpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

