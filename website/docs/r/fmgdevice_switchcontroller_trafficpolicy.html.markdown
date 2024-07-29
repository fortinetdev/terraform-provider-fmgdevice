---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_switchcontroller_trafficpolicy"
description: |-
  Configure FortiSwitch traffic policy.
---

# fmgdevice_switchcontroller_trafficpolicy
Configure FortiSwitch traffic policy.

## Example Usage

```hcl
resource "fmgdevice_switchcontroller_trafficpolicy" "trname" {
  cos_queue            = 10
  description          = "your own value"
  guaranteed_bandwidth = 10
  guaranteed_burst     = 10
  fosid                = 10
  name                 = "your own value"
  device_name          = var.device_name # not required if setting is at provider
  device_vdom          = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `cos_queue` - COS queue(0 - 7), or unset to disable.
* `description` - Description of the traffic policy.
* `guaranteed_bandwidth` - Guaranteed bandwidth in kbps (max value = 524287000).
* `guaranteed_burst` - Guaranteed burst size in bytes (max value = 4294967295).
* `fosid` - FSW Policer id
* `maximum_burst` - Maximum burst size in bytes (max value = 4294967295).
* `name` - Traffic policy name.
* `policer_status` - Enable/disable policer config on the traffic policy. Valid values: `disable`, `enable`.

* `type` - Type. Valid values: `ingress`, `egress`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController TrafficPolicy can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_switchcontroller_trafficpolicy.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

