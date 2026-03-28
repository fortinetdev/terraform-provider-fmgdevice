---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_system_nethsm"
description: |-
  Device SystemNethsm
---

# fmgdevice_system_nethsm
Device SystemNethsm

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `hagroups`: `fmgdevice_system_nethsm_hagroups`
>- `partitions`: `fmgdevice_system_nethsm_partitions`
>- `servers`: `fmgdevice_system_nethsm_servers`
>- `slots`: `fmgdevice_system_nethsm_slots`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.

* `ha` - Ha. Valid values: `disable`, `enable`.

* `ha_status_pulling_interval` - Ha-Status-Pulling-Interval.
* `hagroups` - Hagroups. The structure of `hagroups` block is documented below.
* `interface` - Interface.
* `partitions` - Partitions. The structure of `partitions` block is documented below.
* `primus_cfg` - Primus-Cfg.
* `receivetimeout` - Receivetimeout.
* `rsa_mech_remap` - Rsa-Mech-Remap. Valid values: `disable`, `enable`.

* `secret_content` - Secret-Content.
* `servers` - Servers. The structure of `servers` block is documented below.
* `slots` - Slots. The structure of `slots` block is documented below.
* `status` - Status. Valid values: `disable`, `enable`.

* `vendor` - Vendor. Valid values: `SafeNet`, `safenet`, `primus`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `hagroups` block supports:

* `member` - Member.
* `name` - Name.

The `partitions` block supports:

* `name` - Name.
* `pkcs11_pin` - Pkcs11-Pin.
* `slot_id` - Slot-Id.

The `servers` block supports:

* `htl` - Htl. Valid values: `disable`, `enable`.

* `name` - Name.
* `port` - Port.
* `server` - Server.
* `server_cert` - Server-Cert.

The `slots` block supports:

* `for_ha` - For-Ha. Valid values: `no`, `yes`.

* `id` - Id.
* `name` - Name.
* `password` - Password.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Nethsm can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_system_nethsm.labelname SystemNethsm
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

