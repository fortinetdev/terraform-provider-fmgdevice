---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_nsxt_servicechain"
description: |-
  Configure NSX-T service chain.
---

# fmgdevice_nsxt_servicechain
Configure NSX-T service chain.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `service_index`: `fmgdevice_nsxt_servicechain_serviceindex`



## Example Usage

```hcl
resource "fmgdevice_nsxt_servicechain" "trname" {
  fosid = 10
  name  = "your own value"
  service_index {
    fosid         = 10
    name          = "your own value"
    reverse_index = 10
    vd            = ["your own value"]
  }

  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fosid` - Chain ID.
* `name` - Chain name.
* `service_index` - Service-Index. The structure of `service_index` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `service_index` block supports:

* `id` - Service index.
* `name` - Index name.
* `reverse_index` - Reverse service index.
* `vd` - VDOM name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Nsxt ServiceChain can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_nsxt_servicechain.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

