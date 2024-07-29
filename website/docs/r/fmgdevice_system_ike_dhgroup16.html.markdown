---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ike_dhgroup16"
description: |-
  Diffie-Hellman group 16 (MODP-4096).
---

# fmgdevice_system_ike_dhgroup16
Diffie-Hellman group 16 (MODP-4096).

~> This resource is a sub resource for variable `dh_group_16` of resource `fmgdevice_system_ike`. Conflict and overwrite may occur if use both of them.



## Example Usage

```hcl
resource "fmgdevice_system_ike_dhgroup16" "trname" {
  fosid         = 10
  keypair_cache = "custom"
  keypair_count = 10
  mode          = "software"
  device_name   = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `fosid` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System IkeDhGroup16 can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ike_dhgroup16.labelname SystemIkeDhGroup16
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

