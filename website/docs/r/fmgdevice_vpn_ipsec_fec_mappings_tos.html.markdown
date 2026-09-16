---
subcategory: "VPN"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ipsec_fec_mappings_tos"
description: |-
  FEC redundancy mapping table for specific type of service (TOS).
---

# fmgdevice_vpn_ipsec_fec_mappings_tos
FEC redundancy mapping table for specific type of service (TOS).

~> This resource is a sub resource for variable `tos` of resource `fmgdevice_vpn_ipsec_fec_mappings`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `fec` - Fec.
* `mappings` - Mappings.

* `base` - Number of base FEC packets (1 - 40).
* `redundant` - Number of redundant FEC packets (0 - 20).
* `seqno` - Sequence number (1 - 8).
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seqno}}.

## Import

Vpn IpsecFecMappingsTos can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "fec=YOUR_VALUE", "mappings=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ipsec_fec_mappings_tos.labelname {{seqno}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

