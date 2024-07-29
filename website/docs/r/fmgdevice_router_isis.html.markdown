---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_router_isis"
description: |-
  Configure IS-IS.
---

# fmgdevice_router_isis
Configure IS-IS.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `isis_interface`: `fmgdevice_router_isis_isisinterface`
>- `isis_net`: `fmgdevice_router_isis_isisnet`
>- `redistribute`: `fmgdevice_router_isis_redistribute`
>- `redistribute6`: `fmgdevice_router_isis_redistribute6`
>- `summary_address`: `fmgdevice_router_isis_summaryaddress`
>- `summary_address6`: `fmgdevice_router_isis_summaryaddress6`



## Example Usage

```hcl
resource "fmgdevice_router_isis" "trname" {
  adjacency_check   = "enable"
  adjacency_check6  = "disable"
  adv_passive_only  = "enable"
  adv_passive_only6 = "enable"
  auth_keychain_l1  = ["your own value"]
  device_name       = var.device_name # not required if setting is at provider
  device_vdom       = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `adjacency_check` - Enable/disable adjacency check. Valid values: `disable`, `enable`.

* `adjacency_check6` - Enable/disable IPv6 adjacency check. Valid values: `disable`, `enable`.

* `adv_passive_only` - Enable/disable IS-IS advertisement of passive interfaces only. Valid values: `disable`, `enable`.

* `adv_passive_only6` - Enable/disable IPv6 IS-IS advertisement of passive interfaces only. Valid values: `disable`, `enable`.

* `auth_keychain_l1` - Authentication key-chain for level 1 PDUs.
* `auth_keychain_l2` - Authentication key-chain for level 2 PDUs.
* `auth_mode_l1` - Level 1 authentication mode. Valid values: `md5`, `password`.

* `auth_mode_l2` - Level 2 authentication mode. Valid values: `md5`, `password`.

* `auth_password_l1` - Authentication password for level 1 PDUs.
* `auth_password_l2` - Authentication password for level 2 PDUs.
* `auth_sendonly_l1` - Enable/disable level 1 authentication send-only. Valid values: `disable`, `enable`.

* `auth_sendonly_l2` - Enable/disable level 2 authentication send-only. Valid values: `disable`, `enable`.

* `default_originate` - Enable/disable distribution of default route information. Valid values: `disable`, `enable`.

* `default_originate6` - Enable/disable distribution of default IPv6 route information. Valid values: `disable`, `enable`.

* `dynamic_hostname` - Enable/disable dynamic hostname. Valid values: `disable`, `enable`.

* `ignore_lsp_errors` - Enable/disable ignoring of LSP errors with bad checksums. Valid values: `disable`, `enable`.

* `is_type` - IS type. Valid values: `level-1-2`, `level-1`, `level-2-only`.

* `isis_interface` - Isis-Interface. The structure of `isis_interface` block is documented below.
* `isis_net` - Isis-Net. The structure of `isis_net` block is documented below.
* `lsp_gen_interval_l1` - Minimum interval for level 1 LSP regenerating.
* `lsp_gen_interval_l2` - Minimum interval for level 2 LSP regenerating.
* `lsp_refresh_interval` - LSP refresh time in seconds.
* `max_lsp_lifetime` - Maximum LSP lifetime in seconds.
* `metric_style` - Use old-style (ISO 10589) or new-style packet formats. Valid values: `narrow`, `narrow-transition`, `narrow-transition-l1`, `narrow-transition-l2`, `wide`, `wide-l1`, `wide-l2`, `wide-transition`, `wide-transition-l1`, `wide-transition-l2`, `transition`, `transition-l1`, `transition-l2`.

* `overload_bit` - Enable/disable signal other routers not to use us in SPF. Valid values: `disable`, `enable`.

* `overload_bit_on_startup` - Overload-bit only temporarily after reboot.
* `overload_bit_suppress` - Suppress overload-bit for the specific prefixes. Valid values: `external`, `interlevel`.

* `redistribute` - Redistribute. The structure of `redistribute` block is documented below.
* `redistribute_l1` - Enable/disable redistribution of level 1 routes into level 2. Valid values: `disable`, `enable`.

* `redistribute_l1_list` - Access-list for route redistribution from l1 to l2.
* `redistribute_l2` - Enable/disable redistribution of level 2 routes into level 1. Valid values: `disable`, `enable`.

* `redistribute_l2_list` - Access-list for route redistribution from l2 to l1.
* `redistribute6` - Redistribute6. The structure of `redistribute6` block is documented below.
* `redistribute6_l1` - Enable/disable redistribution of level 1 IPv6 routes into level 2. Valid values: `disable`, `enable`.

* `redistribute6_l1_list` - Access-list for IPv6 route redistribution from l1 to l2.
* `redistribute6_l2` - Enable/disable redistribution of level 2 IPv6 routes into level 1. Valid values: `disable`, `enable`.

* `redistribute6_l2_list` - Access-list for IPv6 route redistribution from l2 to l1.
* `spf_interval_exp_l1` - Level 1 SPF calculation delay.
* `spf_interval_exp_l2` - Level 2 SPF calculation delay.
* `summary_address` - Summary-Address. The structure of `summary_address` block is documented below.
* `summary_address6` - Summary-Address6. The structure of `summary_address6` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `isis_interface` block supports:

* `auth_keychain_l1` - Authentication key-chain for level 1 PDUs.
* `auth_keychain_l2` - Authentication key-chain for level 2 PDUs.
* `auth_mode_l1` - Level 1 authentication mode. Valid values: `md5`, `password`.

* `auth_mode_l2` - Level 2 authentication mode. Valid values: `md5`, `password`.

* `auth_password_l1` - Authentication password for level 1 PDUs.
* `auth_password_l2` - Authentication password for level 2 PDUs.
* `auth_send_only_l1` - Enable/disable authentication send-only for level 1 PDUs. Valid values: `disable`, `enable`.

* `auth_send_only_l2` - Enable/disable authentication send-only for level 2 PDUs. Valid values: `disable`, `enable`.

* `circuit_type` - IS-IS interface's circuit type. Valid values: `level-1-2`, `level-1`, `level-2`.

* `csnp_interval_l1` - Level 1 CSNP interval.
* `csnp_interval_l2` - Level 2 CSNP interval.
* `hello_interval_l1` - Level 1 hello interval.
* `hello_interval_l2` - Level 2 hello interval.
* `hello_multiplier_l1` - Level 1 multiplier for Hello holding time.
* `hello_multiplier_l2` - Level 2 multiplier for Hello holding time.
* `hello_padding` - Enable/disable padding to IS-IS hello packets. Valid values: `disable`, `enable`.

* `lsp_interval` - LSP transmission interval (milliseconds).
* `lsp_retransmit_interval` - LSP retransmission interval (sec).
* `mesh_group` - Enable/disable IS-IS mesh group. Valid values: `disable`, `enable`.

* `mesh_group_id` - Mesh group ID &lt;0-4294967295&gt;, 0: mesh-group blocked.
* `metric_l1` - Level 1 metric for interface.
* `metric_l2` - Level 2 metric for interface.
* `name` - IS-IS interface name.
* `network_type` - IS-IS interface's network type. Valid values: `broadcast`, `point-to-point`, `loopback`.

* `priority_l1` - Level 1 priority.
* `priority_l2` - Level 2 priority.
* `status` - Enable/disable interface for IS-IS. Valid values: `disable`, `enable`.

* `status6` - Enable/disable IPv6 interface for IS-IS. Valid values: `disable`, `enable`.

* `wide_metric_l1` - Level 1 wide metric for interface.
* `wide_metric_l2` - Level 2 wide metric for interface.

The `isis_net` block supports:

* `id` - ISIS network ID.
* `net` - IS-IS networks (format = xx.xxxx.  .xxxx.xx.).

The `redistribute` block supports:

* `level` - Level. Valid values: `level-1-2`, `level-1`, `level-2`.

* `metric` - Metric.
* `metric_type` - Metric type. Valid values: `external`, `internal`.

* `protocol` - Protocol name.
* `routemap` - Route map name.
* `status` - Status. Valid values: `disable`, `enable`.


The `redistribute6` block supports:

* `level` - Level. Valid values: `level-1-2`, `level-1`, `level-2`.

* `metric` - Metric.
* `metric_type` - Metric type. Valid values: `external`, `internal`.

* `protocol` - Protocol name.
* `routemap` - Route map name.
* `status` - Enable/disable redistribution. Valid values: `disable`, `enable`.


The `summary_address` block supports:

* `id` - Summary address entry ID.
* `level` - Level. Valid values: `level-1-2`, `level-1`, `level-2`.

* `prefix` - Prefix.

The `summary_address6` block supports:

* `id` - Prefix entry ID.
* `level` - Level. Valid values: `level-1-2`, `level-1`, `level-2`.

* `prefix6` - IPv6 prefix.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Router Isis can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_router_isis.labelname RouterIsis
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

