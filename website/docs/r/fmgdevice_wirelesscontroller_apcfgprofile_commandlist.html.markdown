---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_apcfgprofile_commandlist"
description: |-
  AP local configuration command list.
---

# fmgdevice_wirelesscontroller_apcfgprofile_commandlist
AP local configuration command list.

~> This resource is a sub resource for variable `command_list` of resource `fmgdevice_wirelesscontroller_apcfgprofile`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_apcfgprofile_commandlist" "trname" {
  fosid        = 10
  name         = "your own value"
  passwd_value = ["your own value"]
  type         = "non-password"
  value        = "your own value"
  device_name  = var.device_name # not required if setting is at provider
  device_vdom  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `apcfg_profile` - Apcfg Profile.

* `fosid` - Command ID.
* `name` - AP local configuration command name.
* `passwd_value` - AP local configuration command password value.
* `type` - The command type (default = non-password). Valid values: `non-password`, `password`.

* `value` - AP local configuration command value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

WirelessController ApcfgProfileCommandList can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "apcfg_profile=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_apcfgprofile_commandlist.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

