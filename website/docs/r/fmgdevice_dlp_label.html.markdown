---
subcategory: "DLP"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_label"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure labels used by DLP blocking.
---

# fmgdevice_dlp_label
<i>This object will be purged after policy copy and install.</i> Configure labels used by DLP blocking.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fmgdevice_dlp_label_entries`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Optional comments.
* `connector` - Name of SDN connector.
* `entries` - Entries. The structure of `entries` block is documented below.
* `fabric_force_sync` - Enable/disable forced synchronization of configuration objects from the root FortiGate unit to the downstream devices.  Configuration conflict check is skipped. Valid values: `disable`, `enable`.

* `fabric_object` - Security Fabric global object setting. Valid values: `disable`, `enable`.

* `fabric_object_source` - Source of truth for fabric object. Valid values: `member`, `local`, `root`.

* `mpip_type` - MPIP label type. Valid values: `local`, `remote`.

* `name` - Name of table containing the label.
* `type` - Label type. Valid values: `mpip`, `fortidata`.

* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `fortidata_label_name` - Name of FortiData label
* `guid` - MPIP label guid.
* `id` - ID.
* `mpip_label_name` - Name of MPIP label.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp Label can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_label.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

