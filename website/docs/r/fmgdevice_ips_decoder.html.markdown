---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ips_decoder"
description: |-
  Configure IPS decoder.
---

# fmgdevice_ips_decoder
Configure IPS decoder.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `parameter`: `fmgdevice_ips_decoder_parameter`



## Example Usage

```hcl
resource "fmgdevice_ips_decoder" "trname" {
  name = "your own value"
  parameter {
    name  = "your own value"
    value = "your own value"
  }

  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `name` - Decoder name.
* `parameter` - Parameter. The structure of `parameter` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `parameter` block supports:

* `name` - Parameter name.
* `value` - Parameter value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ips Decoder can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ips_decoder.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

