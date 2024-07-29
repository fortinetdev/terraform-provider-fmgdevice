---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_federatedupgrade_knownhamembers"
description: |-
  Known members of the HA cluster. If a member is missing at upgrade time, the upgrade will be cancelled.
---

# fmgdevice_system_federatedupgrade_knownhamembers
Known members of the HA cluster. If a member is missing at upgrade time, the upgrade will be cancelled.

~> This resource is a sub resource for variable `known_ha_members` of resource `fmgdevice_system_federatedupgrade`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_federatedupgrade_knownhamembers" "trname" {
  serial      = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `serial` - Serial number of HA member


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{serial}}.

## Import

System FederatedUpgradeKnownHaMembers can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_federatedupgrade_knownhamembers.labelname {{serial}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

