---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ospf6_area"
description: |-
  OSPF6 area configuration.
---

# fmgdevice_router_ospf6_area
OSPF6 area configuration.

~> This resource is a sub resource for variable `area` of resource `fmgdevice_router_ospf6`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `ipsec_keys`: `fmgdevice_router_ospf6_area_ipseckeys`
>- `range`: `fmgdevice_router_ospf6_area_range`
>- `virtual_link`: `fmgdevice_router_ospf6_area_virtuallink`



## Example Usage

```hcl
resource "fmgdevice_router_ospf6_area" "trname" {
  authentication = "none"
  default_cost   = 10
  fosid          = "your own value"
  ipsec_auth_alg = "sha256"
  ipsec_enc_alg  = "aes256"
  device_name    = var.device_name # not required if setting is at provider
  device_vdom    = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `authentication` - Authentication mode. Valid values: `none`, `ah`, `esp`.

* `default_cost` - Summary default cost of stub or NSSA area.
* `fosid` - Area entry IP address.
* `ipsec_auth_alg` - Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.

* `ipsec_enc_alg` - Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.

* `ipsec_keys` - Ipsec-Keys. The structure of `ipsec_keys` block is documented below.
* `key_rollover_interval` - Key roll-over interval.
* `nssa_default_information_originate` - Enable/disable originate type 7 default into NSSA area. Valid values: `disable`, `enable`.

* `nssa_default_information_originate_metric` - OSPFv3 default metric.
* `nssa_default_information_originate_metric_type` - OSPFv3 metric type for default routes. Valid values: `1`, `2`.

* `nssa_redistribution` - Enable/disable redistribute into NSSA area. Valid values: `disable`, `enable`.

* `nssa_translator_role` - NSSA translator role type. Valid values: `candidate`, `never`, `always`.

* `range` - Range. The structure of `range` block is documented below.
* `stub_type` - Stub summary setting. Valid values: `summary`, `no-summary`.

* `type` - Area type setting. Valid values: `regular`, `nssa`, `stub`.

* `virtual_link` - Virtual-Link. The structure of `virtual_link` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ipsec_keys` block supports:

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.

The `range` block supports:

* `advertise` - Enable/disable advertise status. Valid values: `disable`, `enable`.

* `id` - Range entry ID.
* `prefix6` - IPv6 prefix.

The `virtual_link` block supports:

* `authentication` - Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.

* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `ipsec_auth_alg` - Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.

* `ipsec_enc_alg` - Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.

* `ipsec_keys` - Ipsec-Keys. The structure of `ipsec_keys` block is documented below.
* `key_rollover_interval` - Key roll-over interval.
* `name` - Virtual link entry name.
* `peer` - A.B.C.D, peer router ID.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.

The `ipsec_keys` block supports:

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router Ospf6Area can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ospf6_area.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

