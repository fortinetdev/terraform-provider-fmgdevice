---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_vap_vlanname"
description: |-
  Table for mapping VLAN name to VLAN ID.
---

# fmgdevice_wirelesscontroller_vap_vlanname
Table for mapping VLAN name to VLAN ID.

~> This resource is a sub resource for variable `vlan_name` of resource `fmgdevice_wirelesscontroller_vap`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_vap_vlanname" "trname" {
  name        = "your own value"
  vlan_id     = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `vap` - Vap.

* `name` - VLAN name.
* `vlan_id` - VLAN IDs (maximum 8 VLAN IDs).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController VapVlanName can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "vap=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_vap_vlanname.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

