---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_evpn"
description: |-
  Configure EVPN instance.
---

# fmgdevice_system_evpn
Configure EVPN instance.

## Example Usage

```hcl
resource "fmgdevice_system_evpn" "trname" {
  arp_suppression   = "disable"
  export_rt         = ["your own value"]
  fosid             = 10
  import_rt         = ["your own value"]
  ip_local_learning = "disable"
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `arp_suppression` - Enable/disable ARP suppression. Valid values: `disable`, `enable`.

* `export_rt` - List of export route targets.
* `fosid` - ID.
* `import_rt` - List of import route targets.
* `ip_local_learning` - Enable/disable IP address local learning. Valid values: `disable`, `enable`.

* `rd` - Route Distinguisher: AA:NN|A.B.C.D:NN.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System Evpn can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_evpn.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

