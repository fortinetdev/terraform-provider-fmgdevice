---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_initialconfig_vlans"
description: |-
  Configure initial template for auto-generated VLAN interfaces.
---

# fmgdevice_switchcontroller_initialconfig_vlans
Configure initial template for auto-generated VLAN interfaces.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_initialconfig_vlans" "trname" {
  default_vlan = ["your own value"]
  nac          = ["your own value"]
  nac_segment  = ["your own value"]
  quarantine   = ["your own value"]
  rspan        = ["your own value"]
  device_name  = var.device_name # not required if setting is at provider
  device_vdom  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `default_vlan` - Default VLAN (native) assigned to all switch ports upon discovery.
* `nac` - VLAN for NAC onboarding devices.
* `nac_segment` - VLAN for NAC segment primary interface.
* `quarantine` - VLAN for quarantined traffic.
* `rspan` - VLAN for RSPAN/ERSPAN mirrored traffic.
* `video` - VLAN dedicated for video devices.
* `voice` - VLAN dedicated for voice devices.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController InitialConfigVlans can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_initialconfig_vlans.labelname SwitchControllerInitialConfigVlans
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

