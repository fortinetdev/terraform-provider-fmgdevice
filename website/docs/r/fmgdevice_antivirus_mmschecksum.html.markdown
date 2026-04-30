---
subcategory: "Antivirus"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_antivirus_mmschecksum"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure MMS content checksum list.
---

# fmgdevice_antivirus_mmschecksum
<i>This object will be purged after policy copy and install.</i> Configure MMS content checksum list.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fmgdevice_antivirus_mmschecksum_entries`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Optional comments.
* `entries` - Entries. The structure of `entries` block is documented below.
* `fosid` - ID.
* `name` - Name of table.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `checksum` - MMS attachment checksum value (CRC32)
* `name` - entry name, for administrator reference
* `status` - apply this entry during attachment inspection Valid values: `disable`, `enable`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Antivirus MmsChecksum can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_antivirus_mmschecksum.labelname {{fosid}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

