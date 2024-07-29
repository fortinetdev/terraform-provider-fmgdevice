---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_ospf6"
description: |-
  Configure IPv6 OSPF.
---

# fmgdevice_router_ospf6
Configure IPv6 OSPF.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `area`: `fmgdevice_router_ospf6_area`
>- `ospf6_interface`: `fmgdevice_router_ospf6_ospf6interface`
>- `redistribute`: `fmgdevice_router_ospf6_redistribute`
>- `summary_address`: `fmgdevice_router_ospf6_summaryaddress`



## Example Usage

```hcl
resource "fmgdevice_router_ospf6" "trname" {
  abr_type = "standard"
  area {
    authentication = "esp"
    default_cost   = 10
    fosid          = "your own value"
    ipsec_auth_alg = "sha512"
    ipsec_enc_alg  = "aes256"
    ipsec_keys {
      auth_key = ["your own value"]
      enc_key  = ["your own value"]
      spi      = 10
    }

    key_rollover_interval                          = 10
    nssa_default_information_originate             = "disable"
    nssa_default_information_originate_metric      = 10
    nssa_default_information_originate_metric_type = "1"
    nssa_redistribution                            = "disable"
    nssa_translator_role                           = "always"
    range {
      advertise = "disable"
      fosid     = 10
      prefix6   = "your own value"
    }

    stub_type = "summary"
    type      = "stub"
    virtual_link {
      authentication = "ah"
      dead_interval  = 10
      hello_interval = 10
      ipsec_auth_alg = "md5"
      ipsec_enc_alg  = "des"
      ipsec_keys {
        auth_key = ["your own value"]
        enc_key  = ["your own value"]
        spi      = 10
      }

      key_rollover_interval = 10
      name                  = "your own value"
      peer                  = "your own value"
      retransmit_interval   = 10
      transmit_delay        = 10
    }

  }

  auto_cost_ref_bandwidth    = 10
  bfd                        = "enable"
  default_information_metric = 10
  device_name                = var.device_name # not required if setting is at provider
  device_vdom                = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `abr_type` - Area border router type. Valid values: `cisco`, `ibm`, `standard`.

* `area` - Area. The structure of `area` block is documented below.
* `auto_cost_ref_bandwidth` - Reference bandwidth in terms of megabits per second.
* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `disable`, `enable`.

* `default_information_metric` - Default information metric.
* `default_information_metric_type` - Default information metric type. Valid values: `1`, `2`.

* `default_information_originate` - Enable/disable generation of default route. Valid values: `disable`, `enable`, `always`.

* `default_information_route_map` - Default information route map.
* `default_metric` - Default metric of redistribute routes.
* `log_neighbour_changes` - Log OSPFv3 neighbor changes. Valid values: `disable`, `enable`.

* `ospf6_interface` - Ospf6-Interface. The structure of `ospf6_interface` block is documented below.
* `passive_interface` - Passive interface configuration.
* `redistribute` - Redistribute. The structure of `redistribute` block is documented below.
* `restart_mode` - OSPFv3 restart mode (graceful or none). Valid values: `none`, `graceful-restart`.

* `restart_on_topology_change` - Enable/disable continuing graceful restart upon topology change. Valid values: `disable`, `enable`.

* `restart_period` - Graceful restart period in seconds.
* `router_id` - A.B.C.D, in IPv4 address format.
* `spf_timers` - SPF calculation frequency.
* `summary_address` - Summary-Address. The structure of `summary_address` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `area` block supports:

* `authentication` - Authentication mode. Valid values: `none`, `ah`, `esp`.

* `default_cost` - Summary default cost of stub or NSSA area.
* `id` - Area entry IP address.
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

The `ospf6_interface` block supports:

* `area_id` - A.B.C.D, in IPv4 address format.
* `authentication` - Authentication mode. Valid values: `none`, `ah`, `esp`, `area`.

* `bfd` - Enable/disable Bidirectional Forwarding Detection (BFD). Valid values: `global`, `enable`, `disable`.

* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `dead_interval` - Dead interval.
* `hello_interval` - Hello interval.
* `interface` - Configuration interface name.
* `ipsec_auth_alg` - Authentication algorithm. Valid values: `md5`, `sha1`, `sha256`, `sha384`, `sha512`.

* `ipsec_enc_alg` - Encryption algorithm. Valid values: `null`, `des`, `3des`, `aes128`, `aes192`, `aes256`.

* `ipsec_keys` - Ipsec-Keys. The structure of `ipsec_keys` block is documented below.
* `key_rollover_interval` - Key roll-over interval.
* `mtu` - MTU for OSPFv3 packets.
* `mtu_ignore` - Enable/disable ignoring MTU field in DBD packets. Valid values: `disable`, `enable`.

* `name` - Interface entry name.
* `neighbor` - Neighbor. The structure of `neighbor` block is documented below.
* `network_type` - Network type. Valid values: `broadcast`, `non-broadcast`, `point-to-point`, `point-to-multipoint`, `point-to-multipoint-non-broadcast`.

* `priority` - Priority.
* `retransmit_interval` - Retransmit interval.
* `status` - Enable/disable OSPF6 routing on this interface. Valid values: `disable`, `enable`.

* `transmit_delay` - Transmit delay.

The `ipsec_keys` block supports:

* `auth_key` - Authentication key.
* `enc_key` - Encryption key.
* `spi` - Security Parameters Index.

The `neighbor` block supports:

* `cost` - Cost of the interface, value range from 0 to 65535, 0 means auto-cost.
* `ip6` - IPv6 link local address of the neighbor.
* `poll_interval` - Poll interval time in seconds.
* `priority` - Priority.

The `redistribute` block supports:

* `metric` - Redistribute metric setting.
* `metric_type` - Metric type. Valid values: `1`, `2`.

* `name` - Redistribute name.
* `routemap` - Route map name.
* `status` - Status. Valid values: `disable`, `enable`.


The `summary_address` block supports:

* `advertise` - Enable/disable advertise status. Valid values: `disable`, `enable`.

* `id` - Summary address entry ID.
* `prefix6` - IPv6 prefix.
* `tag` - Tag value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Ospf6 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_ospf6.labelname RouterOspf6
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

