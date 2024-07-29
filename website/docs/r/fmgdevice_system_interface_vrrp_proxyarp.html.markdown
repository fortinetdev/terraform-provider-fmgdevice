---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_interface_vrrp_proxyarp"
description: |-
  VRRP Proxy ARP configuration.
---

# fmgdevice_system_interface_vrrp_proxyarp
VRRP Proxy ARP configuration.

~> This resource is a sub resource for variable `proxy_arp` of resource `fmgdevice_system_interface_vrrp`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_interface_vrrp_proxyarp" "trname" {
  fosid       = 10
  ip          = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `interface` - Interface.
* `vrrp` - Vrrp.

* `fosid` - ID.
* `ip` - Set IP addresses of proxy ARP.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System InterfaceVrrpProxyArp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "interface=YOUR_VALUE", "vrrp=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_interface_vrrp_proxyarp.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

