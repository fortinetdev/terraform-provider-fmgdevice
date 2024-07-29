---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ospf_area"
description: |-
  OSPF area configuration.
---

# fmgdevice_router_ospf_area
OSPF area configuration.

~> This resource is a sub resource for variable `area` of resource `fmgdevice_router_ospf`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `filter_list`: `fmgdevice_router_ospf_area_filterlist`
>- `range`: `fmgdevice_router_ospf_area_range`
>- `virtual_link`: `fmgdevice_router_ospf_area_virtuallink`



## Example Usage

```hcl
resource "fmgdevice_router_ospf_area" "trname" {
  authentication = "none"
  comments       = "your own value"
  default_cost   = 10
  filter_list {
    direction = "out"
    fosid     = 10
    list      = ["your own value"]
  }

  fosid       = "your own value"
  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `authentication` - Authentication type. Valid values: `none`, `text`, `md5`, `message-digest`.

* `comments` - Comment.
* `default_cost` - Summary default cost of stub or NSSA area.
* `filter_list` - Filter-List. The structure of `filter_list` block is documented below.
* `fosid` - Area entry IP address.
* `nssa_default_information_originate` - Redistribute, advertise, or do not originate Type-7 default route into NSSA area. Valid values: `disable`, `enable`, `always`.

* `nssa_default_information_originate_metric` - OSPF default metric.
* `nssa_default_information_originate_metric_type` - OSPF metric type for default routes. Valid values: `2`, `1`.

* `nssa_redistribution` - Enable/disable redistribute into NSSA area. Valid values: `disable`, `enable`.

* `nssa_translator_role` - NSSA translator role type. Valid values: `candidate`, `never`, `always`.

* `range` - Range. The structure of `range` block is documented below.
* `shortcut` - Enable/disable shortcut option. Valid values: `disable`, `enable`, `default`.

* `stub_type` - Stub summary setting. Valid values: `summary`, `no-summary`.

* `type` - Area type setting. Valid values: `regular`, `nssa`, `stub`.

* `virtual_link` - Virtual-Link. The structure of `virtual_link` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `filter_list` block supports:

* `direction` - Direction. Valid values: `out`, `in`.

* `id` - Filter list entry ID.
* `list` - Access-list or prefix-list name.

The `range` block supports:

* `advertise` - Enable/disable advertise status. Valid values: `disable`, `enable`.

* `id` - Range entry ID.
* `prefix` - Prefix.
* `substitute` - Substitute prefix.
* `substitute_status` - Enable/disable substitute status. Valid values: `disable`, `enable`.


The `virtual_link` block supports:

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

The `md5_keys` block supports:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Router OspfArea can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ospf_area.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

