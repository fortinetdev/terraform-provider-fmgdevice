---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ipam_pools"
description: |-
  Configure IPAM pools.
---

# fmgdevice_system_ipam_pools
Configure IPAM pools.

~> This resource is a sub resource for variable `pools` of resource `fmgdevice_system_ipam`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `exclude`: `fmgdevice_system_ipam_pools_exclude`



## Example Usage

```hcl
resource "fmgdevice_system_ipam_pools" "trname" {
  description = "your own value"
  exclude {
    fosid          = 10
    exclude_subnet = ["your own value"]
  }

  name        = "your own value"
  subnet      = ["your own value"]
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `description` - Description.
* `exclude` - Exclude. The structure of `exclude` block is documented below.
* `name` - IPAM pool name.
* `subnet` - Configure IPAM pool subnet, Class A - Class B subnet.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `exclude` block supports:

* `id` - Exclude ID.
* `exclude_subnet` - Configure subnet to exclude from the IPAM pool.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System IpamPools can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ipam_pools.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

