---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_vlanpolicy"
description: |-
  Configure VLAN policy to be applied on the managed FortiSwitch ports through dynamic-port-policy.
---

# fmgdevice_switchcontroller_vlanpolicy
Configure VLAN policy to be applied on the managed FortiSwitch ports through dynamic-port-policy.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_vlanpolicy" "trname" {
  allowed_vlans     = ["your own value"]
  allowed_vlans_all = "enable"
  description       = "your own value"
  discard_mode      = "all-tagged"
  fortilink         = ["your own value"]
  name              = "your own value"
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `allowed_vlans` - Allowed VLANs to be applied when using this VLAN policy.
* `allowed_vlans_all` - Enable/disable all defined VLANs when using this VLAN policy. Valid values: `disable`, `enable`.

* `description` - Description for the VLAN policy.
* `discard_mode` - Discard mode to be applied when using this VLAN policy. Valid values: `none`, `all-untagged`, `all-tagged`.

* `fortilink` - FortiLink interface for which this VLAN policy belongs to.
* `name` - VLAN policy name.
* `untagged_vlans` - Untagged VLANs to be applied when using this VLAN policy.
* `vlan` - Native VLAN to be applied when using this VLAN policy.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController VlanPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_vlanpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

