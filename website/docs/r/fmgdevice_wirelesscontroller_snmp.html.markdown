---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_wirelesscontroller_snmp"
description: |-
  Configure SNMP.
---

# fmgdevice_wirelesscontroller_snmp
Configure SNMP.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `community`: `fmgdevice_wirelesscontroller_snmp_community`
>- `user`: `fmgdevice_wirelesscontroller_snmp_user`



## Example Usage

```hcl
resource "fmgdevice_wirelesscontroller_snmp" "trname" {
  community {
    hosts {
      fosid = 10
      ip    = "your own value"
    }

    fosid            = 10
    name             = "your own value"
    query_v1_status  = "enable"
    query_v2c_status = "enable"
    status           = "enable"
    trap_v1_status   = "enable"
    trap_v2c_status  = "disable"
  }

  contact_info            = "your own value"
  engine_id               = "your own value"
  trap_high_cpu_threshold = 10
  trap_high_mem_threshold = 10
  device_name             = var.device_name # not required if setting is at provider
  device_vdom             = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `community` - Community. The structure of `community` block is documented below.
* `contact_info` - Contact Information.
* `engine_id` - AC SNMP engineID string (maximum 24 characters).
* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_high_mem_threshold` - Memory usage when trap is sent.
* `user` - User. The structure of `user` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `community` block supports:

* `hosts` - Hosts. The structure of `hosts` block is documented below.
* `id` - Community ID.
* `name` - Community name.
* `query_v1_status` - Enable/disable SNMP v1 queries. Valid values: `disable`, `enable`.

* `query_v2c_status` - Enable/disable SNMP v2c queries. Valid values: `disable`, `enable`.

* `status` - Enable/disable this SNMP community. Valid values: `disable`, `enable`.

* `trap_v1_status` - Enable/disable SNMP v1 traps. Valid values: `disable`, `enable`.

* `trap_v2c_status` - Enable/disable SNMP v2c traps. Valid values: `disable`, `enable`.


The `hosts` block supports:

* `id` - Host entry ID.
* `ip` - IPv4 address of the SNMP manager (host).

The `user` block supports:

* `auth_proto` - Authentication protocol. Valid values: `md5`, `sha`.

* `auth_pwd` - Password for authentication protocol.
* `name` - SNMP user name.
* `notify_hosts` - Configure SNMP User Notify Hosts.
* `priv_proto` - Privacy (encryption) protocol. Valid values: `aes`, `des`, `aes256`, `aes256cisco`.

* `priv_pwd` - Password for privacy (encryption) protocol.
* `queries` - Enable/disable SNMP queries for this user. Valid values: `disable`, `enable`.

* `security_level` - Security level for message authentication and encryption. Valid values: `no-auth-no-priv`, `auth-no-priv`, `auth-priv`.

* `status` - SNMP user enable. Valid values: `disable`, `enable`.

* `trap_status` - Enable/disable traps for this SNMP user. Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController Snmp can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_wirelesscontroller_snmp.labelname WirelessControllerSnmp
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

