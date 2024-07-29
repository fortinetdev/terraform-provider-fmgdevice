---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_snmpsysinfo"
description: |-
  Configuration method to edit Simple Network Management Protocol (SNMP) system info.
---

# fmgdevice_switchcontroller_managedswitch_snmpsysinfo
Configuration method to edit Simple Network Management Protocol (SNMP) system info.

~> This resource is a sub resource for variable `snmp_sysinfo` of resource `fmgdevice_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_managedswitch_snmpsysinfo" "trname" {
  contact_info = "your own value"
  description  = "your own value"
  engine_id    = "your own value"
  location     = "your own value"
  status       = "enable"
  device_name  = var.device_name # not required if setting is at provider
  device_vdom  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `contact_info` - Contact information.
* `description` - System description.
* `engine_id` - Local SNMP engine ID string (max 24 char).
* `location` - System location.
* `status` - Enable/disable SNMP. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController ManagedSwitchSnmpSysinfo can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_snmpsysinfo.labelname SwitchControllerManagedSwitchSnmpSysinfo
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

