---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ipam_pools_exclude"
description: |-
  Configure pool exclude subnets.
---

# fmgdevice_system_ipam_pools_exclude
Configure pool exclude subnets.

~> This resource is a sub resource for variable `exclude` of resource `fmgdevice_system_ipam_pools`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_ipam_pools_exclude" "trname" {
  fosid          = 10
  exclude_subnet = ["your own value"]
  device_name    = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `pools` - Pools.

* `fosid` - Exclude ID.
* `exclude_subnet` - Configure subnet to exclude from the IPAM pool.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System IpamPoolsExclude can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "pools=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ipam_pools_exclude.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

