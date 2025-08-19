---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_extensioncontroller_extender"
description: |-
  Extender controller configuration.
---

# fmgdevice_extensioncontroller_extender
Extender controller configuration.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `wan_extension`: `fmgdevice_extensioncontroller_extender_wanextension`



## Example Usage

```hcl
resource "fmgdevice_extensioncontroller_extender" "trname" {
  allowaccess     = ["ssh"]
  authorized      = "enable"
  bandwidth_limit = 10
  description     = "your own value"
  device_id       = 10
  name            = "your own value"
  device_name     = var.device_name # not required if setting is at provider
  device_vdom     = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `allowaccess` - Control management access to the managed extender. Separate entries with a space. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`.

* `authorized` - FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`, `discovered`.

* `bandwidth_limit` - FortiExtender LAN extension bandwidth limit (Mbps).
* `description` - Description.
* `device_id` - Device ID.
* `enforce_bandwidth` - Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `disable`, `enable`.

* `ext_name` - FortiExtender name.
* `extension_type` - Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.

* `firmware_provision_latest` - Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.

* `fosid` - FortiExtender serial number.
* `login_password` - Set the managed extender's administrator password.
* `login_password_change` - Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `no`, `yes`, `default`.

* `name` - FortiExtender entry name.
* `override_allowaccess` - Enable to override the extender profile management access configuration. Valid values: `disable`, `enable`.

* `override_enforce_bandwidth` - Enable to override the extender profile enforce-bandwidth setting. Valid values: `disable`, `enable`.

* `override_login_password_change` - Enable to override the extender profile login-password (administrator password) setting. Valid values: `disable`, `enable`.

* `profile` - FortiExtender profile configuration.
* `vdom` - Vdom.
* `wan_extension` - Wan-Extension. The structure of `wan_extension` block is documented below.

The `wan_extension` block supports:

* `modem1_extension` - FortiExtender interface name.
* `modem1_pdn1_interface` - FortiExtender interface name.
* `modem1_pdn2_interface` - FortiExtender interface name.
* `modem1_pdn3_interface` - FortiExtender interface name.
* `modem1_pdn4_interface` - FortiExtender interface name.
* `modem2_extension` - FortiExtender interface name.
* `modem2_pdn1_interface` - FortiExtender interface name.
* `modem2_pdn2_interface` - FortiExtender interface name.
* `modem2_pdn3_interface` - FortiExtender interface name.
* `modem2_pdn4_interface` - FortiExtender interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtensionController Extender can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_extensioncontroller_extender.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

