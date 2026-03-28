---
subcategory: "No Category"
layout: "fmgdevice"
page_title: "FortiManager Device: fmgdevice_dlp_profile"
description: |-
  <i>This object will be purged after policy copy and install.</i> Configure DLP profiles.
---

# fmgdevice_dlp_profile
<i>This object will be purged after policy copy and install.</i> Configure DLP profiles.

~> The following variables have sub resource. Avoid using them together, otherwise conflicts and overwrites may occur.
>- `rule`: `fmgdevice_dlp_profile_rule`



## Argument Reference


The following arguments are supported:

* `device_name` - FortiManager managed device name. This variable is used in the request URL. If not specified, it will inherit the variable `device_name` of the provider.
* `device_vdom` - FortiManager managed device vdom. This variable is used in the request URL. If not specified, it will inherit the variable `device_vdom` of the provider.

* `comment` - Comment.
* `dlp_log` - Enable/disable DLP logging. Valid values: `disable`, `enable`.

* `extended_log` - Enable/disable extended logging for data loss prevention. Valid values: `disable`, `enable`.

* `feature_set` - Flow/proxy feature set. Valid values: `flow`, `proxy`.

* `fortidata_error_action` - Action to take if FortiData query fails. Valid values: `block`, `log-only`, `ignore`.

* `full_archive_proto` - Protocols to always content archive. Valid values: `smtp`, `pop3`, `imap`, `http-post`, `http-get`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.

* `nac_quar_log` - Enable/disable NAC quarantine logging. Valid values: `disable`, `enable`.

* `name` - Name of the DLP profile.
* `replacemsg_group` - Replacement message group used by this DLP profile.
* `rule` - Rule. The structure of `rule` block is documented below.
* `summary_proto` - Protocols to always log summary. Valid values: `smtp`, `pop3`, `imap`, `http-post`, `http-get`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.

* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `rule` block supports:

* `action` - Action to take with content that this DLP profile matches. Valid values: `log-only`, `block`, `quarantine-ip`, `allow`.

* `archive` - Enable/disable DLP archiving. Valid values: `disable`, `enable`.

* `expiry` - Quarantine duration in days, hours, minutes (format = dddhhmm).
* `file_size` - Match files greater than or equal to this size (KB).
* `file_type` - Select the number of a DLP file pattern table to match.
* `filter_by` - Select the type of content to match. Valid values: `fingerprint`, `sensor`, `encrypted`, `none`, `mip`, `label`.

* `id` - ID.
* `label` - Select DLP label.
* `match_percentage` - Percentage of fingerprints in the fingerprint databases designated with the selected sensitivity to match.
* `name` - Filter name.
* `proto` - Check messages or files over one or more of these protocols. Valid values: `smtp`, `pop3`, `imap`, `http-post`, `http-get`, `ftp`, `nntp`, `mapi`, `ssh`, `cifs`.

* `sensitivity` - Select a DLP file pattern sensitivity to match.
* `sensor` - Select DLP sensors.
* `severity` - Select the severity or threat level that matches this filter. Valid values: `info`, `low`, `medium`, `high`, `critical`.

* `type` - Select whether to check the content of messages (an email message) or files (downloaded files or email attachments). Valid values: `file`, `message`.



## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Dlp Profile can be imported using any of these accepted formats:
```
Set import_options = ["device_name=YOUR_VALUE", "device_vdom=YOUR_VALUE"] in the provider section.

$ export "FORTIMANAGER_IMPORT_TABLE"="true"
$ terraform import fmgdevice_dlp_profile.labelname {{name}}
$ unset "FORTIMANAGER_IMPORT_TABLE"
```

