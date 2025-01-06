---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ospf_ospfinterface"
description: |-
  OSPF interface configuration.
---

# fmgdevice_router_ospf_ospfinterface
OSPF interface configuration.

~> This resource is a sub resource for variable `ospf_interface` of resource `fmgdevice_router_ospf`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `md5_keys`: `fmgdevice_router_ospf_ospfinterface_md5keys`



## Example Usage

```hcl
resource "fmgdevice_router_ospf_ospfinterface" "trname" {
  authentication     = "none"
  authentication_key = ["your own value"]
  bfd                = "global"
  comments           = "your own value"
  cost               = 10
  name               = "your own value"
  device_name        = var.device_name # not required if setting is at provider
  device_vdom        = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `authentication` - Authentication type. Valid values: `none`, `text`, `md5`, `message-digest`.

* `authentication_key` - Authentication key.
* `bfd` - Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.

* `comments` - Comment.
* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `database_filter_out` - Enable/disable control of flooding out LSAs. Valid values: `disable`, `enable`.

* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `hello_multiplier` - Number of hello packets within dead interval.
* `interface` - Configuration interface name.
* `ip` - IP address.
* `keychain` - Message-digest key-chain name.
* `linkdown_fast_failover` - Enable/disable fast link failover. Valid values: `disable`, `enable`.

* `md5_keychain` - Md5-Keychain.
* `md5_keys` - Md5-Keys. The structure of `md5_keys` block is documented below.
* `mtu` - MTU for database description packets.
* `mtu_ignore` - Enable/disable ignore MTU. Valid values: `disable`, `enable`.

* `name` - Interface entry name.
* `network_type` - Network type. Valid values: `broadcast`, `non-broadcast`, `point-to-point`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.

* `prefix_length` - Prefix length.
* `priority` - Priority.
* `resync_timeout` - Graceful restart neighbor resynchronization timeout.
* `retransmit_interval` - Retransmit interval.
* `status` - Enable/disable status. Valid values: `disable`, `enable`.

* `transmit_delay` - Transmit delay.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `md5_keys` block supports:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router OspfOspfInterface can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ospf_ospfinterface.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

