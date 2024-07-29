---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_fortilinksettings"
description: |-
  Configure integrated FortiLink settings for FortiSwitch.
---

# fmgdevice_switchcontroller_fortilinksettings
Configure integrated FortiLink settings for FortiSwitch.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `nac_ports`: `fmgdevice_switchcontroller_fortilinksettings_nacports`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_fortilinksettings" "trname" {
  access_vlan_mode = "legacy"
  fortilink        = ["your own value"]
  inactive_timer   = 10
  link_down_flush  = "enable"
  nac_ports {
    bounce_nac_port   = "enable"
    lan_segment       = "disabled"
    member_change     = 10
    nac_lan_interface = ["port2"]
    nac_segment_vlans = ["your own value"]
    onboarding_vlan   = ["your own value"]
    parent_key        = "your own value"
  }

  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `access_vlan_mode` - Intra VLAN traffic behavior with loss of connection to the FortiGate. Valid values: `legacy`, `fail-open`, `fail-close`.

* `fortilink` - FortiLink interface to which this fortilink-setting belongs.
* `inactive_timer` - Time interval(minutes) to be included in the inactive devices expiry calculation (mac age-out + inactive-time + periodic scan interval).
* `link_down_flush` - Clear NAC and dynamic devices on switch ports on link down event. Valid values: `disable`, `enable`.

* `nac_ports` - Nac-Ports. The structure of `nac_ports` block is documented below.
* `name` - FortiLink settings name.

The `nac_ports` block supports:

* `bounce_nac_port` - Enable/disable bouncing (administratively bring the link down, up) of a switch port when NAC mode is configured on the port. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.

* `lan_segment` - Enable/disable LAN segment feature on the FortiLink interface. Valid values: `disabled`, `enabled`.

* `member_change` - Member-Change.
* `nac_lan_interface` - Configure NAC LAN interface.
* `nac_segment_vlans` - Configure NAC segment VLANs.
* `onboarding_vlan` - Default NAC Onboarding VLAN when NAC devices are discovered.
* `parent_key` - Parent-Key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController FortilinkSettings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_fortilinksettings.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

