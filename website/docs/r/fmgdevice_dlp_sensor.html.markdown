---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_sensor"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure sensors used by DLP blocking.
---

# fmgdevice_dlp_sensor
<i>This object will be purged after policy copy and install.</i> Configure sensors used by DLP blocking.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `entries`: `fmgdevice_dlp_sensor_entries`
>- `filter`: `fmgdevice_dlp_sensor_filter`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Optional comments.
* `entries` - Entries. The structure of `entries` block is documented below.
* `eval` - Expression to evaluate.
* `fgd_id` - ID of object in FortiGuard database.
* `match_type` - Logical relation between entries (default = match-any). Valid values: `match-all`, `match-any`, `match-eval`.

* `name` - Name of table containing the sensor.
* `dlp_log` - Dlp-Log. Valid values: `disable`, `enable`.

* `extended_log` - Extended-Log. Valid values: `disable`, `enable`.

* `feature_set` - Flow/proxy feature set. Valid values: `proxy`, `flow`.

* `filter` - Filter. The structure of `filter` block is documented below.
* `full_archive_proto` - Full-Archive-Proto. Valid values: `imap`, `smtp`, `pop3`, `ftp`, `nntp`, `mapi`, `http-get`, `http-post`, `cifs`, `ssh`.

* `nac_quar_log` - Nac-Quar-Log. Valid values: `disable`, `enable`.

* `replacemsg_group` - Replacemsg-Group.
* `summary_proto` - Summary-Proto. Valid values: `imap`, `smtp`, `pop3`, `ftp`, `nntp`, `mapi`, `http-get`, `http-post`, `cifs`, `ssh`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `count` - Count of dictionary matches to trigger sensor entry match (Dictionary might not be able to trigger more than once based on its 'repeat' option, 1 - 255, default = 1).
* `dictionary` - Select a DLP dictionary or exact-data-match.
* `id` - ID.
* `status` - Enable/disable this entry. Valid values: `disable`, `enable`.


The `filter` block supports:

* `action` - Action. Valid values: `log-only`, `block`, `quarantine-ip`, `allow`.

* `archive` - Archive. Valid values: `disable`, `enable`.

* `company_identifier` - Company-Identifier.
* `expiry` - Expiry.
* `file_size` - File-Size.
* `file_type` - File-Type.
* `filter_by` - Filter-By. Valid values: `credit-card`, `ssn`, `regexp`, `file-type`, `file-size`, `fingerprint`, `watermark`, `encrypted`, `file-type-and-size`.

* `id` - Id.
* `match_percentage` - Match-Percentage.
* `name` - Name.
* `proto` - Proto. Valid values: `imap`, `smtp`, `pop3`, `ftp`, `nntp`, `mapi`, `http-get`, `http-post`, `cifs`, `ssh`.

* `regexp` - Regexp.
* `sensitivity` - Sensitivity.
* `severity` - Severity. Valid values: `info`, `low`, `medium`, `high`, `critical`.

* `type` - Type. Valid values: `file`, `message`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp Sensor can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_sensor.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

