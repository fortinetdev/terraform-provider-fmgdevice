---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_managedswitch_ipsourceguard"
description: |-
  IP source guard.
---

# fmgdevice_switchcontroller_managedswitch_ipsourceguard
IP source guard.

~> This resource is a sub resource for variable `ip_source_guard` of resource `fmgdevice_switchcontroller_managedswitch`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `binding_entry`: `fmgdevice_switchcontroller_managedswitch_ipsourceguard_bindingentry`



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_managedswitch_ipsourceguard" "trname" {
  binding_entry {
    entry_name = "your own value"
    ip         = "your own value"
    mac        = "your own value"
  }

  description = "your own value"
  port        = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `managed_switch` - Managed Switch.

* `binding_entry` - Binding-Entry. The structure of `binding_entry` block is documented below.
* `description` - Description.
* `port` - Ingress interface to which source guard is bound.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `binding_entry` block supports:

* `entry_name` - Configure binding pair.
* `ip` - Source IP for this rule.
* `mac` - MAC address for this rule.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{port}}.

## Import

SwitchController ManagedSwitchIpSourceGuard can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "managed_switch=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_managedswitch_ipsourceguard.labelname {{port}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

