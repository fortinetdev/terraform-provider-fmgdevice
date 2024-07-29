---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_firewall_internetservicedefinition_entry"
description: |-
  Protocol and port information in an Internet Service entry.
---

# fmgdevice_firewall_internetservicedefinition_entry
Protocol and port information in an Internet Service entry.

~> This resource is a sub resource for variable `entry` of resource `fmgdevice_firewall_internetservicedefinition`. Conflict and overwrite may occur if use both of them.
The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `port_range`: `fmgdevice_firewall_internetservicedefinition_entry_portrange`



## Example Usage

```hcl
resource "fmgdevice_firewall_internetservicedefinition_entry" "trname" {
  category_id = 10
  name        = "your own value"
  port_range {
    end_port   = 10
    fosid      = 10
    start_port = 10
  }

  protocol    = 10
  seq_num     = 10
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `internet_service_definition` - Internet Service Definition.

* `category_id` - Category-Id.
* `name` - Name.
* `port_range` - Port-Range. The structure of `port_range` block is documented below.
* `protocol` - Protocol.
* `seq_num` - Entry sequence number.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `port_range` block supports:

* `end_port` - End-Port.
* `id` - Custom entry port range ID.
* `start_port` - Start-Port.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Firewall InternetServiceDefinitionEntry can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "internet_service_definition=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_firewall_internetservicedefinition_entry.labelname {{seq_num}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

