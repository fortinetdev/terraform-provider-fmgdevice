---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_internetservicedefinition"
description: |-
  Configure Internet Service definition.
---

# fmgdevice_firewall_internetservicedefinition
Configure Internet Service definition.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entry`: `fmgdevice_firewall_internetservicedefinition_entry`



## Example Usage

```hcl
resource "fmgdevice_firewall_internetservicedefinition" "trname" {
  entry {
    category_id = 10
    name        = "your own value"
    port_range {
      end_port   = 10
      fosid      = 10
      start_port = 10
    }

    protocol = 10
    seq_num  = 10
  }

  fosid       = 10
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `entry` - Entry. The structure of `entry` block is documented below.
* `fosid` - Id.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entry` block supports:

* `category_id` - Category-Id.
* `name` - Name.
* `port_range` - Port-Range. The structure of `port_range` block is documented below.
* `protocol` - Protocol.
* `seq_num` - Entry sequence number.

The `port_range` block supports:

* `end_port` - End-Port.
* `id` - Custom entry port range ID.
* `start_port` - Start-Port.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceDefinition can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_internetservicedefinition.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

