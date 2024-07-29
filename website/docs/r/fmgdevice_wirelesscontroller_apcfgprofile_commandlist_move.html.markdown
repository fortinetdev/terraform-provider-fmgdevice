---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_apcfgprofile_commandlist_move"
description: |-
  AP local configuration command list.
---

# fmgdevice_wirelesscontroller_apcfgprofile_commandlist_move
AP local configuration command list.

## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_apcfgprofile_commandlist_move" "trname" {
  target        = 10
  command_list  = 11
  option        = "after"
  apcfg_profile = fmgdevice_wirelesscontroller_apcfgprofile.trname1.name
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
  depends_on    = [fmgdevice_wirelesscontroller_apcfgprofile_commandlist.trname1, fmgdevice_wirelesscontroller_apcfgprofile_commandlist.trname2]
}

resource "fmgdevice_wirelesscontroller_apcfgprofile_commandlist" "trname2" {
  fosid         = 10
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
  apcfg_profile = fmgdevice_wirelesscontroller_apcfgprofile.trname1.name
  depends_on    = [fmgdevice_wirelesscontroller_apcfgprofile.trname1]
}

resource "fmgdevice_wirelesscontroller_apcfgprofile_commandlist" "trname1" {
  fosid         = 11
  device_name   = var.device_name # not required if setting is at provider
  device_vdom   = var.device_vdom # not required if setting is at provider
  apcfg_profile = fmgdevice_wirelesscontroller_apcfgprofile.trname1.name
  depends_on    = [fmgdevice_wirelesscontroller_apcfgprofile.trname1]
}

resource "fmgdevice_wirelesscontroller_apcfgprofile" "trname1" {
  name        = "test3"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `apcfg_profile` - Apcfg Profile.
* `command_list` - Command List.

* `target` - Key to the target entry.
* `option` - Option. Valid values: `before`, `after`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.
* `state_pos` - The parameter is read-only, it is used to get the lastest relative position of the two items. This can help check whether the latest relative position of the two items matches the configuration, and help check whether they have been deleted. If the latest relative position of the two items matches the configuration, the value of state_pos is an empty string.

## Others

~> **Warning:** Since the command list changes caused by modifications outside the terraform may be beyond the control of the resource, terraform destroy for the resource will not restore the original sequence state of command lists.
