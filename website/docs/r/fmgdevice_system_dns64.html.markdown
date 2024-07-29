---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_dns64"
description: |-
  Configure DNS64.
---

# fmgdevice_system_dns64
Configure DNS64.

## Example Usage

```hcl
resource "fmgdevice_system_dns64" "trname" {
  always_synthesize_aaaa_record = "enable"
  dns64_prefix                  = "your own value"
  status                        = "disable"
  device_name                   = var.device_name # not required if setting is at provider
  device_vdom                   = var.device_vdom # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `always_synthesize_aaaa_record` - Enable/disable AAAA record synthesis (default = enable). Valid values: `disable`, `enable`.

* `dns64_prefix` - DNS64 prefix must be ::/96 (default = 64:ff9b::/96).
* `status` - Enable/disable DNS64 (default = disable). Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Dns64 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_dns64.labelname SystemDns64
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

