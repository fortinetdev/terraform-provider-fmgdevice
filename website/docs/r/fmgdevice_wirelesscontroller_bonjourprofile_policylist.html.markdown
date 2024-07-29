---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_bonjourprofile_policylist"
description: |-
  Bonjour policy list.
---

# fmgdevice_wirelesscontroller_bonjourprofile_policylist
Bonjour policy list.

~> This resource is a sub resource for variable `policy_list` of resource `fmgdevice_wirelesscontroller_bonjourprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_bonjourprofile_policylist" "trname" {
  description = "your own value"
  from_vlan   = "your own value"
  policy_id   = 10
  services    = ["scanners"]
  to_vlan     = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `bonjour_profile` - Bonjour Profile.

* `description` - Description.
* `from_vlan` - VLAN ID from which the Bonjour service is advertised (0 - 4094, default = 0).
* `policy_id` - Policy ID.
* `services` - Bonjour services for the VLAN connecting to the Bonjour network. Valid values: `airplay`, `afp`, `bit-torrent`, `ftp`, `ichat`, `itunes`, `printers`, `samba`, `scanners`, `ssh`, `chromecast`, `all`, `miracast`.

* `to_vlan` - VLAN ID to which the Bonjour service is made available (0 - 4094, default = all).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policy_id}}.

## Import

WirelessController BonjourProfilePolicyList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "bonjour_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_bonjourprofile_policylist.labelname {{policy_id}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

