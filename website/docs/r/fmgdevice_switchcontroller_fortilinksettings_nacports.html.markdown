---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_fortilinksettings_nacports"
description: |-
  NAC specific configuration.
---

# fmgdevice_switchcontroller_fortilinksettings_nacports
NAC specific configuration.

~> This resource is a sub resource for variable `nac_ports` of resource `fmgdevice_switchcontroller_fortilinksettings`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_fortilinksettings_nacports" "trname" {
  bounce_nac_port   = "enable"
  lan_segment       = "disabled"
  member_change     = 10
  nac_lan_interface = ["port2"]
  nac_segment_vlans = ["your own value"]
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `fortilink_settings` - Fortilink Settings.

* `bounce_nac_port` - Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.

* `lan_segment` - Enable/disable LAN segment feature on the FortiLink interface. Valid values: `disabled`, `enabled`.

* `member_change` - Member-Change.
* `nac_lan_interface` - Configure NAC LAN interface.
* `nac_segment_vlans` - Configure NAC segment VLANs.
* `onboarding_vlan` - Default NAC Onboarding VLAN when NAC devices are discovered.
* `parent_key` - Parent-Key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController FortilinkSettingsNacPorts can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "fortilink_settings=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_fortilinksettings_nacports.labelname SwitchControllerFortilinkSettingsNacPorts
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

