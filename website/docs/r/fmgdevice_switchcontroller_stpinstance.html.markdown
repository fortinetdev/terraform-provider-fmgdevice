---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_stpinstance"
description: |-
  Configure FortiSwitch multiple spanning tree protocol (MSTP) instances.
---

# fmgdevice_switchcontroller_stpinstance
Configure FortiSwitch multiple spanning tree protocol (MSTP) instances.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_stpinstance" "trname" {
  fosid       = "your own value"
  vlan_range  = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `fosid` - Instance ID.
* `vlan_range` - Configure VLAN range for STP instance.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SwitchController StpInstance can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_stpinstance.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

