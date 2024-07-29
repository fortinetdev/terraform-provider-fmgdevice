---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_apcfgprofile"
description: |-
  Configure AP local configuration profiles.
---

# fmgdevice_wirelesscontroller_apcfgprofile
Configure AP local configuration profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `command_list`: `fmgdevice_wirelesscontroller_apcfgprofile_commandlist`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_apcfgprofile" "trname" {
  ac_ip       = "your own value"
  ac_port     = 10
  ac_timer    = 10
  ac_type     = "specify"
  ap_family   = "fap-c"
  name        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `ac_ip` - IP address of the validation controller that AP must be able to join after applying AP local configuration.
* `ac_port` - Port of the validation controller that AP must be able to join after applying AP local configuration (1024 - 49150, default = 5246).
* `ac_timer` - Maximum waiting time for the AP to join the validation controller after applying AP local configuration (3 - 30 min, default = 10).
* `ac_type` - Validation controller type (default = default). Valid values: `default`, `specify`, `apcfg`.

* `ap_family` - FortiAP family type (default = fap). Valid values: `fap`, `fap-u`, `fap-c`.

* `command_list` - Command-List. The structure of `command_list` block is documented below.
* `comment` - Comment.
* `name` - AP local configuration profile name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `command_list` block supports:

* `id` - Command ID.
* `name` - AP local configuration command name.
* `passwd_value` - AP local configuration command password value.
* `type` - The command type (default = non-password). Valid values: `non-password`, `password`.

* `value` - AP local configuration command value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController ApcfgProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_apcfgprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

