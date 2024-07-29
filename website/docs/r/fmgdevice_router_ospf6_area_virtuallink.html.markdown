---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ospf6_area_virtuallink"
description: |-
  OSPF6 virtual link configuration.
---

# fmgdevice_router_ospf6_area_virtuallink
OSPF6 virtual link configuration.

~> This resource is a sub resource for variable `virtual_link` of resource `fmgdevice_router_ospf6_area`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `ipsec_keys`: `fmgdevice_router_ospf6_area_virtuallink_ipseckeys`



## Example Usage

```hcl
resource "fmgdevice_router_ospf6_area_virtuallink" "trname" {
  authentication = "area"
  dead_interval  = 10
  hello_interval = 10
  ipsec_auth_alg = "sha384"
  ipsec_enc_alg  = "aes192"
  name           = "your own value"
  device_name    = var.device_name # not required if setting is at provider
  device_vdom    = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `area` - Area.

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
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `ipsec_keys` block supports:

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router Ospf6AreaVirtualLink can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "area=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ospf6_area_virtuallink.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

