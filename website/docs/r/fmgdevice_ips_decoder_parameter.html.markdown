---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_ips_decoder_parameter"
description: |-
  IPS group parameters.
---

# fmgdevice_ips_decoder_parameter
IPS group parameters.

~> This resource is a sub resource for variable `parameter` of resource `fmgdevice_ips_decoder`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_ips_decoder_parameter" "trname" {
  name        = "your own value"
  value       = "your own value"
  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `decoder` - Decoder.

* `name` - Parameter name.
* `value` - Parameter value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ips DecoderParameter can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "decoder=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_ips_decoder_parameter.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

