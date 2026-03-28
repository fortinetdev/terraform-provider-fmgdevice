---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ipsec_fec_mappings"
description: |-
  FEC redundancy mapping table.
---

# fmgdevice_vpn_ipsec_fec_mappings
FEC redundancy mapping table.

~> This resource is a sub resource for variable `mappings` of resource `fmgdevice_vpn_ipsec_fec`. Conflict and overwrite may occur if use both of them.



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `fec` - Fec.

* `bandwidth_bi_threshold` - Apply FEC parameters when available bi-bandwidth is &gt;= threshold (kbps, 0 means no threshold).
* `bandwidth_down_threshold` - Apply FEC parameters when available down bandwidth is &gt;= threshold (kbps, 0 means no threshold).
* `bandwidth_up_threshold` - Apply FEC parameters when available up bandwidth is &gt;= threshold (kbps, 0 means no threshold).
* `base` - Number of base FEC packets (1 - 20).
* `latency_threshold` - Apply FEC parameters when latency is &lt;= threshold (0 means no threshold).
* `packet_loss_threshold` - Apply FEC parameters when packet loss is &gt;= threshold (0 - 100, 0 means no threshold).
* `redundant` - Number of redundant FEC packets (1 - 5).
* `seqno` - Sequence number (1 - 64).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seqno}}.

## Import

Vpn IpsecFecMappings can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "fec=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ipsec_fec_mappings.labelname {{seqno}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

