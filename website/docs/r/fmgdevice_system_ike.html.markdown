---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_ike"
description: |-
  Configure IKE global attributes.
---

# fmgdevice_system_ike
Configure IKE global attributes.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `dh_group_1`: `fmgdevice_system_ike_dhgroup1`
>- `dh_group_14`: `fmgdevice_system_ike_dhgroup14`
>- `dh_group_15`: `fmgdevice_system_ike_dhgroup15`
>- `dh_group_16`: `fmgdevice_system_ike_dhgroup16`
>- `dh_group_17`: `fmgdevice_system_ike_dhgroup17`
>- `dh_group_18`: `fmgdevice_system_ike_dhgroup18`
>- `dh_group_19`: `fmgdevice_system_ike_dhgroup19`
>- `dh_group_2`: `fmgdevice_system_ike_dhgroup2`
>- `dh_group_20`: `fmgdevice_system_ike_dhgroup20`
>- `dh_group_21`: `fmgdevice_system_ike_dhgroup21`
>- `dh_group_27`: `fmgdevice_system_ike_dhgroup27`
>- `dh_group_28`: `fmgdevice_system_ike_dhgroup28`
>- `dh_group_29`: `fmgdevice_system_ike_dhgroup29`
>- `dh_group_30`: `fmgdevice_system_ike_dhgroup30`
>- `dh_group_31`: `fmgdevice_system_ike_dhgroup31`
>- `dh_group_32`: `fmgdevice_system_ike_dhgroup32`
>- `dh_group_5`: `fmgdevice_system_ike_dhgroup5`



## Example Usage

```hcl
resource "fmgdevice_system_ike" "trname" {
  dh_group_1 {
    fosid         = 10
    keypair_cache = "global"
    keypair_count = 10
    mode          = "hardware"
  }

  dh_group_14 {
    fosid         = 10
    keypair_cache = "global"
    keypair_count = 10
    mode          = "software"
  }

  dh_group_15 {
    fosid         = 10
    keypair_cache = "custom"
    keypair_count = 10
    mode          = "hardware"
  }

  dh_group_16 {
    fosid         = 10
    keypair_cache = "custom"
    keypair_count = 10
    mode          = "global"
  }

  dh_group_17 {
    fosid         = 10
    keypair_cache = "custom"
    keypair_count = 10
    mode          = "hardware"
  }

  device_name = var.device_name # not required if setting is at provider
}
```

## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `dh_group_1` - Dh-Group-1. The structure of `dh_group_1` block is documented below.
* `dh_group_14` - Dh-Group-14. The structure of `dh_group_14` block is documented below.
* `dh_group_15` - Dh-Group-15. The structure of `dh_group_15` block is documented below.
* `dh_group_16` - Dh-Group-16. The structure of `dh_group_16` block is documented below.
* `dh_group_17` - Dh-Group-17. The structure of `dh_group_17` block is documented below.
* `dh_group_18` - Dh-Group-18. The structure of `dh_group_18` block is documented below.
* `dh_group_19` - Dh-Group-19. The structure of `dh_group_19` block is documented below.
* `dh_group_2` - Dh-Group-2. The structure of `dh_group_2` block is documented below.
* `dh_group_20` - Dh-Group-20. The structure of `dh_group_20` block is documented below.
* `dh_group_21` - Dh-Group-21. The structure of `dh_group_21` block is documented below.
* `dh_group_27` - Dh-Group-27. The structure of `dh_group_27` block is documented below.
* `dh_group_28` - Dh-Group-28. The structure of `dh_group_28` block is documented below.
* `dh_group_29` - Dh-Group-29. The structure of `dh_group_29` block is documented below.
* `dh_group_30` - Dh-Group-30. The structure of `dh_group_30` block is documented below.
* `dh_group_31` - Dh-Group-31. The structure of `dh_group_31` block is documented below.
* `dh_group_32` - Dh-Group-32. The structure of `dh_group_32` block is documented below.
* `dh_group_5` - Dh-Group-5. The structure of `dh_group_5` block is documented below.
* `dh_keypair_cache` - Enable/disable Diffie-Hellman key pair cache. Valid values: `disable`, `enable`.

* `dh_keypair_count` - Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
* `dh_keypair_throttle` - Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `disable`, `enable`.

* `dh_mode` - Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.

* `dh_multiprocess` - Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `disable`, `enable`.

* `dh_worker_count` - Number of Diffie-Hellman workers to start.
* `embryonic_limit` - Maximum number of IPsec tunnels to negotiate simultaneously.

The `dh_group_1` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_14` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_15` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_16` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_17` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_18` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_19` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_2` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_20` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_21` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_27` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_28` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_29` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_30` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_31` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_32` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.


The `dh_group_5` block supports:

* `id` - Diffie-Hellman group ID.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.

* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).
* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ike can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_ike.labelname SystemIke
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

