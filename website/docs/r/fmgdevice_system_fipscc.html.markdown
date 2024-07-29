---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_fipscc"
description: |-
  Configure FIPS-CC mode.
---

# fmgdevice_system_fipscc
Configure FIPS-CC mode.

## Example Usage

```hcl
resource "fmgdevice_system_fipscc" "trname" {
  entropy_token            = "disable"
  key_generation_self_test = "disable"
  status                   = "disable"
  device_name              = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `entropy_token` - Enable/disable/dynamic entropy token. Valid values: `disable`, `enable`, `dynamic`.

* `key_generation_self_test` - Enable/disable self tests after key generation. Valid values: `disable`, `enable`.

* `self_test_period` - Self test period.
* `status` - Enable/disable ciphers for FIPS mode of operation. Valid values: `disable`, `enable`, `fips-ciphers`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FipsCc can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_fipscc.labelname SystemFipsCc
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

