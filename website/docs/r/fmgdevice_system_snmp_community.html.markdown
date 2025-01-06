---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_snmp_community"
description: |-
  SNMP community configuration.
---

# fmgdevice_system_snmp_community
SNMP community configuration.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `hosts`: `fmgdevice_system_snmp_community_hosts`
>- `hosts6`: `fmgdevice_system_snmp_community_hosts6`



## Example Usage

```hcl
resource "fmgdevice_system_snmp_community" "trname" {
  events = ["intf-ip"]
  hosts {
    ha_direct = "disable"
    host_type = "trap"
    fosid     = 10
    ip        = ["your own value"]
    source_ip = "your own value"
  }

  hosts6 {
    ha_direct   = "disable"
    host_type   = "trap"
    fosid       = 10
    ipv6        = "your own value"
    source_ipv6 = "your own value"
  }

  fosid       = 10
  mib_view    = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `events` - SNMP trap events. Valid values: `cpu-high`, `mem-low`, `log-full`, `intf-ip`, `vpn-tun-up`, `vpn-tun-down`, `ha-switch`, `ha-hb-failure`, `ips-signature`, `ips-anomaly`, `av-virus`, `av-oversize`, `av-pattern`, `av-fragmented`, `fm-if-change`, `fm-conf-change`, `temperature-high`, `voltage-alert`, `ha-member-up`, `ha-member-down`, `ent-conf-change`, `av-conserve`, `av-bypass`, `av-oversize-passed`, `av-oversize-blocked`, `ips-pkg-update`, `power-supply-failure`, `amc-bypass`, `faz-disconnect`, `fan-failure`, `bgp-established`, `bgp-backward-transition`, `wc-ap-up`, `wc-ap-down`, `fswctl-session-up`, `fswctl-session-down`, `ips-fail-open`, `load-balance-real-server-down`, `device-new`, `enter-intf-bypass`, `exit-intf-bypass`, `per-cpu-high`, `power-blade-down`, `confsync_failure`, `dhcp`, `pool-usage`, `power-redundancy-degrade`, `power-redundancy-failure`, `ospf-nbr-state-change`, `ospf-virtnbr-state-change`, `disk-failure`, `disk-overload`, `faz-main-failover`, `faz-alt-failover`, `slbc`, `faz`, `power-supply`.

* `hosts` - Hosts. The structure of `hosts` block is documented below.
* `hosts6` - Hosts6. The structure of `hosts6` block is documented below.
* `fosid` - Community ID.
* `mib_view` - SNMP access control MIB view.
* `name` - Community name.
* `query_v1_port` - SNMP v1 query port (default = 161).
* `query_v1_status` - Enable/disable SNMP v1 queries. Valid values: `disable`, `enable`.

* `query_v2c_port` - SNMP v2c query port (default = 161).
* `query_v2c_status` - Enable/disable SNMP v2c queries. Valid values: `disable`, `enable`.

* `status` - Enable/disable this SNMP community. Valid values: `disable`, `enable`.

* `trap_v1_lport` - SNMP v1 trap local port (default = 162).
* `trap_v1_rport` - SNMP v1 trap remote port (default = 162).
* `trap_v1_status` - Enable/disable SNMP v1 traps. Valid values: `disable`, `enable`.

* `trap_v2c_lport` - SNMP v2c trap local port (default = 162).
* `trap_v2c_rport` - SNMP v2c trap remote port (default = 162).
* `trap_v2c_status` - Enable/disable SNMP v2c traps. Valid values: `disable`, `enable`.

* `vdoms` - SNMP access control VDOMs.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `hosts` block supports:

* `ha_direct` - Enable/disable direct management of HA cluster members. Valid values: `disable`, `enable`.

* `host_type` - Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. No traps will be sent when IP type is subnet. Valid values: `any`, `query`, `trap`.

* `id` - Host entry ID.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ip` - IPv4 address of the SNMP manager (host).
* `source_ip` - Source IPv4 address for SNMP traps.

The `hosts6` block supports:

* `ha_direct` - Enable/disable direct management of HA cluster members. Valid values: `disable`, `enable`.

* `host_type` - Control whether the SNMP manager sends SNMP queries, receives SNMP traps, or both. Valid values: `any`, `query`, `trap`.

* `id` - Host6 entry ID.
* `interface` - Specify outgoing interface to reach server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.

* `ipv6` - SNMP manager IPv6 address prefix.
* `source_ipv6` - Source IPv6 address for SNMP traps.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SnmpCommunity can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_snmp_community.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

