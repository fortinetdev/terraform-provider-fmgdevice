---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_bonjourprofile"
description: |-
  Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.
---

# fmgdevice_wirelesscontroller_bonjourprofile
Configure Bonjour profiles. Bonjour is Apple's zero configuration networking protocol. Bonjour profiles allow APs and FortiAPs to connnect to networks using Bonjour.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `policy_list`: `fmgdevice_wirelesscontroller_bonjourprofile_policylist`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_bonjourprofile" "trname" {
  comment = "your own value"
  name    = "your own value"
  policy_list {
    description = "your own value"
    from_vlan   = "your own value"
    policy_id   = 10
    services    = ["printers"]
    to_vlan     = "your own value"
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `micro_location` - Enable/disable Micro location for Bonjour profile (default = disable). Valid values: `disable`, `enable`.

* `name` - Bonjour profile name.
* `policy_list` - Policy-List. The structure of `policy_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `policy_list` block supports:

* `description` - Description.
* `from_vlan` - VLAN ID from which the Bonjour service is advertised (0 - 4094, default = 0).
* `policy_id` - Policy ID.
* `services` - Bonjour services for the VLAN connecting to the Bonjour network. Valid values: `airplay`, `afp`, `bit-torrent`, `ftp`, `ichat`, `itunes`, `printers`, `samba`, `scanners`, `ssh`, `chromecast`, `all`, `miracast`.

* `to_vlan` - VLAN ID to which the Bonjour service is made available (0 - 4094, default = all).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController BonjourProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_bonjourprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

