---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_ipsourceguard_bindingentry"
description: |-
  IP and MAC address configuration.
---

# fmgdevice_switchcontroller_managedswitch_ipsourceguard_bindingentry
IP and MAC address configuration.

~> This resource is a sub resource for variable `binding_entry` of resource `fmgdevice_switchcontroller_managedswitch_ipsourceguard`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_managedswitch_ipsourceguard_bindingentry" "trname" {
  entry_name  = "your own value"
  ip          = "your own value"
  mac         = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.
* `ip_source_guard` - Ip Source Guard.

* `entry_name` - Configure binding pair.
* `ip` - Source IP for this rule.
* `mac` - MAC address for this rule.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{entry_name}}.

## Import

SwitchController ManagedSwitchIpSourceGuardBindingEntry can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE", "ip_source_guard=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_ipsourceguard_bindingentry.labelname {{entry_name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

