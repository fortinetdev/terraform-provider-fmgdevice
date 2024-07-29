---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_securitypolicy_localaccess"
description: |-
  Configure allowaccess list for mgmt and internal interfaces on managed FortiSwitch units.
---

# fmgdevice_switchcontroller_securitypolicy_localaccess
Configure allowaccess list for mgmt and internal interfaces on managed FortiSwitch units.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_securitypolicy_localaccess" "trname" {
  internal_allowaccess = ["ping"]
  mgmt_allowaccess     = ["ping"]
  name                 = "your own value"
  device_name          = var.device_name # not required if setting is at provider
  device_vdom          = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `internal_allowaccess` - Allowed access on the switch internal interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.

* `mgmt_allowaccess` - Allowed access on the switch management interface. Valid values: `https`, `ping`, `ssh`, `snmp`, `http`, `telnet`, `radius-acct`.

* `name` - Policy name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController SecurityPolicyLocalAccess can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_securitypolicy_localaccess.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

