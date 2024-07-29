---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_macpolicy"
description: |-
  Configure MAC policy to be applied on the managed FortiSwitch devices through NAC device.
---

# fmgdevice_switchcontroller_macpolicy
Configure MAC policy to be applied on the managed FortiSwitch devices through NAC device.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_macpolicy" "trname" {
  bounce_port_link = "enable"
  count            = "disable"
  description      = "your own value"
  drop             = "enable"
  fortilink        = ["your own value"]
  name             = "your own value"
  device_name      = var.device_name # not required if setting is at provider
  device_vdom      = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `bounce_port_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where this mac-policy is applied. Valid values: `disable`, `enable`.

* `fmgcount` - Enable/disable packet count on the NAC device. Valid values: `disable`, `enable`.

* `description` - Description for the MAC policy.
* `drop` - Drop. Valid values: `disable`, `enable`.

* `fortilink` - FortiLink interface for which this MAC policy belongs to.
* `name` - MAC policy name.
* `traffic_policy` - Traffic policy to be applied when using this MAC policy.
* `vlan` - Ingress traffic VLAN assignment for the MAC address matching this MAC policy.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController MacPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_macpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

