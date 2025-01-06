---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ospf"
description: |-
  Configure OSPF.
---

# fmgdevice_router_ospf
Configure OSPF.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `area`: `fmgdevice_router_ospf_area`
>- `distribute_list`: `fmgdevice_router_ospf_distributelist`
>- `neighbor`: `fmgdevice_router_ospf_neighbor`
>- `network`: `fmgdevice_router_ospf_network`
>- `ospf_interface`: `fmgdevice_router_ospf_ospfinterface`
>- `redistribute`: `fmgdevice_router_ospf_redistribute`
>- `summary_address`: `fmgdevice_router_ospf_summaryaddress`



## Example Usage

```hcl
resource "fmgdevice_router_ospf" "trname" {
  abr_type = "standard"
  area {
    authentication = "message-digest"
    comments       = "your own value"
    default_cost   = 10
    filter_list {
      direction = "in"
      fosid     = 10
      list      = ["your own value"]
    }

    fosid                                          = "your own value"
    nssa_default_information_originate             = "enable"
    nssa_default_information_originate_metric      = 10
    nssa_default_information_originate_metric_type = "1"
    nssa_redistribution                            = "enable"
    nssa_translator_role                           = "candidate"
    range {
      advertise         = "enable"
      fosid             = 10
      prefix            = ["your own value"]
      substitute        = ["your own value"]
      substitute_status = "disable"
    }

    shortcut  = "default"
    stub_type = "no-summary"
    type      = "nssa"
    virtual_link {
      authentication     = "none"
      authentication_key = ["your own value"]
      dead_interval      = 10
      hello_interval     = 10
      keychain           = ["your own value"]
      md5_keychain       = ["your own value"]
      md5_keys {
        fosid      = 10
        key_string = ["your own value"]
      }

      name                = "your own value"
      peer                = "your own value"
      retransmit_interval = 10
      transmit_delay      = 10
    }

  }

  auto_cost_ref_bandwidth = 10
  bfd                     = "enable"
  database_overflow       = "disable"
  device_name             = var.device_name # not required if setting is at provider
  device_vdom             = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `abr_type` - Area border router type. Valid values: `cisco`, `ibm`, `shortcut`, `standard`.

* `area` - Area. The structure of `area` block is documented below.
* `auto_cost_ref_bandwidth` - Reference bandwidth in terms of megabits per second.
* `bfd` - Bidirectional Forwarding Detection (BFD). Valid values: `disable`, `enable`.

* `database_overflow` - Enable/disable database overflow. Valid values: `disable`, `enable`.

* `database_overflow_max_lsas` - Database overflow maximum LSAs.
* `database_overflow_time_to_recover` - Database overflow time to recover (sec).
* `default_information_metric` - Default information metric.
* `default_information_metric_type` - Default information metric type. Valid values: `2`, `1`.

* `default_information_originate` - Enable/disable generation of default route. Valid values: `disable`, `enable`, `always`.

* `default_information_route_map` - Default information route map.
* `default_metric` - Default metric of redistribute routes.
* `distance` - Distance of the route.
* `distance_external` - Administrative external distance.
* `distance_inter_area` - Administrative inter-area distance.
* `distance_intra_area` - Administrative intra-area distance.
* `distribute_list` - Distribute-List. The structure of `distribute_list` block is documented below.
* `distribute_list_in` - Filter incoming routes.
* `distribute_route_map_in` - Filter incoming external routes by route-map.
* `log_neighbour_changes` - Log of OSPF neighbor changes. Valid values: `disable`, `enable`.

* `lsa_refresh_interval` - The minimal OSPF LSA update time interval
* `neighbor` - Neighbor. The structure of `neighbor` block is documented below.
* `network` - Network. The structure of `network` block is documented below.
* `ospf_interface` - Ospf-Interface. The structure of `ospf_interface` block is documented below.
* `passive_interface` - Passive interface configuration.
* `redistribute` - Redistribute. The structure of `redistribute` block is documented below.
* `restart_mode` - OSPF restart mode (graceful or LLS). Valid values: `none`, `lls`, `graceful-restart`.

* `restart_on_topology_change` - Enable/disable continuing graceful restart upon topology change. Valid values: `disable`, `enable`.

* `restart_period` - Graceful restart period.
* `rfc1583_compatible` - Enable/disable RFC1583 compatibility. Valid values: `disable`, `enable`.

* `router_id` - Router ID.
* `spf_timers` - SPF calculation frequency.
* `summary_address` - Summary-Address. The structure of `summary_address` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `area` block supports:

* `authentication` - Authentication type. Valid values: `none`, `text`, `md5`, `message-digest`.

* `comments` - Comment.
* `default_cost` - Summary default cost of stub or NSSA area.
* `filter_list` - Filter-List. The structure of `filter_list` block is documented below.
* `id` - Area entry IP address.
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

The `distribute_list` block supports:

* `access_list` - Access list name.
* `id` - Distribute list entry ID.
* `protocol` - Protocol type. Valid values: `connected`, `static`, `rip`.


The `neighbor` block supports:

* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `id` - Neighbor entry ID.
* `ip` - Interface IP address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `priority` - Priority.

The `network` block supports:

* `area` - Attach the network to area.
* `comments` - Comment.
* `id` - Network entry ID.
* `prefix` - Prefix.

The `ospf_interface` block supports:

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

The `md5_keys` block supports:

* `id` - Key ID (1 - 255).
* `key_string` - Password for the key.

The `redistribute` block supports:

* `metric` - Redistribute metric setting.
* `metric_type` - Metric type. Valid values: `2`, `1`.

* `name` - Redistribute name.
* `routemap` - Route map name.
* `status` - Status. Valid values: `disable`, `enable`.

* `tag` - Tag value.

The `summary_address` block supports:

* `advertise` - Enable/disable advertise status. Valid values: `disable`, `enable`.

* `id` - Summary address entry ID.
* `prefix` - Prefix.
* `tag` - Tag value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Ospf can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ospf.labelname RouterOspf
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

