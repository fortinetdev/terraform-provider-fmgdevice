---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ospf6_area_ipseckeys"
description: |-
  IPsec authentication and encryption keys.
---

# fmgdevice_router_ospf6_area_ipseckeys
IPsec authentication and encryption keys.

~> This resource is a sub resource for variable `ipsec_keys` of resource `fmgdevice_router_ospf6_area`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_router_ospf6_area_ipseckeys" "trname" {
  auth_key    = ["your own value"]
  enc_key     = ["your own value"]
  spi         = 10
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `area` - Area.

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{spi}}.

## Import

Router Ospf6AreaIpsecKeys can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "area=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ospf6_area_ipseckeys.labelname {{spi}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

