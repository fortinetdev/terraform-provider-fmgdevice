---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ipv6neighborcache"
description: |-
  Configure IPv6 neighbor cache table.
---

# fmgdevice_system_ipv6neighborcache
Configure IPv6 neighbor cache table.

## Example Usage

```hcl
resource "fmgdevice_system_ipv6neighborcache" "trname" {
  fosid       = 10
  interface   = ["port2"]
  ipv6        = "your own value"
  mac         = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `fosid` - Unique integer ID of the entry.
* `interface` - Select the associated interface name from available options.
* `ipv6` - IPv6 address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `mac` - MAC address (format: xx:xx:xx:xx:xx:xx).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System Ipv6NeighborCache can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ipv6neighborcache.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

