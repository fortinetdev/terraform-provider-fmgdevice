---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_vpn_ipsec_fec"
description: |-
  Configure Forward Error Correction (FEC) mapping profiles.
---

# fmgdevice_vpn_ipsec_fec
Configure Forward Error Correction (FEC) mapping profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `mappings`: `fmgdevice_vpn_ipsec_fec_mappings`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `mappings` - Mappings. The structure of `mappings` block is documented below.
* `name` - Profile name.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `mappings` block supports:

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
* `id` - an identifier for the resource with format {{name}}.

## Import

Vpn IpsecFec can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_vpn_ipsec_fec.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

