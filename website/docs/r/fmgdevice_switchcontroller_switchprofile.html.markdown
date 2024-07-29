---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_switchprofile"
description: |-
  Configure FortiSwitch switch profile.
---

# fmgdevice_switchcontroller_switchprofile
Configure FortiSwitch switch profile.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_switchprofile" "trname" {
  login                     = "enable"
  login_passwd              = ["your own value"]
  login_passwd_override     = "disable"
  name                      = "your own value"
  revision_backup_on_logout = "disable"
  device_name               = var.device_name # not required if setting is at provider
  device_vdom               = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `login` - Enable/disable FortiSwitch serial console. Valid values: `disable`, `enable`.

* `login_passwd` - Login password of managed FortiSwitch.
* `login_passwd_override` - Enable/disable overriding the admin administrator password for a managed FortiSwitch with the FortiGate admin administrator account password. Valid values: `disable`, `enable`.

* `name` - FortiSwitch Profile name.
* `revision_backup_on_logout` - Enable/disable automatic revision backup upon logout from FortiSwitch. Valid values: `disable`, `enable`.

* `revision_backup_on_upgrade` - Enable/disable automatic revision backup upon FortiSwitch image upgrade. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController SwitchProfile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_switchprofile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

