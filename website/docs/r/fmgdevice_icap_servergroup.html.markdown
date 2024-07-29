---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_icap_servergroup"
description: |-
  Configure an ICAP server group consisting of multiple forward servers. Supports failover and load balancing.
---

# fmgdevice_icap_servergroup
Configure an ICAP server group consisting of multiple forward servers. Supports failover and load balancing.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `server_list`: `fmgdevice_icap_servergroup_serverlist`



## Example Usage

```hcl
resource "fmgdevice_icap_servergroup" "trname" {
  ldb_method = "weighted"
  name       = "your own value"
  server_list {
    name   = ["your own value"]
    weight = 10
  }

  device_name = var.device_name # not required if setting is at provider
  device_vdom = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `ldb_method` - Load balance method. Valid values: `weighted`, `least-session`, `active-passive`.

* `name` - Configure an ICAP server group consisting one or multiple forward servers. Supports failover and load balancing.
* `server_list` - Server-List. The structure of `server_list` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `server_list` block supports:

* `name` - ICAP server name.
* `weight` - Optionally assign a weight of the forwarding server for weighted load balancing (1 - 100, default = 10).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Icap ServerGroup can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_icap_servergroup.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

