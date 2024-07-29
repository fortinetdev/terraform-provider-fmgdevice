---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_multicast_pimsmglobal_rpaddress"
description: |-
  Statically configure RP addresses.
---

# fmgdevice_router_multicast_pimsmglobal_rpaddress
Statically configure RP addresses.

~> This resource is a sub resource for variable `rp_address` of resource `fmgdevice_router_multicast_pimsmglobal`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_multicast_pimsmglobal_rpaddress" "trname" {
  group       = ["your own value"]
  fosid       = 10
  ip_address  = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `group` - Groups to use this RP.
* `fosid` - ID.
* `ip_address` - RP router address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router MulticastPimSmGlobalRpAddress can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_multicast_pimsmglobal_rpaddress.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

