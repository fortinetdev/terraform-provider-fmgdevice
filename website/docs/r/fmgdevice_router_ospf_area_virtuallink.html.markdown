---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ospf_area_virtuallink"
description: |-
  OSPF virtual link configuration.
---

# fmgdevice_router_ospf_area_virtuallink
OSPF virtual link configuration.

~> This resource is a sub resource for variable `virtual_link` of resource `fmgdevice_router_ospf_area`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `md5_keys`: `fmgdevice_router_ospf_area_virtuallink_md5keys`



## Example Usage

```hcl
resource "fmgdevice_router_ospf_area_virtuallink" "trname" {
  authentication     = "md5"
  authentication_key = ["your own value"]
  dead_interval      = 10
  hello_interval     = 10
  keychain           = ["your own value"]
  name               = "your own value"
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.
* `area` - Area.

* `authentication` - Authentication type. Valid values: `none`, `text`, `md5`, `message-digest`.

* `authentication_key` - Authentication key.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `keychain` - Message-digest key-chain name.
* `md5_keychain` - Md5-Keychain.
* `md5_keys` - Md5-Keys. The structure of `md5_keys` block is documented below.
* `name` - Virtual link entry name.
* `peer` - Peer IP.
* `retransmit_interval` - Retransmit interval.
* `transmit_delay` - Transmit delay.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `md5_keys` block supports:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router OspfAreaVirtualLink can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE", "area=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ospf_area_virtuallink.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

