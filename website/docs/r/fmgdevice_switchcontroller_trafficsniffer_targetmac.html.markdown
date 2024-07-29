---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_trafficsniffer_targetmac"
description: |-
  Sniffer MACs to filter.
---

# fmgdevice_switchcontroller_trafficsniffer_targetmac
Sniffer MACs to filter.

~> This resource is a sub resource for variable `target_mac` of resource `fmgdevice_switchcontroller_trafficsniffer`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_switchcontroller_trafficsniffer_targetmac" "trname" {
  description  = "your own value"
  dst_entry_id = 10
  mac          = "your own value"
  src_entry_id = 10
  device_name  = var.device_name # not required if setting is at provider
  device_vdom  = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `description` - Description for the sniffer MAC.
* `dst_entry_id` - FortiSwitch dest entry ID for the sniffer MAC.
* `mac` - Sniffer MAC.
* `src_entry_id` - FortiSwitch source entry ID for the sniffer MAC.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{mac}}.

## Import

SwitchController TrafficSnifferTargetMac can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_trafficsniffer_targetmac.labelname {{mac}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

